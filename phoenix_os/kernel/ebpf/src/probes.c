#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>

/* 
 * PhoenixOS Kernel Telemetry Probes (P2A)
 * Goal: Low-overhead ingestion of process, I/O, and network events.
 */

struct event_t {
    __u64 timestamp;
    __u32 pid;
    __u32 ppid;
    char process[16];
    char syscall[16];
    __u32 cpu;
};

// Ring buffer for high-throughput event transfer to user space
struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24); // 16MB buffer
} events SEC(".maps");

// Minimal task_struct for CO-RE (Compile Once - Run Everywhere)
struct task_struct {
    struct task_struct *real_parent;
    int tgid;
} __attribute__((preserve_access_index));

static __always_inline void submit_event(void *ctx, const char *syscall_name) {
    struct event_t *e;

    e = bpf_ringbuf_reserve(&events, sizeof(*e), 0);
    if (!e) {
        return;
    }

    e->timestamp = bpf_ktime_get_ns();
    e->pid = bpf_get_current_pid_tgid() >> 32;

    struct task_struct *task = (struct task_struct *)bpf_get_current_task();
    e->ppid = BPF_CORE_READ(task, real_parent, tgid);

    bpf_get_current_comm(&e->process, sizeof(e->process));
    
    // Copy syscall name directly
    int i;
    for (i = 0; i < 15 && syscall_name[i] != '\0'; i++) {
        e->syscall[i] = syscall_name[i];
    }
    e->syscall[i] = '\0';

    e->cpu = bpf_get_smp_processor_id();

    bpf_ringbuf_submit(e, 0);
}

// --- Process Lifecycle ---
SEC("tracepoint/syscalls/sys_enter_execve")
int tp_execve(void *ctx) { submit_event(ctx, "execve"); return 0; }

SEC("tracepoint/syscalls/sys_enter_fork")
int tp_fork(void *ctx) { submit_event(ctx, "fork"); return 0; }

SEC("tracepoint/syscalls/sys_enter_clone")
int tp_clone(void *ctx) { submit_event(ctx, "clone"); return 0; }

SEC("tracepoint/syscalls/sys_enter_exit")
int tp_exit(void *ctx) { submit_event(ctx, "exit"); return 0; }

// --- I/O Operations ---
SEC("tracepoint/syscalls/sys_enter_open")
int tp_open(void *ctx) { submit_event(ctx, "open"); return 0; }

SEC("tracepoint/syscalls/sys_enter_read")
int tp_read(void *ctx) { submit_event(ctx, "read"); return 0; }

SEC("tracepoint/syscalls/sys_enter_write")
int tp_write(void *ctx) { submit_event(ctx, "write"); return 0; }

// --- Network Operations ---
SEC("tracepoint/syscalls/sys_enter_connect")
int tp_connect(void *ctx) { submit_event(ctx, "connect"); return 0; }

SEC("tracepoint/syscalls/sys_enter_accept")
int tp_accept(void *ctx) { submit_event(ctx, "accept"); return 0; }

SEC("tracepoint/syscalls/sys_enter_bind")
int tp_bind(void *ctx) { submit_event(ctx, "bind"); return 0; }

char LICENSE[] SEC("license") = "GPL";
