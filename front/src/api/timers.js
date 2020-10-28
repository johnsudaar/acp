import axios from 'axios'
import {cleanError} from './utils'

// Timers subsystem of the ACP API
export default class TimersClient {
  constructor(url) {
    this._url = url;
  }

  // All timers currently loaded
  all() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/timers`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }

  // Create a new timer and return it's representation in the database
  create(params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/timers`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  update(id, params) {
    return new Promise((resolve, reject) => {
      axios.put(`${this._url}/api/timers/${id}`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  // Destroy a timer
  destroy(id) {
    return axios.delete(`${this._url}/api/timers/${id}`)
  }

  action(id, name, param) {
    let payload = {
      action: name,
      param: param,
    }
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/timers/${id}/action`, payload)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  deviceSources(id) {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/devices/${id}/timers/sources`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }
}
