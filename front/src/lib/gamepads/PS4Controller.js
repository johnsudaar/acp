import {normalizeAxis} from './utils'

const LB = 4;
const LT = 6;
const RT = 7;
const RB = 5;
const DEADZONE = 0.25;

export default class XBoxController {
  toActions(gamepad) {
    let camAxis = normalizeAxis({x: gamepad.axes[0], y: gamepad.axes[1]}, DEADZONE);
    camAxis.y = -1 * camAxis.y;

    let zoomOut = gamepad.buttons[LT].value;
    let zoomIn = gamepad.buttons[RT].value;
    let zoom = zoomIn - zoomOut;

    let otherAxis = normalizeAxis({x: gamepad.axes[2], y: gamepad.axes[3], DEADZONE})

    let diaphOpen = gamepad.buttons(LB).value;
    let diaphClose = gamepad.buttons(RB).value;
    let diaph = diaphOpen - diaphClose;

    return {id: gamepad.index, name: gamepad.id, cam: camAxis, zoom: zoom}
  }
}
