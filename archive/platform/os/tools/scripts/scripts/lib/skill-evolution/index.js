/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
'use strict';

const provenance = require('./provenance');
const versioning = require('./versioning');
const tracker = require('./tracker');
const health = require('./health');
const dashboard = require('./dashboard');

module.exports = {
  ...provenance,
  ...versioning,
  ...tracker,
  ...health,
  ...dashboard,
  provenance,
  versioning,
  tracker,
  health,
  dashboard,
};
