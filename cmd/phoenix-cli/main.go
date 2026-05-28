package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
	unloadCmd := flag.NewFlagSet("unload", flag.ExitOnError)
	statusCmd := flag.NewFlagSet("status", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'load', 'unload' or 'status' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "load":
		loadCmd.Parse(os.Args[2:])
		fmt.Println("Sending LOAD_PROBE directive to Warden...")
		// TODO: Implement Unix Domain Socket (UDS) communication for secure local control.
	case "unload":
		unloadCmd.Parse(os.Args[2:])
		fmt.Println("Sending UNLOAD_PROBE directive to Warden...")
		// TODO: Implement Unix Domain Socket (UDS) communication for secure local control.
	case "status":
		statusCmd.Parse(os.Args[2:])
		fmt.Println("Querying Warden for eBPF status...")
		// TODO: Implement Unix Domain Socket (UDS) communication for secure local control.
	default:
		fmt.Println("expected 'load', 'unload' or 'status' subcommands")
		os.Exit(1)
	}
}
