import Vue from 'vue'

const MAX_CHAT_HISTORY = 50;

const state = {
    chats: {},
}

const mutations = {
    addMessage(state, {id, params}) {
        let chat = state.chats[id];
        if(!chat) {
            chat = [];
        }

        // Prevent duplication if there are multiple listeners
        for(let msg of chat) {
            if(msg.id == paramd.id) {
                return
            }
        }

        chat.unshift(params);
        if(chat.length >= MAX_CHAT_HISTORY) {
            chat.slice(0, MAX_CHAT_HISTORY);
        }
        Vue.set(state.chats, id, chat);
    }
}

const getters = {}

const actions = {}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
}