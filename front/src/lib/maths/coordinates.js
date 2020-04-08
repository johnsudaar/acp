export function toPolar(v) {
  let r = Math.sqrt(Math.pow(v.x, 2) + Math.pow(v.y, 2))
  let t = Math.atan(v.y/v.x)
  if(v.x < 0) {
    t += Math.PI
  }
  return {r: r, theta: t}
}

export function toCarthesian(p) {
  let x = p.r * Math.cos(p.theta)
  let y = p.r * Math.sin(p.theta)
  return  {x: x, y: y}
}
