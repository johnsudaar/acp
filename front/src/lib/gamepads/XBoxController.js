import {normalizeAxis} from './utils'

const LB = 4;
const LT = 6;
const RT = 7;
const RB = 5;
const X = 2;
const Y = 3;
const A = 0;
const B = 1;
const START = 9;
const RAxisButton = 11;

const DEADZONE = 0.01;

export default class XBoxController {
  constructor() {
    this.leftHandedMode = false
  }

  toActions(gamepad) {
    let camAxis = normalizeAxis({x: gamepad.axes[0], y: gamepad.axes[1]}, DEADZONE);
    let foxcusAxis = normalizeAxis({x: gamepad.axes[2], y: gamepad.axes[3]}, DEADZONE)
    if(this.leftHandedMode) {
      let temp = foxcusAxis
      foxcusAxis = camAxis
      camAxis = temp
    }

    camAxis.y = -1 * camAxis.y;

    let zoomOut = gamepad.buttons[LT].value;
    let zoomIn = gamepad.buttons[RT].value;
    let zoom = zoomIn - zoomOut;
    if(gamepad.buttons[START].pressed) {
      this.leftHandedMode = true
    }

    let buttons = {
      focus_push_auto: gamepad.buttons[RAxisButton].pressed,
      iris_close: gamepad.buttons[LB].pressed,
      iris_open:  gamepad.buttons[RB].pressed,
      button_1:  gamepad.buttons[X].pressed,
      button_2:  gamepad.buttons[Y].pressed,
      button_3:  gamepad.buttons[A].pressed,
      button_4:  gamepad.buttons[B].pressed,
    }

    return {id: gamepad.index, name: gamepad.id, cam: camAxis, zoom: zoom, focus: -foxcusAxis.y, buttons: buttons}
  }
}
