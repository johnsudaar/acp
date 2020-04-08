import {toPolar, toCarthesian} from '@/lib/maths/coordinates'
export function normalizeAxis(axis, deadzone) {
  let slowdown = 2;
  return {
    x: transform(axis.x),
    y: transform(axis.y)
  }
  if(!deadzone) {
    deadzone = 0.15
  }
  let p = toPolar(axis);
  if(p.r < deadzone && false) {
    p.r = 0;
  } else {
    //p.r = (p.r - deadzone) / (1 - deadzone)
    p.r = Math.pow(p.r, 1)
    p.r /= 5
  }
  return toCarthesian(p)
}

function transform(v) {
  let slowdown = 2.5;
  let r = Math.pow(Math.abs(v), 2)/slowdown
  if(v < 0) {
    return -r
  } else {
    return r
  }
}
