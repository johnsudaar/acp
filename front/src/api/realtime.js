import Centrifuge from 'centrifuge'

export default class RealtimeClient {
  constructor(ip, port) {
    this._ip = ip;
    this._port = port;
    this._ws = new Centrifuge(`ws://${this._ip}:${this._port}/connection/websocket/`)
    if(this._port == 443) {
      this._ws = new Centrifuge(`wss://${this._ip}:${this._port}/connection/websocket/`)
    }
  }

  on(event, cb) {
    this._ws.on(event, cb)
  }

  connect() {
    this._ws.connect()
  }

  subscribe(channel, cb) {
    return this._ws.subscribe(channel, cb)
  }

  ptzJoystick(camId, params) {
    this._ws.publish("ptz", {
      device_id: camId,
      data: params
    })
  }
}
