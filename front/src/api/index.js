import DeviceClient from './devices'
import LinkClient from './links'
import TallyRecorderClient from './tally_recorder'
import JVCClient from './jvc'
import PTZClient from './ptz'
import SwitcherClient from './switcher'
import RealtimeClient from './realtime'
import TimersClient from './timers'
import ScenesClient from './scenes'
import {cleanError} from './utils'
import axios from 'axios'
import PositionGroupsClient from './position_groups'

// ACP API client
export default class Client {
  constructor(ip, port) {
    // Construct the base URL
    this._url = `http://${ip}:${port}`;
    if(port == 443) {
      this._url = `https://${ip}:${port}`;
    }

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

    // Init the timers client
    this.timers = new TimersClient(this._url)

    // Init the scene client
    this.scenes = new ScenesClient(this._url)

    // Init the position groups client
    this.positionGroups = new PositionGroupsClient(this._url);
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
