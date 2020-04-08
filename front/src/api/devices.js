import axios from 'axios'
import {cleanError} from './utils'

// Device subsystem of the ACP API
export default class DeviceClient {
  constructor(url) {
    this._url = url;
  }

  // All get all devices currently loaded
  all() {
    return new Promise((resolve, reject) => {
      // We need to intercept and add or own promise to convert the raw json into a Device Array object
      axios.get(`${this._url}/api/devices`)
        .then((response) => {
          let devices = [];
          for (var id in response.data) {
            devices.push(new Device(id, response.data[id]))
          }
          resolve(devices)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  // Get (id: string) get a single device from the graph
  get(id) {
    // We need to intercept and add or own promise to convert the raw json into a Device object
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/devices/${id}`)
        .then((response) =>{
          resolve(new Device(id, response.data))
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  types() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/device_types`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  typeParams(name) {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/device_types/${name}/params`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }

  create(name, type, params) {
    return new Promise((resolve, reject) => {
      axios.post(`${this._url}/api/devices`, {
        type,
        name,
        params
      }).then((response) => {
        resolve(new Device(response.data["id"], response.data))
      }).catch((error) => {
        reject(cleanError(error))
      })
    })
  }

  update(id, params) {
    return axios.put(`${this._url}/api/devices/${id}`, params)
  }

  destroy(id) {
    return axios.delete(`${this._url}/api/devices/${id}`)
  }
}

class Device {
  constructor(id, data) {
    this.id = id;
    for(let key in data) {
      this[key] = data[key];
    }

    // If we fetch a single device, json will be:
    // {
    //  "type": "device_type",
    //  "device": {...},
    // }
    //
    // We want to flatten that to make it easier to use
    for(let key in this.device) {
      this[key] = this.device[key]
    }
  }

  // Front end path to this device
  path() {
    return `/device/${this.id}`
  }
}
