import axios from 'axios'
import {cleanError} from './utils'

// Scenes subsystem of the ACP API
export default class ScenesClient {
  constructor(url) {
    this._url = url;
  }

  // All scenes currently loaded
  all() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/scenes`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }

  // Create a new scenes and return it's representation in the database
  create(params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/scenes`, params)
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
      axios.put(`${this._url}/api/scenes/${id}`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  // Destroy a scene
  destroy(id) {
    return axios.delete(`${this._url}/api/scenes/${id}`)
  }

  launch(id) {
    return axios.post(`${this._url}/api/scenes/${id}/launch`)
  }

  active() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/scenes/_active`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  show(id) {
   return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/scenes/${id}`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }
}
