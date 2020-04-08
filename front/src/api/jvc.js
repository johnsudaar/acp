import axios from 'axios'
import {cleanError} from './utils'

// Client for the JVC HM 660 camcoder

export default class JVCClient {
  constructor(url) {
    this._url = url
  }

  // Get Recorder status
  async recorderStatus(deviceID) {
    let response = null;
    try {
      response = await axios.get(`${this._url}/api/devices/${deviceID}/rec/status`)
    } catch(error) {
      throw cleanError(error)
    }
    return response.data
  }
}
