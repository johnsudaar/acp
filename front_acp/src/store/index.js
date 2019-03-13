import Vue from 'vue'
import Vuex from 'vuex'

import config from './modules/config'
import devices from './modules/devices'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    config,
    devices
  }
})
