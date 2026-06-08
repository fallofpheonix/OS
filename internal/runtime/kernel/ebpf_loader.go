// Package kernel implements the sensing and enforcement layer for PhoenixOS.
// Domain Logic: Loads eBPF programs into the kernel to capture telemetry and enforce reflexive security policies.
// Responsibility: Bridges kernel-space observability to user-space analytical layers and provides kernel-level actuation.
package kernel

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
)

// ExecEvent represents a process execution event captured in kernel space.
// Concurrency: Read-only instances are thread-safe after decoding.
// State Management: Maps binary data from the kernel ring buffer to a structured format.
type ExecEvent struct {
	Pid        uint32
	Ppid       uint32
	Tgid       uint32
	NsproxyIno uint32
	Uid        uint32
	Comm       [16]byte
	Filename   [128]byte
}

// Loader is the main controller for eBPF lifecycle and event polling.
// Concurrency: Thread-safe for BPF map updates. Event polling runs in a dedicated goroutine.
// State Management: Manages BPF collections, links, and the kernel-level blocklist map.
type Loader struct {
	coll        *ebpf.Collection
	links       []link.Link
	Pub         EventPublisher
	rb          *ringbuf.Reader
	blockedPids *ebpf.Map
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewLoader initializes a new eBPF loader instance.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewLoader(pub EventPublisher) *Loader {
	return &Loader{Pub: pub}
}

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// Load performs eBPF initialization, attaches hooks, and starts event polling.
// I/O: Loads eBPF object from disk, attaches hooks to the kernel.
// Side Effects: Modifies kernel state by attaching BPF programs. Starts a polling goroutine.
// Complexity: O(A) where A is the number of attachment points.
func (l *Loader) Load(path string) error {
	spec, err := ebpf.LoadCollectionSpec(path)
	if err != nil {
		return fmt.Errorf("failed to load ebpf spec: %w", err)
	}

	l.coll, err = ebpf.NewCollection(spec)
	if err != nil {
		return fmt.Errorf("failed to create ebpf collection: %w", err)
	}

	// T4.3: Record eBPF program version in ledger
	if l.Pub != nil {
		bytecode, _ := os.ReadFile(path)
		hash := sha256.Sum256(bytecode)
		l.Pub.Publish("system.lifecycle", TelemetryEvent{
			EventID:   "EBPF-LOAD",
			Source:    "kernel.loader",
			EventType: "EBPFProgramLoaded",
			Severity:  0.1,
			Payload:   []byte(fmt.Sprintf(`{"program": "%s", "sha256": "%x"}`, path, hash)),
		})
	}

	// 1. Attach Tracepoint
	prog := l.coll.Programs["handle_execve"]
	if prog == nil {
		return fmt.Errorf("program handle_execve not found")
	}

	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", prog, nil)
	if err != nil {
		return fmt.Errorf("failed to attach tracepoint: %w", err)
	}
	l.links = append(l.links, tp)

	// 2. Attach LSM Programs
	lsmProg := l.coll.Programs["phoenix_enforce_exec"]
	if lsmProg != nil {
		lsmLink, err := link.AttachLSM(link.LSMOptions{Program: lsmProg})
		if err == nil {
			l.links = append(l.links, lsmLink)
		}
	}

	mproProg := l.coll.Programs["phoenix_enforce_mprotect"]
	if mproProg != nil {
		mproLink, err := link.AttachLSM(link.LSMOptions{Program: mproProg})
		if err == nil {
			l.links = append(l.links, mproLink)
		}
	}

	// 3. Setup Maps and Ring Buffer
	l.blockedPids = l.coll.Maps["blocked_pids"]
	if l.blockedPids == nil {
		return fmt.Errorf("map blocked_pids not found")
	}

	rbMap := l.coll.Maps["rb"]
	l.rb, err = ringbuf.NewReader(rbMap)
	if err != nil {
		return fmt.Errorf("failed to create ringbuf reader: %w", err)
	}

	go l.pollEvents()

	return nil
}

func (l *Loader) pollEvents() {
	if l.rb == nil {
		return
	}

	for {
		record, err := l.rb.Read()
		if err != nil {
			if errors.Is(err, ringbuf.ErrClosed) {
				return
			}
			continue
		}

		var event ExecEvent
		if err := binary.Read(bytes.NewReader(record.RawSample), binary.LittleEndian, &event); err != nil {
			continue
		}

		if l.Pub != nil {
			l.Pub.Publish("kernel.exec", TelemetryEvent{
				Source:    "kernel",
				EventType: "process.exec",
				Severity:  0.5,
				Payload:   record.RawSample,
			})
		}
	}
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// BlockPID writes a PID to the kernel-space blocked_pids map.
// I/O: None.
// Side Effects: Modifies kernel-level BPF map. Causes immediate denial of execve for target PID.
// Complexity: O(1) BPF map update.
func (l *Loader) BlockPID(pid uint32) error {
	if l.blockedPids == nil {
		return nil
	}
	var action uint32 = 1
	return l.blockedPids.Update(pid, action, ebpf.UpdateAny)
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// UnblockPID removes a PID from the kernel blocklist.
// I/O: None.
// Side Effects: Modifies kernel-level BPF map.
// Complexity: O(1) BPF map deletion.
func (l *Loader) UnblockPID(pid uint32) error {
	if l.blockedPids == nil {
		return nil
	}
	return l.blockedPids.Delete(pid)
}

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// Close releases all eBPF resources and kernel links.
// I/O: None.
// Side Effects: Unloads BPF programs and tracepoints. Closes ring buffer.
// Complexity: O(L) where L is the number of links.
func (l *Loader) Close() {
	if l.rb != nil {
		_ = l.rb.Close()
	}
	for _, link := range l.links {
		_ = link.Close()
	}
	if l.coll != nil {
		l.coll.Close()
	}
}
