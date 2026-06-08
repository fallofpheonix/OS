/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import express from 'express'
import * as dotenv from 'dotenv'

import authHandler from '../handlers/auth'
import checkAuth from '../middleware/checkAuth'

const router = express.Router()
dotenv.config()

router.get('/callback', authHandler.callback)
router.get('/login', checkAuth, authHandler.login)
router.get('/logout', authHandler.logout)

export default router
