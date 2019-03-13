export class DeviceManager {
  constructor(paper) {
    this._paper = paper
    this._models = {}
  }

  delete(id) {
    if(this._models[id] === undefined) {
      return
    }
    this._models[id].remove()
    delete this._models[id]
  }
  async add(client, id) {
    let device = await client.devices.get(id)

    let height = Math.max(device.input_ports.length, device.output_ports.length) * 20;
    let position = null
    if(device.display_opts) {
      position = device.display_opts["position"]
    }

    if(position == null) {
      position = {x: 20, y: 20}
    }

    let shape = new joint.shapes.devs.Model({
      id: device.id,
      size: { width: 175, height: height},
      inPorts: device.input_ports,
      outPorts: device.output_ports,
      position: position,
      attrs: {
        '.label': {
          'ref-y': '-20px',
          fill: 'white',
          position: 'outside',
          text: device.name,
        },
      },
      ports: {
        groups: {
          'in': {
            attrs: {
              '.port-body': {
                r: 5,
              },
              '.port-label': {
                'font-size': '15px',
              }
            }
          },
          'out': {
            attrs: {
              '.port-body': {
                r: 5,
              },
              '.port-label': {
                'font-size': '15px',
              }

            }
          }
        }
      }
    })
    shape.changeInGroup({label: {position: 'inside'}})
    shape.changeOutGroup({label: {position: 'inside'}})

    shape.addTo(this._paper.graph)
    this._models[id] = shape
  }
}
