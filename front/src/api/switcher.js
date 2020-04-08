import axios from 'axios'

import {cleanError} from './utils'

export default class SwitcherClient {
  constructor(url) {
    this._url = url
  }

  switchOutput(id, output, input) {
    let params = {
      output: output,
      input: input,
    }
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/devices/${id}/switcher/switch`, params)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }
}
