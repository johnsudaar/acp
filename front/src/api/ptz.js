import axios from 'axios'

import {cleanError} from './utils'

export default class PTZClient {
  constructor(url) {
    this._url = url
  }

  joystick(camId, params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/devices/${camId}/ptz/joystick`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  position(camId, params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/devices/${camId}/ptz/position`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  storePosition(camId, params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/devices/${camId}/ptz/store`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  updatePosition(camId, positionId, params) {
    return new Promise((resolve, reject) => {
      axios.put(`${this._url}/api/devices/${camId}/ptz/store/${positionId}`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  deletePosition(camId, positionId) {
    return new Promise((resolve, reject) => {
      axios.delete(`${this._url}/api/devices/${camId}/ptz/store/${positionId}`)
        .then(() => {
          resolve()
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  positionsFor(camId) {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/devices/${camId}/ptz/store`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }
}
