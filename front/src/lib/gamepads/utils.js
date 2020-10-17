import {toPolar, toCarthesian} from '@/lib/maths/coordinates'
export function normalizeAxis(axis, speed, deadzone) {
  //return {
  //  x: transform(axis.x),
  //  y: transform(axis.y)
  //}
  let p = toPolar(axis);
  if(p.r < deadzone) {
    p.r = 0;
  } else {
    //p.r = (p.r - deadzone) / (1 - deadzone)
    p.r *= speed;
    p.r = Math.pow(p.r, 1.5)
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
