/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
/**
 * FILE: main.go
 * PATH: Phoenix.Terminus/cli/main.go
 *
 * PURPOSE:
 * Prototype for WARDEN.EXE - The Sovereign Auditor Interface.
 */

package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fallofpheonix/phoenix/platform/crucible/game/context"
	"github.com/fallofpheonix/phoenix/platform/crucible/game/ecs"
	"github.com/fallofpheonix/phoenix/platform/crucible/game/engines"
	"github.com/fallofpheonix/phoenix/platform/crucible/game/simulation"
	"github.com/fallofpheonix/phoenix/foundation/runtime/authority"
)

// Env holds the shared state for the CLI commands.
type Env struct {
	Bus         *simulation.SimulationBus
	Ctx         *context.ApplicationContext
	Entropy     *engines.EntropyEngine
	Court       *engines.CourtEngine
	ActiveToken string
}

// Command defines the interface for all CLI commands.
type Command interface {
	Name() string
	Description() string
	Execute(env *Env, args []string) bool
}

// HelpCommand displays available commands.
type HelpCommand struct {
	registry []Command
}

func (c *HelpCommand) Name() string        { return "help" }
func (c *HelpCommand) Description() string { return "Display available commands" }
func (c *HelpCommand) Execute(env *Env, args []string) bool {
	fmt.Println("Available Commands:")
	for _, cmd := range c.registry {
		fmt.Printf("  %-10s - %s\n", cmd.Name(), cmd.Description())
	}
	return true
}

// StatusCommand displays system health.
type StatusCommand struct{}

func (c *StatusCommand) Name() string        { return "status" }
func (c *StatusCommand) Description() string { return "Display system health and resources" }
func (c *StatusCommand) Execute(env *Env, args []string) bool {
	fmt.Printf("Substrate Entropy: %.2f%%\n", env.Entropy.GlobalValue()*100)
	fmt.Println("Operator Credits: 100") // Mocked
	return true
}

// AuthCommand handles operator authentication.
type AuthCommand struct{}

func (c *AuthCommand) Name() string        { return "auth" }
func (c *AuthCommand) Description() string { return "Authenticate as a Sovereign Auditor (Usage: auth <token>)" }
func (c *AuthCommand) Execute(env *Env, args []string) bool {
	if len(args) < 2 {
		fmt.Println("Usage: auth <token>")
		return true
	}
	env.ActiveToken = args[1]
	fmt.Printf("IDENTITY REGISTERED LOCALLY. ACTIVE TOKEN: %s\n", args[1])
	return true
}

// DriftCommand displays real-time telemetry.
type DriftCommand struct{}

func (c *DriftCommand) Name() string        { return "drift" }
func (c *DriftCommand) Description() string { return "View real-time anomaly telemetry" }
func (c *DriftCommand) Execute(env *Env, args []string) bool {
	fmt.Println("[!] ANOMALY DETECTED: PID 5011 (AuthService) | Drift: 0.42")
	fmt.Println("[!] ANOMALY DETECTED: PID 9022 (KernelMod)   | Drift: 0.08")
	return true
}

// AuditCommand runs deep verification on a target.
type AuditCommand struct{}

func (c *AuditCommand) Name() string        { return "audit" }
func (c *AuditCommand) Description() string { return "Run deep verification on a target" }
func (c *AuditCommand) Execute(env *Env, args []string) bool {
	if len(args) < 2 {
		fmt.Println("Usage: audit <id>")
		return true
	}
	id := args[1]
	fmt.Printf("RUNNING VERIFICATION ON %s...\n", id)
	time.Sleep(1 * time.Second)
	if id == "5011" {
		fmt.Println("EVIDENCE: Unauthorized access to /etc/shadow detected.")
		fmt.Println("CERTAINTY: 0.88")
	} else {
		fmt.Println("NO CRITICAL VIOLATIONS DETECTED.")
	}
	return true
}

// VerdictCommand adjudicates an anomaly.
type VerdictCommand struct{}

