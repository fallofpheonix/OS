/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { Request, Response, NextFunction } from 'express'

import { SDK } from '../utils/helpers'
import { Method } from '../utils/enums'

const getUser = async (request: Request, response: Response, next: NextFunction) => {
  const { accessToken } = response.locals

  try {
    const { data } = await SDK(Method.GET, '/user', accessToken)
    response.locals.userData = data

    next()
  } catch (error) {
    response.status(403).send({ message: 'Forbidden Resource', status: 403 })
  }
}

export default getUser
