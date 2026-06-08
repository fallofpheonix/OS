/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#!/usr/bin/env node

const fs = require("node:fs")
const path = require("node:path")
const { execFileSync } = require("node:child_process")

const rootDir = path.resolve(__dirname, "..")
const opencodeDir = path.join(rootDir, ".opencode")
const distDir = path.join(opencodeDir, "dist")

fs.rmSync(distDir, { recursive: true, force: true })

let tscEntrypoint

try {
  tscEntrypoint = require.resolve("typescript/bin/tsc", { paths: [rootDir] })
} catch {
  throw new Error(
    "TypeScript compiler not found. Install root dev dependencies before publishing so .opencode/dist can be built."
  )
}

execFileSync(process.execPath, [tscEntrypoint, "-p", path.join(opencodeDir, "tsconfig.json")], {
  cwd: rootDir,
  stdio: "inherit",
})