func (c *VerdictCommand) Name() string        { return "verdict" }
func (c *VerdictCommand) Description() string { return "Adjudicate an anomaly" }
func (c *VerdictCommand) Execute(env *Env, args []string) bool {
	if len(args) < 2 {
		fmt.Println("Usage: verdict <id>")
		return true
	}
	id := args[1]
	if id == "5011" {
		anomaly := &engines.Anomaly{ID: "A-5011", Target: engines.EntityID(id), Drift: 0.42}
		
		// 1. Get the event and verdict from the court
		ev, v := env.Court.Adjudicate(anomaly, nil)
		fmt.Println(env.Court.Summary(anomaly, v))

		// 2. Validate Capability
		if env.ActiveToken == "" {
			fmt.Println("Authority denied: No active token. Use 'auth <token>' first.")
			return true
		}
		if err := env.Ctx.Auth().Validate(env.ActiveToken, "ENFORCE"); err != nil {
			fmt.Printf("Authority denied: %v\n", err)
			return true
		}
		
		// 3. Dispatch via Bus
		err := env.Bus.Dispatch(ev)
		if err != nil {
			fmt.Printf("Validation error: %v\n", err)
		} else {
			// Update projections
			env.Entropy.Apply(ev.Type, ev.Payload)
			fmt.Println("TRUST +10 | CREDITS +20")
		}
	} else {
		fmt.Println("ID not found or already adjudicated.")
	}
	return true
}

// ExitCommand terminates the session.
type ExitCommand struct{}

func (c *ExitCommand) Name() string        { return "exit" }
func (c *ExitCommand) Description() string { return "Exit WARDEN.EXE" }
func (c *ExitCommand) Execute(env *Env, args []string) bool {
	fmt.Println("SHUTTING DOWN...")
	return false // returning false breaks the REPL loop
}

func main() {
	// Generate random root authority token
	b := make([]byte, 16)
	rand.Read(b)
	rootToken := fmt.Sprintf("ROOT-%x", b)

	// 1. Initialize Engines
	entropy := engines.NewEntropyEngine(0.12, 0.001)
	trust := engines.NewTrustEngine()
	reputation := engines.NewReputationEngine()
	registry := ecs.NewRegistry() // Initialize Registry
	validator := engines.NewSemanticValidator()
	authMgr := authority.NewManager()
	authMgr.RegisterToken(&authority.CapabilityToken{
		ID:    rootToken,
		Scope: []string{"ENFORCE"},
	})
	projs := &context.Projections{
		Entropy:    entropy,
		Trust:      trust,
		Reputation: reputation,
		Registry:   registry,
	}
	ctx := context.NewApplicationContext(validator, projs, authMgr, 1000)
	bus := simulation.NewSimulationBus(ctx, 100)
	court := engines.NewCourtEngine(reputation, trust)

	env := &Env{
		Bus:     bus,
		Ctx:     ctx,
		Entropy: entropy,
		Court:   court,
	}

	commands := []Command{
		&StatusCommand{},
		&AuthCommand{},
		&DriftCommand{},
		&AuditCommand{},
		&VerdictCommand{},
		&ExitCommand{},
	}
	help := &HelpCommand{registry: append(commands, &HelpCommand{})}
	commands = append(commands, help)

	cmdMap := make(map[string]Command)
	for _, cmd := range commands {
		cmdMap[cmd.Name()] = cmd
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("########################################")
	fmt.Println("#           WARDEN.EXE v1.0.0          #")
	fmt.Println("#      SOVEREIGN AUDITOR INTERFACE     #")
	fmt.Println("########################################")
	fmt.Println("BOOTING...")
	time.Sleep(1 * time.Second)
	fmt.Printf("ROOT AUTHORITY TOKEN GENERATED: %s\n", rootToken)
	fmt.Println("Store this token securely. It is required to issue verdicts.")
	fmt.Println("SUBSTRATE STATUS: ACTIVE")
	fmt.Printf("ENTROPY: %.2f%% (STABLE)\n", env.Entropy.GlobalValue()*100)
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("warden@phoenix:~$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.Split(input, " ")
		cmdName := parts[0]

		if cmdName == "quit" {
			fmt.Println("SHUTTING DOWN...")
			break
		}

		cmd, exists := cmdMap[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s. Type 'help' for options.\n", cmdName)
			continue
		}

		if !cmd.Execute(env, parts) {
			break
		}
	}
}
