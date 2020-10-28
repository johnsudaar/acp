import Vue from 'vue'

const state = {
  timers: {},
}

const mutations = {
  addTimer(state, timer) {
    Vue.set(state.timers, timer.id, timer);
  },

  updateTimer(state, {id, params}) {
    let timer = state.timers[id];
    if(!timer) {
      return
    }
    for(let k in params) {
      Vue.set(timer, k, params[k]);
    }
  },

  removeTimer(state, id) {
    Vue.delete(state.timers, id)
  }
}

const getters = {
  timerValue: (state) => (id) => {
    let timer = state.timers[id]
    if(!timer) {
      return "-E0:E0:E0"
    }
    return timer.value;
  }
}

const actions = {
  async create(context, params) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }

    let timer = await apiClient.timers.create(params)
    context.commit('addTimer', timer)
    return timer
  },

  async destroy(context, id) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    await apiClient.timers.destroy(id)
    context.commit('removeTimer', id)
  },

  async load(context) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    let timers = await apiClient.timers.all();

    if(timers === null) {
      return
    }

    for(let timer of timers) {
      context.commit('addTimer', timer)
    }
  },

  async updateTimer(context, {id, params}) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    await apiClient.timers.update(id, params)
    context.commit('updateTimer', {id: id, params: params})
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
