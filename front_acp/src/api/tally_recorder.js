import axios from 'axios'
import {cleanError} from './utils'

// Client for the tally recorder API
//
export default class TallyRecorderClient {
  constructor(url) {
    this._url = url;
  }

  // Search tally events
  async search(deviceID) {
    let response = [];
    try {
      response = await axios.get(`${this._url}/api/devices/${deviceID}/search`)
    } catch(error) {
      throw cleanError(error)
    }
    return response.data
  }
}
