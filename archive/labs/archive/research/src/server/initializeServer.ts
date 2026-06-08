/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import path from 'path'

import express, { Router } from 'express'
import cookieParser from 'cookie-parser'
import cors from 'cors'
import helmet from 'helmet'
import compression from 'compression'

export default function initializeServer(router: Router) {
  const app = express()
  const isProduction = process.env.NODE_ENV === 'production'
  const origin = { origin: isProduction ? false : '*' }

  app.set('trust proxy', 1)
  app.use(express.json())
  app.use(cookieParser())
  app.use(cors(origin))
  app.use(helmet())
  app.use(compression())

  app.use((request, response, next) => {
    response.header('Content-Security-Policy', "img-src 'self' *.githubusercontent.com")

    return next()
  })

  app.use(express.static(path.join(__dirname, '../../dist/')))
  app.use('/api', router)
  app.get('*', (request, response) => {
    response.sendFile(path.join(__dirname, '../../dist/index.html'))
  })

  return app
}
