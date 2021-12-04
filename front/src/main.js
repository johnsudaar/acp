import 'normalize.css'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'
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
import Notifications from './components/Misc/Notifications.vue'
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
Vue.component('notifications', Notifications)

// Network
import NetworkAddDevice from './components/Network/AddDevice.vue'
import NetworkAddDeviceModal from './components/Network/AddDeviceModal.vue'
import NetworkRemoveDevice from './components/Network/RemoveDevice.vue'
Vue.component('network-add-device', NetworkAddDevice)
Vue.component('network-add-device-modal', NetworkAddDeviceModal)
Vue.component('network-remove-device', NetworkRemoveDevice)

// Config
import ServerConfig from './components/Config/Server.vue'
Vue.component('config-server', ServerConfig)

// Atem
import AtemBtnGrp from './components/Atem/ButtonGroup.vue'
Vue.component('atem-btn-grp', AtemBtnGrp)

// Recording
import RecStatus from './components/Rec/Status.vue'
Vue.component('rec-status', RecStatus)

// Devices
import DeviceAtem from './components/Devices/Atem.vue'
import DeviceJVCHM660 from './components/Devices/JVCHM660.vue'
import DeviceJVCRemote from './components/Devices/JVCRemote.vue'
import DeviceTallyRecorder from './components/Devices/TallyRecorder.vue'
Vue.component('device-atem', DeviceAtem)
Vue.component('device-jvc-hm-660', DeviceJVCHM660)
Vue.component('device-jvc-remote', DeviceJVCRemote)
Vue.component('device-tally-rec', DeviceTallyRecorder)

// Remote

import JVCRemoteInput from './components/Remote/JVCInput.vue'
Vue.component('jvc-remote-input', JVCRemoteInput)

// CCU Stuff
import CCU from './components/CCU/Main.vue'
import SingleCCU from './components/CCU/Single.vue'
import OverviewCCU from './components/CCU/Overview.vue'

Vue.component('ccu', CCU)
Vue.component('ccu-single', SingleCCU)
Vue.component('ccu-overview', OverviewCCU)

import Pad from './components/Gamepads/Pad'
Vue.component('pad', Pad)

// PTZ
import PtzAddPosition from './components/Ptz/AddPosition.vue'
import PtzEditPosition from './components/Ptz/EditPosition.vue'
import PtzDeletePosition from './components/Ptz/DeletePosition.vue'
import PtzPositions from './components/Ptz/Positions.vue'
import PtzEdit from './components/Ptz/Edit.vue'
import PtzForm from './components/Ptz/Form.vue'

Vue.component('ptz-add-position', PtzAddPosition)
Vue.component('ptz-edit-position', PtzEditPosition)
Vue.component('ptz-delete-position', PtzDeletePosition)
Vue.component('ptz-positions', PtzPositions)
Vue.component('ptz-edit', PtzEdit)
Vue.component('ptz-form', PtzForm)

// Switcher
import Switcher from './components/Switcher/Switcher.vue'
Vue.component('switcher', Switcher)

// Timers
import AddTimer from './components/Timers/AddTimer.vue'
import TimerForm from './components/Timers/Form.vue'
import TimerCard from './components/Timers/TimerCard.vue'
Vue.component('add-timer', AddTimer)
Vue.component('timer-form', TimerForm)
Vue.component('timer-card', TimerCard)

// Scenes
import SceneElemForm from './components/Scenes/ElemForm.vue'
import SceneView from './components/Scenes/View.vue'
import ChatWidget from './components/Scenes/ChatWidget.vue'
Vue.component('scene-elem-form', SceneElemForm)
Vue.component('scene-view', SceneView)
Vue.component('chat-widget', ChatWidget)

// Position Groups

import PositionGroupsAdd from './components/PositionGroups/Add'
import PositionGroupsDestroy from './components/PositionGroups/Destroy.vue'
Vue.component('position-groups-add', PositionGroupsAdd)
Vue.component('position-groups-destroy', PositionGroupsDestroy)

// Form validation
import VeeValidate from 'vee-validate'
Vue.use(VeeValidate)

// Resize text

import AutoText from 'vue-auto-text'
Vue.component(AutoText.name, AutoText)

// UUID
import UUID from "vue-uuid";
Vue.use(UUID);

// Store
import store from './store'

// View
Vue.use(Vuetify)

Vue.config.productionTip = false

window.$ = require('jquery');
window.joint = require('jointjs');

let vuetify = new Vuetify({
  theme: {
    dark: true,
  }
})

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
