/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
export const environment = {
  production: false,
  apiUrl:
    window.location.hostname === 'localhost'
      ? 'https://notes-app-eta-coral.vercel.app/api'
      : 'https://notes-app-eta-coral.vercel.app/api',
  trackingApiUrl: 'https://visitor-tracking-api.vercel.app/api/visit',
};
