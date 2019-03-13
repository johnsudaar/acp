import axios from 'axios'
import {cleanError} from './utils'

// Link subsystem of the ACP API
export default class LinkClient {
  constructor(url) {
    this._url = url;
  }

  // All links currently loaded
  all() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/links`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }

  // Create a new link and return it's representation in the database
  create(input, output) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/links`, {input, output})
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  // Destroy a link
  destroy(id) {
    return axios.delete(`${this._url}/api/links/${id}`)
  }
}
