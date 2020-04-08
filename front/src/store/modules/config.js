import API from '../../api'

const state = {
  ip: null,
  port: null,
  apiClient: null,
  connected: null,
  version: null,
  updateAvailable: null,
  nextRelease: null,
}

const mutations = {
  setServerEndpoint(state, {ip, port}) {
    state.ip = ip
    state.port = port
    state.connected = false
    state.apiClient = null
  },

  setClient(state, client) {
    state.apiClient = client
    state.connected = client !== null
  },

  setVersion(state, {version, updateAvailable, nextRelease}) {
    state.version = version
    state.updateAvailable = updateAvailable
    state.nextRelease = nextRelease
  },
}

const actions = {
  save(context) {
    // Save current configuration to local storage
    localStorage.setItem('config/serverEndpoint/ip', context.state.ip)
    localStorage.setItem('config/serverEndpoint/port', context.state.port)
  },

  async load(context) {
    // load current configuration from local storage
    let ip = localStorage.getItem('config/serverEndpoint/ip')
    let port = localStorage.getItem('config/serverEndpoint/port')
    context.commit('setServerEndpoint', {ip, port})
    if(ip == null || port == null) {
      return
    }
    let api = new API(ip, port)
    try {
      // Try to ping
      await api.ping()
    } catch(e) {
      return
    }
    context.dispatch('connected', api)
  },

  async loadServerVersion(context) {
    let version = await context.state.apiClient.version()
    context.commit('setVersion', {
      version:         version.version,
      updateAvailable: version.update_available,
      nextRelease:     version.next_release,
    })
  },

  connected(context, client) {
    context.commit('setClient', client)
    if(client != null) {
      context.dispatch('devices/refresh', null, {root: true})
      context.dispatch('loadServerVersion')
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
}
