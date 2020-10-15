import DeviceClient from './devices'
import LinkClient from './links'
import TallyRecorderClient from './tally_recorder'
import JVCClient from './jvc'
import PTZClient from './ptz'
import SwitcherClient from './switcher'
import RealtimeClient from './realtime'
import {cleanError} from './utils'
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

    // Init the tally recorder client
    this.tally_recorder = new TallyRecorderClient(this._url)

    // Init the JVC client
    this.jvc = new JVCClient(this._url)

    // Init the PTZ client
    this.ptz = new PTZClient(this._url)

    // Init the Switcher client
    this.switcher = new SwitcherClient(this._url)

    // Init the realtime client
    this.realtime = new RealtimeClient(ip, port)
  }

  // Ping method: used to check connection between the client and the API server
  ping() {
    return axios.get(`${this._url}/api/ping`, {
      timeout: 10000,
    })
  }

  version() {
    return new Promise((resolve, reject) => {
      axios.get(`${this._url}/api/version`)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(cleanError(error))
        })
    })
  }
}
