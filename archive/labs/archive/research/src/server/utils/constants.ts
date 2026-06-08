/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
const isProduction = process.env.NODE_ENV === 'production'

export const thirtyDayCookie = {
  maxAge: 60 * 60 * 1000 * 24 * 30,
  secure: isProduction,
  httpOnly: true,
  sameSite: 'strict' as const,
}
