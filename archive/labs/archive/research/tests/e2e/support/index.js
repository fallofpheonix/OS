/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import './commands'

// Since before unload alert hangs console tests, the alert has to be disabled
// Check https://github.com/cypress-io/cypress/issues/2118 for more info
Cypress.on('window:before:load', function (win) {
  const original = win.EventTarget.prototype.addEventListener

  win.EventTarget.prototype.addEventListener = function () {
    if (arguments && arguments[0] === 'beforeunload') return

    return original.apply(this, arguments)
  }

  Object.defineProperty(win, 'onbeforeunload', {
    get: function () {},
    set: function () {},
  })
})
