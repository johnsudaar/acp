import DeviceClient from './devices'
import LinkClient from './links'
import axios from 'axios'

// ACP API client
export default class Client {
  constructor(ip, port) {
    // Construct the base URL
    this._url = `http://${ip}:${port}`;

    // Init the devices subservice of this API
    this.devices = new DeviceClient(this._url);

    // Init the links subservice of this API
    this.links = new LinkClient(this._url)
  }

  // Ping method: used to check connection between the client and the API server
  ping() {
    return axios.get(`${this._url}/api/ping`, {
      timeout: 10000,
    })
  }
}

