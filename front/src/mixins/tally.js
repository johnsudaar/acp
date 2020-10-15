export default {
  data() {
    return {
      tallies: {},
      tallyConnection: null,
    }
  },
  mounted() {
    console.log("Tally Mixin loaded");
    this.tallyConnection = this.$store.state.config.apiClient.realtime.subscribe("tally", this.onTallyMessage)
  },

  beforeDestroy() {
    if(this.tallyConnection) {
      this.tallyConnection.unsubscribe();
    }
  },
  methods: {
    onTallyMessage(message) {
      message = message.data
      this.$set(this.tallies, message.sender_id, message.data)
    },
    isPreview(deviceId) {
      if(this.tallies[deviceId]) {
        return this.tallies[deviceId].preview
      }
      return false;
    },
    isProgram(deviceId) {
      if(this.tallies[deviceId]) {
        return this.tallies[deviceId].program
      }
      return false;
    }
  }
}
