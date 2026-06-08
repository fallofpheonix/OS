package bus

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	ledgerv1 "github.com/fallofpheonix/phoenix/foundation/contracts/ledger/v1"
	validationv1 "github.com/fallofpheonix/phoenix/foundation/contracts/validation/v1"
)

// PURPOSE: Orchestrates the serialized ingestion of events into the ledger.
// CONTRACT: Acts as the single-threaded gatekeeper for the durable forensic record.
//           It ensures that concurrent events from the Bus are assigned
//           a deterministic linear order before persistence.
// CONNECTS: foundation/runtime/bus (source)
//           foundation/ledger (sink)

// ApplierConfig defines the operational parameters for the event loop.
type ApplierConfig struct {
	BufferSize int
	Ledger     ledgerv1.LedgerWriter
	Bus        *Bus
	Validator  validationv1.SemanticValidator
	Topics     []string
}

// Applier manages the lifecycle and execution of the ledger ingestion loop.
type Applier struct {
	config       ApplierConfig
	intake       chan TelemetryEvent
	droppedCount uint64
	running      atomic.Bool
}

// NewApplier initializes the serialization engine.
// PURPOSE: Constructor for the Applier.
// CONTRACT: config.Ledger and config.Bus must not be nil.
func NewApplier(config ApplierConfig) *Applier {
	if config.BufferSize <= 0 {
		config.BufferSize = 1024
	}
	return &Applier{
		config: config,
		intake: make(chan TelemetryEvent, config.BufferSize),
	}
}

// Start launches the ingestion goroutine and begins subscribing to topics.
// PURPOSE: Activates the Event Loop.
// CONTRACT: Can only be called once. Returns error if already running.
// FAILURE: Returns error if subscription fails or context is already cancelled.
func (a *Applier) Start(ctx context.Context) error {
	if a.running.Swap(true) {
		return fmt.Errorf("applier is already running")
	}

	// WHY: Subscribing to each topic and bridging it to our internal serial intake.
	for _, topic := range a.config.Topics {
		ch := a.config.Bus.Subscribe(topic)
		go a.bridgeTopic(ctx, ch)
	}

	go a.intakeLoop(ctx)
	return nil
}

// bridgeTopic moves events from a Bus subscriber channel to the Applier intake.
// WHY: Converts many-to-many Bus distribution into many-to-one Applier ingestion.
func (a *Applier) bridgeTopic(ctx context.Context, busCh chan TelemetryEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		case ev := <-busCh:
			// CONTRACT: Non-blocking send to ensure slow disk/ledger doesn't
			// deadlock the high-performance telemetry bus.
			// REASON: Q895 Revision (Error and Escalate policy).
			select {
			case a.intake <- ev:
			default:
				atomic.AddUint64(&a.droppedCount, 1)
			}
		}
	}
}

// intakeLoop processes the serial stream and writes to the ledger.
// PURPOSE: The authoritative single-threaded write loop.
// CONTRACT: Only one instance of this function may run at any time.
func (a *Applier) intakeLoop(ctx context.Context) {
	defer a.running.Store(false)

	for {
		select {
		case <-ctx.Done():
			// WHY: Draining the channel during graceful shutdown to minimize
			// forensic gaps.
			a.drain(500 * time.Millisecond)
			return
		case ev := <-a.intake:
			if err := a.handleEvent(ev); err != nil {
				// FAILURE: Disk full or integrity break triggers halt.
				// WHY: Q898 compliance.
				fmt.Printf("[APPLIER FATAL] Ledger write failed: %v\n", err)
				return
			}
		}
	}
}

// handleEvent translates and persists a single event.
// PURPOSE: Translates, validates, and persists a single event.
func (a *Applier) handleEvent(ev TelemetryEvent) error {
	id, cause, tick, payload := ToLedgerParams(ev)

	// T5.3: Wire SemanticValidator
	if a.config.Validator != nil {
		if err := a.config.Validator.Validate(id, cause, tick, payload); err != nil {
			// REASON: Physics violation or temporal paradox. Log and drop for now,
			// or return error to halt applier if integrity is broken.
			fmt.Printf("[APPLIER REJECT] Semantic validation failed for %s: %v\n", id, err)
			return nil // Don't halt the whole OS for one bad event, but don't ledger it.
		}
	}

	return a.config.Ledger.AddEntry(id, cause, tick, payload)
}

// drain attempts to flush pending events before shutdown.
func (a *Applier) drain(timeout time.Duration) {
	deadline := time.After(timeout)
	for {
		select {
		case ev := <-a.intake:
			_ = a.handleEvent(ev)
		case <-deadline:
			return
		default:
			return
		}
	}
}

// DroppedEvents returns the count of events that failed to enter the intake channel.
func (a *Applier) DroppedEvents() uint64 {
	return atomic.LoadUint64(&a.droppedCount)
}
