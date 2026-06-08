/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
#!/usr/bin/env node
'use strict';

const { isHookEnabled } = require('../lib/hook-flags');

const [, , hookId, profilesCsv] = process.argv;
if (!hookId) {
  process.stdout.write('yes');
  process.exit(0);
}

process.stdout.write(isHookEnabled(hookId, { profiles: profilesCsv }) ? 'yes' : 'no');
