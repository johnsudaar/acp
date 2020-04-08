import Vue from 'vue'
import Vuex from 'vuex'

import config from './modules/config'
import devices from './modules/devices'
import ptzpositions from './modules/ptzpositions'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    config,
    devices,
    ptzpositions
  }
})
