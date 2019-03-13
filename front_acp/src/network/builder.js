import {Paper} from './paper'
import {addLink} from './links'

export default {
  async build(id, client) {
    let paper = new Paper(id, client)

    let devices = await client.devices.all()
    let promises = []
    devices = devices || []
    for(let i in devices) {
      promises.push(paper.deviceManager.add(client, devices[i].id))
    }
    await Promise.all(promises)

    let links = await client.links.all()
    for(let i in links) {
      addLink(paper.graph, links[i].id, links[i].input, links[i].output)
    }
    return paper
  }
}
