export default {
  deviceTypeIcon(type) {
    switch(type){
      case 'ATEM':
        return "view_comfy"
      case 'JVC_HM_660':
        return "videocam"
      default:
        return "setting"
    }
  }
}
