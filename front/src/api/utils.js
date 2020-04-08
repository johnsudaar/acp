export function cleanError(e) {
  if(e.response && e.response.data && e.response.data.error) {
    return e.response.data.error
  }
  return e
}
