/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import express from 'express'

import syncHandler from '../handlers/sync'
import checkAuth from '../middleware/checkAuth'
import getUser from '../middleware/getUser'

const router = express.Router()

router.post('/', checkAuth, getUser, syncHandler.sync)
router.get('/notes', checkAuth, getUser, syncHandler.getNotes)
router.get('/categories', checkAuth, getUser, syncHandler.getCategories)

export default router
