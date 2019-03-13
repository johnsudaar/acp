import {DeviceManager} from './device'

export class Paper {
  constructor(id, client) {
    this._id = id
    this._events = {}
    this._client = client

    this._graph = new joint.dia.Graph;
    this._paper = new joint.dia.Paper({
      el: document.getElementById(id),
      model: this._graph,
      width: '100%',
      height: '100%',
      gridSize: 1,
      defaultRouter: {
        name: 'metro',
      },
      background: {
        color: 'transparent',
      },
      validateConnection: this.noLoopback,
      snapLinks: { radius: 75 },
      linkPinning: false
    });

    this.addDragAndDropToPanAndTilt()
    this.itemMoveEventSender();
    this.highlighting();
    this.linking()

    // TODO: Save link vertices

    this.deviceManager = new DeviceManager(this)
  }

  get graph() {
    return this._graph
  }

  get paper() {
    return this._paper
  }


  noLoopback(cellViewS, magnetS, cellViewT, magnetT, end, linkView) {
    // Prevent linking from input ports.
    if (magnetS && magnetS.getAttribute('port-group') === 'in') return false;
    // Prevent linking from output ports to input ports within one element.
    if (cellViewS === cellViewT) return false;
    // Prevent linking to input ports.
    return magnetT && magnetT.getAttribute('port-group') === 'in';
  }

  addDragAndDropToPanAndTilt() {
    // Manage drag and drop to pan and tilt

    // Beginning position of the event
    let dragStartPosition = null;

    // When we click on a blanck space register the position
    this._paper.on('blank:pointerdown', (event, x, y) => {
      dragStartPosition = { x: x, y: y};
    });

    // Cleaning: free the dragStartPosition when the pointer is released
    this._paper.on('cell:pointerup blank:pointerup', (cellView, x, y) => {
      dragStartPosition = null;
    });

    $("#"+this._id).mousemove((event) => {
      // Since JointJS does not provide a cell:move event we use jQuery to get the mouse move and move the paper

      // If the start position has been set (we're clicking on a blank space)
      if (dragStartPosition) {
        // Translate the paer by this delta
        this._paper.translate(event.offsetX - dragStartPosition.x, event.offsetY - dragStartPosition.y)
      }
    });
  }

  on(event, callback) {
    // Subscribe to an event. This will just push the callback to the list of callbacks registered for this event.
    // If you want to send event see emit
    if(this._events[event] === undefined) {
      this._events[event] = []
    }
    this._events[event].push(callback)
  }

  emit(event, params) {
    // Emit a specific event
    // This will send the params to all the callbacks registered to this event
    let cbs = this._events[event]
    if(cbs === undefined) {
      return
    }

    for(let i = 0; i < cbs.length; i++) {
      cbs[i](params)
    }
  }

  itemMoveEventSender() {
    // Send an event where an item has been moved
    this._paper.on('cell:pointerup', (cellView, event, x, y) => {
      if(cellView.model.isLink()) {
        // If it's a link do not do anything
        return
      }
      this.emit('cell:move', {
        id: cellView.model.id,
        position: cellView.model.attributes.position,
      })
    })
  }

  highlighting() {
    let curCellView = null;
    this._paper.on('cell:pointerclick', (newCellView) => {
      if(newCellView.model.isLink()) {
        // if it's a link treat the event as a deselect event
        if(curCellView !== null) {
          curCellView.model.attr({".body": {fill: "white"}})
          curCellView = null
          this.emit('cell:selected', null)
        }
        return
      }

      if(curCellView === null) {
        // No cells where selected: Select the new one
        curCellView = newCellView
        curCellView.model.attr({".body": {fill: "#ffd7b2"}})
      } else if(curCellView.model.id === newCellView.model.id){
        // We clicked on the same cell: deselect it
        curCellView.model.attr({".body": {fill: "white"}})
        curCellView = null
      } else {
        // We clicked on another cell: Change the selected cell
        curCellView.model.attr({".body": {fill: "white"}})
        curCellView = newCellView
        curCellView.model.attr({".body": {fill: "#ffd7b2"}})
      }

      if(curCellView === null) {
        // If no cells where selected
        this.emit('cell:selected', null)
      } else {
        // If a cell is selected
        this.emit('cell:selected', curCellView.model.id)
      }
    })

    this._paper.on("blank:pointerclick", () => {
      // We clicked on the blang: Deselect the current cell
      if(curCellView !== null) {
        curCellView.model.attr({".body": {fill: "white"}})
        curCellView = null
        this.emit('cell:selected', null)
      }
    })
  }

  // Manage linking and unlinking events
  linking() {
    this._paper.on('link:connect', async (linkView, event, elementViewConnected, magnet, arrowHead) =>  {
      let output = {
        device_id: linkView.model.attributes.source.id,
        port: linkView.model.attributes.source.port,
      }
      let input = {
        device_id: linkView.model.attributes.target.id,
        port: linkView.model.attributes.target.port,
      }

      let link = null;
      try {
        link = await this._client.links.create(input, output)
      }catch(e) {
        this.emit('error', e)
        linkView.remove()
        return
      }
      linkView.model.attr({acp_id: link.id})
    })

    this._graph.on('remove', (cell) => {
      if(cell.isLink()) {
        let id = cell.attributes.attrs.acp_id;
        if(id === undefined) {
          return
        }
        this._client.links.destroy(id)
      }
    })
  }
}
