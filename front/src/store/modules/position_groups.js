import Vue from 'vue'

const state = {
    groups: {},
}

const mutations = {
    addGroup(state, group) {
        Vue.set(state.groups, group.id, group);
    },

    removeGroup(state, id) {
        Vue.delete(state.groups, id);
    }
}

const getters = {
    asArray(state) {
        let res = [];
        for(let group in state.groups) {
            res.push(state.groups[group]);
        }
        return res;
    },
}

const actions = {
    async create(context, params) {
        let {connected, apiClient} = context.rootState.config
        if(!connected) {
            throw "Not Connected"
        }

        let group = await apiClient.positionGroups.create(params);
        context.commit('addGroup', group);
        return group;
    },

    async destroy(context, {id, destroyMembers}) {
        let {connected, apiClient} = context.rootState.config
        if(!connected) {
            throw "Not Connected"
        }
        await apiClient.positionGroups.destroy(id, destroyMembers)
        context.commit('removeGroup', id)
    },

    async load(context) {
        let {connected, apiClient} = context.rootState.config
        if(!connected) {
            throw "Not Connected"
        }

        let groups = await apiClient.positionGroups.all();
        if(groups === null) {
            return
        }

        for(let group of groups) {
            context.commit('addGroup', group)
        }
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
};
