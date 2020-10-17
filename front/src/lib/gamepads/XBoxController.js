import {normalizeAxis} from './utils'

const LB = 4;
const LT = 6;
const RT = 7;
const RB = 5;
const X = 2;
const Y = 3;
const A = 0;
const B = 1;
const BACK = 9;
const START = 8;
const RAxisButton = 11;

const DEADZONE = 0.16;

export default class XBoxController {
  constructor() {
    this.leftHandedMode = false;
    this.leftHandedPressed = false;
    this.jibMode = false;
    this.jibPressed = false;
    this.speed = 0.9;
  }

  toActions(gamepad) {
    let camAxis = normalizeAxis({x: gamepad.axes[0], y: gamepad.axes[1]}, this.speed, DEADZONE);
    let foxcusAxis = normalizeAxis({x: gamepad.axes[2], y: gamepad.axes[3]}, this.speed, DEADZONE)
    if(this.leftHandedMode) {
      let temp = foxcusAxis
      foxcusAxis = camAxis
      camAxis = temp
    }

    // For normal operations
    if(this.jibMode) {
      camAxis.x *= -1;
    } else {
      camAxis.y = -1 * camAxis.y;
    }

    let zoomOut = gamepad.buttons[LT].value;
    let zoomIn = gamepad.buttons[RT].value;
    let zoom = zoomIn - zoomOut;

    if(gamepad.buttons[START].pressed) {
      if(!this.leftHandedPressed) {
        this.leftHandedMode = !this.leftHandedMode
        this.leftHandedPressed = true;
        console.log(`leftHandedMode: ${this.leftHandedMode}`)
      }
    } else {
      this.leftHandedPressed = false;
    }

    if(gamepad.buttons[BACK].pressed) {
      if(!this.jibPressed) {
        this.jibMode = !this.jibMode;
        this.jibPressed = true;
        console.log(`jibMode: ${this.jibMode}`)
      }
    } else {
      this.jibPressed = false;
    }

    if(gamepad.buttons[X].pressed) {
      this.speed = .35;
    }

    if(gamepad.buttons[Y].pressed) {
      this.speed = .5;
    }
    if(gamepad.buttons[B].pressed) {
      this.speed = .9;
    }

    let buttons = {
      focus_push_auto: gamepad.buttons[RAxisButton].pressed,
      iris_close: gamepad.buttons[LB].pressed,
      iris_open:  gamepad.buttons[RB].pressed,
      //button_1:  gamepad.buttons[X].pressed,
      //button_2:  gamepad.buttons[Y].pressed,
      //button_3:  gamepad.buttons[A].pressed,
      //button_4:  gamepad.buttons[B].pressed,
    }

    return {id: gamepad.index, name: gamepad.id, cam: camAxis, zoom: zoom, focus: -foxcusAxis.y, buttons: buttons}
  }
}
