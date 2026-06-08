/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { PayloadAction } from '@reduxjs/toolkit'

import reducer, {
  initialState,
  toggleSettingsModal,
  togglePreviewMarkdown,
  updateCodeMirrorOption,
  toggleDarkTheme,
} from '@/slices/settings'

describe('settings slice', () => {
  it('should return the initial state on first run', () => {
    const nextState = initialState
    const action = {} as PayloadAction
    const result = reducer(undefined, action)

    expect(result).toEqual(nextState)
  })

  it('should toggle open state', () => {
    const nextState = { ...initialState, isOpen: true }
    const result = reducer(initialState, toggleSettingsModal())

    expect(result).toEqual(nextState)
  })

  it('should update code mirror option', () => {
    const payload = { key: 'key123', value: 'mirror' }
    const state = {
      ...initialState,
      codeMirrorOptions: {
        ...initialState.codeMirrorOptions,
        [payload.key]: payload.value,
      },
    }
    const result = reducer(initialState, updateCodeMirrorOption(payload))

    expect(result).toEqual(state)
  })

  it('should toggle preview markdown state', () => {
    const nextState = { ...initialState, previewMarkdown: !initialState.previewMarkdown }
    const result = reducer(initialState, togglePreviewMarkdown())

    expect(result).toEqual(nextState)
  })

  it('should toggle dark theme state', () => {
    const nextState = { ...initialState, darkTheme: !initialState.darkTheme }
    const result = reducer(initialState, toggleDarkTheme())

    expect(result).toEqual(nextState)
  })
})
