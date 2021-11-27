import Vue from 'vue'

const state = {
  scenes: {},
}

const mutations = {
  addScene(state, scene) {
    Vue.set(state.scenes, scene.id, scene);
  },

  updateScene(state, {id, params}) {
    let scene = state.scenes[id];
    if(!scene) {
      return
    }
    for(let k in params) {
      Vue.set(scene, k, params[k]);
    }
  },

  removeScene(state, id) {
    Vue.delete(state.scenes, id)
  }
}

const getters = {
}

const actions = {
  async create(context, params) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }

    let scene = await apiClient.scenes.create(params)
    context.commit('addScene', scene)
    return scene
  },

  async destroy(context, id) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    await apiClient.scenes.destroy(id)
    context.commit('removeScene', id)
  },

  async load(context) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    let scenes = await apiClient.scenes.all();

    if(scenes === null) {
      return
    }

    for(let scene of scenes) {
      context.commit('addScene', scene)
    }
  },

  async updateScene(context, {id, params}) {
    let {connected, apiClient} = context.rootState.config
    if(!connected) {
      throw "Not connected"
    }
    await apiClient.scenes.update(id, params)
    context.commit('updateScene', {id: id, params: params})
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
