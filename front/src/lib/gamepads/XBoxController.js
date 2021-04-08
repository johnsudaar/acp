import {normalizeAxis} from './utils'

const LB = 4;
const LT = 6;
const RT = 7;
const RB = 5;
const X = 2;
const Y = 3;
const A = 0;
const B = 1;
const BACK = 8;
const START = 9;
const RAxisButton = 11;
const RIGHT = 15;
const LEFT = 14;
const DOWN = 13;
const UP = 12;


const DEADZONE = 0.16;

export default class XBoxController {
  constructor() {
    this.leftHandedMode = false;
    this.leftHandedPressed = false;
    this.jibMode = false;
    this.jibPressed = false;
    this.speed = 0.75;
  }

  toActions(gamepad) {
    let camAxis = normalizeAxis({x: gamepad.axes[0], y: gamepad.axes[1]}, this.speed, DEADZONE);
    let foxcusAxis = normalizeAxis({x: gamepad.axes[2], y: gamepad.axes[3]}, this.speed, DEADZONE)
    if(this.leftHandedMode) {
      let temp = foxcusAxis
      foxcusAxis = camAxis
      camAxis = temp
    }

    if(gamepad.buttons[UP].pressed) {
      camAxis.y = -1 * this.speed;
    }
    if(gamepad.buttons[DOWN].pressed) {
      camAxis.y = 1 * this.speed;
    }

    if(gamepad.buttons[LEFT].pressed)  {
      //camAxis.x = -1 * this.speed;
    }

    if(gamepad.buttons[RIGHT].pressed) {
      //camAxis.x = 1 * this.speed;
    }

    if(!this.jibMode) {
      // For normal operations
      camAxis.y *= -1;
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
      this.speed = .25;
    }

    if(gamepad.buttons[Y].pressed) {
      this.speed = .5;
    }
    if(gamepad.buttons[B].pressed) {
      this.speed = .75;
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
