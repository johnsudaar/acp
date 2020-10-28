export default {
  data() {
    return {
      timersConnection: null,
    }
  },
  mounted() {
    console.log("Timers Mixin loaded");
    this.timersConnection = this.$store.state.config.apiClient.realtime.subscribe("timer", this.onTimerMessage)
  },

  beforeDestroy() {
    if(this.timersConnection) {
      this.timersConnection.unsubscribe();
    }
  },
  methods: {
    onTimerMessage(message) {
      let payload = message.data;
      this.$store.commit('timers/updateTimer', {
        id: payload.sender_id,
        params: payload.data,
      })
    },
  }
}
