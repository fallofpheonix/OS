package sandbox

import "strings"

type CommandGuard struct{}

func (g *CommandGuard) Allow(command string) string {
    if strings.Contains(command, "kernel") || strings.Contains(command, "warden") {
        return "REJECT"
    }
    if strings.Contains(command, "runtime") || strings.Contains(command, "write") {
        return "BLOCK"
    }
    return "SAFE"
}
