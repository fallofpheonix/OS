/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { combineReducers, Reducer } from 'redux'

import authReducer from '@/slices/auth'
import categoryReducer from '@/slices/category'
import noteReducer from '@/slices/note'
import settingsReducer from '@/slices/settings'
import syncReducer from '@/slices/sync'
import { RootState } from '@/types'

const rootReducer: Reducer<RootState> = combineReducers<RootState>({
  authState: authReducer,
  categoryState: categoryReducer,
  noteState: noteReducer,
  settingsState: settingsReducer,
  syncState: syncReducer,
})

export default rootReducer
