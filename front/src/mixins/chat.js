export default {
    data() {
        return {
            chatConnection: null,
        }
    },
    mounted() {
        this.chatConnection = this.$store.state.config.apiClient.realtime.subscribe("chat", this.onChatMessage);
    },
    beforeDestroy() {
        if(this.chatConnection) {
            this.chatConnection.unsubscribe();
        }
    },
    methods: {
        onChatMessage(message) {
            console.log(message)
            let payload = message.data;
            this.$store.commit('chat/addMessage', {
                id: payload.sender_id,
                params: payload.data,
            });
        }
    }
}