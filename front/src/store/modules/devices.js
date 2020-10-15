import Vue from 'vue'
const state = {
  devices: [],
  types: [],
}

const mutations = {
  addDevice(state, device) {
    state.devices.push(device)
  },

  setDevices(state, devices) {
    state.devices = devices
  },

  removeDevice(state, id) {
    // First find the index of this device
    let i = state.devices.map(device => device.id).indexOf(id)
    // If the item is not found
    if(i === -1) {
      return
    }

    // Remove it
    Vue.delete(state.devices, i)
  },

  setTypes(state, deviceTypes) {
    state.types = deviceTypes
  }
}

const actions = {
  async refresh(context) {
    let {connected, apiClient} = context.rootState.config;
    if(!connected) {
      throw "Not connected"
    }

    let devices = await apiClient.devices.all()
    context.commit('setDevices', devices)
  },

  async refreshTypes(context) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }

    let deviceTypes = await apiClient.devices.types()
    context.commit('setTypes', deviceTypes)
  },

  async create(context, {name, type, params}) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }

    let device = await apiClient.devices.create(name, type, params)
    context.commit('addDevice', device)
    return device
  },

  async destroy(context, id) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    await apiClient.devices.destroy(id)
    context.commit('removeDevice', id)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
}
