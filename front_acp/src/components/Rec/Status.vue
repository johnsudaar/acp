<template>
  <v-layout align-end>
    <atem-btn name="Rec" v-bind:red="recStatus.recording"/>
    <seven-seg color="red" :text="recStatus.recording_time | timeFilter"/>
  </v-layout>
</template>

<script>

export default {
  props: {
    deviceId: String
  },
  data() {
    return {
      intervalID: null,
      recStatus: {
        recording: false,
        recording_time: "00:00:00",
      },
    }
  },
  methods: {
    startClock() {
      clearInterval(this.intervalID)
      this.intervalID = setInterval(() => {
        this.refreshStatus()
      }, 1000)
    },
    async refreshStatus() {
      if(!this.$store.state.config.connected) {
        return
      }

      let client = this.$store.state.config.apiClient
      try {
        this.recStatus = await client.jvc.recorderStatus(this.deviceId)
      } catch(error) {
        console.error(error)
      }
    }
  },
  created() {
    this.startClock()
  },
  beforeDestroy() {
    clearInterval(this.intervalID)
  },
  filters: {
    timeFilter: (a)=> {
      return a.replace(/[msh]/gi, ":")
    }
  }
}

</script>
