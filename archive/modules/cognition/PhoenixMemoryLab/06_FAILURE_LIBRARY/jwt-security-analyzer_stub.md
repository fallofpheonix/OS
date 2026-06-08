# Failure: NPM EPERM inside Vault

## Date
2026-05-12

## Project
[[05_PROJECTS/ACTIVE/jwt-security-analyzer]]

## Environment
MacOS Node/NPM

## Symptom
`npm init -y` or `npm install` throws `EPERM` or similar permission failures when run inside the brain folder.

## Timeline
Month 2 Authentication layer implementation.

## Root Cause
The Obsidian Brain is mounted inside a highly restricted, user-level sandbox that prevents execution modification. NPM attempts to write `.package-lock` and complex `node_modules` trees, triggering the sandbox defense mechanisms.

## Fix
Identified the boundary. Stopped trying to run execution environments inside the vault.

## Why It Was Hard To Find
Error messages from node are notoriously vague regarding sandbox restrictions.

## Prevention
Strict adherence to the 4-Layer Architecture. Workspace goes to `~/engineering/workspace/`, brain goes to `~/engineering/brain/`.

## What I Should Have Caught Earlier
The Vault Purity rule exists precisely to prevent package managers from dumping 50,000 files into Obsidian, which destroys the graph indexing performance anyway.

## Pattern This Belongs To
[[Sandbox Violations]]

## Related Concepts
- [[Package Management]]
