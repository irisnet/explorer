/* eslint-disable no-console */

import { unregister } from 'register-service-worker'

if (process.env.NODE_ENV === 'production') {
  unregister()
}
