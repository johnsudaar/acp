import axios from 'axios'
import cleanError from './utils'

export default class Intercom {
  constructor(url) {
    this._url = url
  }

  async createRoom(deviceID, room) {
    let response = null;
    try {
      response = await axios.post(`${this._url}/api/devices/${deviceID}/rooms`, room)
    } catch(error) {
      throw cleanError(error)
    }
    return response.data.room
  }

  async listRooms(deviceID) {
    let response = null;
    try {
      response = await axios.get(`${this._url}/api/devices/${deviceID}/rooms`)
    } catch(error) {
      throw cleanError(error)
    }
    return response.data.rooms
  }
}
