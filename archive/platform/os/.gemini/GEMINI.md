# ECC for Gemini CLI

This file provides Gemini CLI with the baseline ECC workflow, review standards, and security checks for repositories that install the Gemini target.

## Overview

Everything Claude Code (ECC) is a cross-harness coding system with 36 specialized agents, 142 skills, and 68 commands.

Gemini support is currently focused on a strong project-local instruction layer via `.gemini/GEMINI.md`, plus the shared MCP catalog and package-manager setup assets shipped by the installer.

## Core Workflow

1. Plan before editing large features.
2. Prefer test-first changes for bug fixes and new functionality.
3. Review for security before shipping.
4. Keep changes self-contained, readable, and easy to revert.

## Coding Standards

- Prefer immutable updates over in-place mutation.
- Keep functions small and files focused.
- Validate user input at boundaries.
- Never hardcode secrets.
- Fail loudly with clear error messages instead of silently swallowing problems.

## Security Checklist

Before any commit:

- No hardcoded API keys, passwords, or tokens
- All external input validated
- Parameterized queries for database writes
- Sanitized HTML output where applicable
- Authz/authn checked for sensitive paths
- Error messages scrubbed of sensitive internals

### AgentShield Integration

For comprehensive security auditing, `pheonixos` integrates with `AgentShield`, an external security scanner.

**Installation:**
```bash
npm install -g ecc-agentshield
```

**Usage:**
Run a scan using the `security-scan` command (if integrated), or directly via `npx`:
```bash
/security-scan             # If configured as a command
# OR
npx ecc-agentshield scan   # Direct CLI usage
```

**Features:**
- Scans for vulnerabilities, misconfigurations, and injection risks.
- Covers agent configurations, hooks, rules, and sensitive files.
- Provides options for auto-fixing safe issues (`npx ecc-agentshield scan --fix`) and deep analysis (`npx ecc-agentshield scan --opus`).


## Delivery Standards

- Use conventional commits: `feat`, `fix`, `refactor`, `docs`, `test`, `chore`, `perf`, `ci`
- Run targeted verification for touched areas before shipping
- Prefer contained local implementations over adding new third-party runtime dependencies

## ECC Areas To Reuse

- `AGENTS.md` for repo-wide operating rules
- `skills/` for deep workflow guidance
- `commands/` for slash-command patterns worth adapting into prompts/macros
- `mcp-configs/` for shared connector baselines
