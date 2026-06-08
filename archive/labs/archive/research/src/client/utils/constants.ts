/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { Folder, NotesSortKey, DirectionText } from '@/utils/enums'

export const folderMap: Record<Folder, string> = {
  [Folder.ALL]: 'All Notes',
  [Folder.FAVORITES]: 'Favorites',
  [Folder.SCRATCHPAD]: 'Scratchpad',
  [Folder.TRASH]: 'Trash',
  [Folder.CATEGORY]: 'Category',
}

export const iconColor = 'rgba(255, 255, 255, 0.25)'

export const shortcutMap = [
  { action: 'Create a new note', key: 'N' },
  { action: 'Delete a note', key: 'U' },
  { action: 'Create a category', key: 'C' },
  { action: 'Download a note', key: 'O' },
  { action: 'Sync all notes', key: 'L' },
  { action: 'Markdown preview', key: 'P' },
  { action: 'Toggle theme', key: 'K' },
  { action: 'Search notes', key: 'F' },
  { action: 'Prettify a note', key: 'I' },
]

export const notesSortOptions = [
  { value: NotesSortKey.TITLE, label: 'Title' },
  { value: NotesSortKey.CREATED_DATE, label: 'Date Created' },
  { value: NotesSortKey.LAST_UPDATED, label: 'Last Updated' },
]

export const directionTextOptions = [
  { value: DirectionText.LEFT_TO_RIGHT, label: 'Left to right' },
  { value: DirectionText.RIGHT_TO_LEFT, label: 'Right to left' },
]
