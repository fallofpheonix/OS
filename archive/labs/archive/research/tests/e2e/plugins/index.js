/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
const clipboardy = require('clipboardy')

const cypressTypeScriptPreprocessor = require('./cy-ts-preprocessor')

module.exports = (on) => {
  on('file:preprocessor', cypressTypeScriptPreprocessor)

  on('task', {
    getClipboard() {
      return clipboardy.readSync()
    },
  })
}
