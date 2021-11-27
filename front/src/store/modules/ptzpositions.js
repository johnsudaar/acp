import Vue from 'vue'

const state = {
  positions: {},
}

const mutations = {
  addPosition(state, {cam, position}) {
    let positions = state.positions[cam]
    if(!positions) {
      positions = []
    }
    positions.push(position)
    Vue.set(state.positions, cam, positions)
  },
  setPositions(state, {cam, positions}) {
    Vue.set(state.positions, cam, positions)
  },
  removePosition(state, {cam, id}) {
    if(!state.positions[cam]) {
      return
    }
    let i = state.positions[cam].map(pos => pos.id).indexOf(id)
    if(i === -1) {
      return
    }
    Vue.delete(state.positions[cam], i)
  },
  setPosition(state, {cam, position}) {
    let positions = state.positions[cam]
    if(!positions) {
      positions = []
      return
    }
    let i = positions.map(pos => pos.id).indexOf(position.id)
    // If the item is not found
    if(i === -1) {
      positions.push(position)
    } else {
      Vue.set(positions, i, position)
    }
    Vue.set(state.positions, cam, positions)
  }
}

const getters = {
  forDevice: (state) => (cam) => {
    return state.positions[cam]
  },

  find: (state) => (cam, id) => {
    let positions = state.positions[cam]
    if(!positions) {
      positions = []
    }

    return positions.find((elem) => {
      return elem.id == id
    })
  }
}

const actions = {
  async ensure(context, cam) {
    if(state.positions[cam]) {
      return
    }
    await context.dispatch('refresh', cam);
  },
  async refresh(context, cam) {
    let {connected, apiClient} = context.rootState.config;
    if(!connected) {
      throw "Not connected"
    }

    let positions = await apiClient.ptz.positionsFor(cam)
    context.commit('setPositions', {cam, positions})
  }
}

export default {
  namespaced: true,
  actions,
  state,
  getters,
  mutations,
}
