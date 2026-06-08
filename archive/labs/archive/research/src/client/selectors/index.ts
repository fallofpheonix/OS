/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { RootState } from '@/types'

export const getSettings = (state: RootState) => state.settingsState
export const getCategories = (state: RootState) => state.categoryState
export const getNotes = (state: RootState) => state.noteState
export const getSync = (state: RootState) => state.syncState
export const getAuth = (state: RootState) => state.authState
