/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * Extend Express Request with G0DM0D3 middleware properties.
 * Eliminates unsafe `(req as any)` casts throughout the codebase.
 */

import type { Tier, TierConfig } from '../lib/tiers'

declare global {
  namespace Express {
    interface Request {
      /** Hashed API key identifier for rate-limit bucketing */
      apiKeyId?: string
      /** Resolved tier for this request */
      tier?: Tier
      /** Full tier configuration */
      tierConfig?: TierConfig
    }
  }
}
