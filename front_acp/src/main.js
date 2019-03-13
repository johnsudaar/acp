import 'normalize.css'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'jointjs/dist/joint.css'

import Vue from 'vue'
import App from './App.vue'
import router from './routes'
import Vuetify from 'vuetify'


// ------ Components -------
import NavigationLink from './components/NavigationLink.vue'
import Loading from './components/Misc/Loading.vue'
import SevenSeg from './components/Displays/SevenSeg.vue'
import Clock from './components/Misc/Clock.vue'
import ServerStatus from './components/Misc/ServerStatus.vue'
import Group from './components/Displays/Group.vue'
import LED from './components/Displays/LED.vue'
import AtemBtn from './components/Displays/Button.vue'
Vue.component('navigation-link', NavigationLink)
Vue.component('loading', Loading)
Vue.component('seven-seg', SevenSeg)
Vue.component('clock', Clock)
Vue.component('server-status', ServerStatus)
Vue.component('group', Group)
Vue.component('led', LED)
Vue.component('atem-btn', AtemBtn)

// Network
import NetworkAddDevice from './components/Network/AddDevice.vue'
import NetworkAddDeviceModal from './components/Network/AddDeviceModal.vue'
import NetworkRemoveDevice from './components/Network/RemoveDevice.vue'
Vue.component('network-add-device', NetworkAddDevice)
Vue.component('network-add-device-modal', NetworkAddDeviceModal)
Vue.component('network-remove-device', NetworkRemoveDevice)

// Network Forms
import NetworkFormAtem from './components/Network/forms/Atem.vue'
import NetworkFormJVCHM660 from './components/Network/forms/JVCHM660.vue'
import NetworkFormTallyRasp from './components/Network/forms/TallyRasp.vue'
import NetworkFormTallyRecorder from './components/Network/forms/TallyRecorder.vue'
Vue.component('network-form-atem', NetworkFormAtem)
Vue.component('network-form-jvc-hm-660', NetworkFormJVCHM660)
Vue.component('network-form-tally-rasp', NetworkFormTallyRasp)
Vue.component('network-form-tally-rec', NetworkFormTallyRecorder)

// Config
import ServerConfig from './components/Config/Server.vue'
Vue.component('config-server', ServerConfig)

// Atem
import AtemBtnGrp from './components/Atem/ButtonGroup.vue'
Vue.component('atem-btn-grp', AtemBtnGrp)

// Devices
import DeviceAtem from './components/Devices/Atem.vue'
import DeviceJVCHM660 from './components/Devices/JVCHM660.vue'
Vue.component('device-atem', DeviceAtem)
Vue.component('device-jvc-hm-660', DeviceJVCHM660)


// Form validation
import VeeValidate from 'vee-validate'
Vue.use(VeeValidate)

// Store
import store from './store'

// View
Vue.use(Vuetify)

Vue.config.productionTip = false

window.$ = require('jquery');
window.joint = require('jointjs');

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
