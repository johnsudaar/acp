import Vue from 'vue'
import Vuex from 'vuex'

import config from './modules/config'
import devices from './modules/devices'
import ptzpositions from './modules/ptzpositions'
import timers from './modules/timers'
import scenes from './modules/scenes'
import chat from './modules/chat'
import positiongroups from './modules/position_groups'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    config,
    devices,
    ptzpositions,
    timers,
    scenes,
    chat,
    positiongroups,
  }
})
