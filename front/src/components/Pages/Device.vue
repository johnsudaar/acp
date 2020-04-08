<script>
import devices from '../../api/devices'

export default {
  watch: {
    '$route.params.id': function (id) {
      this.resetComponent()
    },
    '$store.state.config.apiClient': function() {
      if(this.device != null) return
      this.resetComponent()
    }
  },
  created() {
    this.resetComponent()
  },
  methods: {
    // Reset all values to their default and load a new device.
    async resetComponent() {
      this.device = null;
      let client = this.$store.state.config.apiClient
      if(client == null) {
        return
      }
      this.device = await client.devices.get(this.$route.params.id)
    }
  },

  data() {
    return {
      device: null,
    }
  },
  computed: {
    showToolbar() {
      if(this.device == null) {
        return false
      }

      switch(this.device.type) {
        case "JVC_REMOTE":
          return false
      }

      return true
    },
    deviceComponent() {
      // Dynamically load the component name from the device type.
      let deviceType = this.device.type.toLowerCase();
      deviceType = deviceType.replace(/_/g, '-');
      return `device-${deviceType}`
    }
  }
}
</script>

<style scoped>
p {
  margin-bottom: 0;
}

</style>


<template>
  <v-layout fill-height align-center justify-center>
    <v-layout v-if="device" align-center justify-start column fill-height>
      <v-toolbar dense class="mb-4" v-if="showToolbar">
        <v-toolbar-title> Device {{device.name}}</v-toolbar-title>
        <v-spacer/>
          <p class="mr-2"> IP: {{device.ip}}</p>
          <p class="mr-2"> Port: {{device.port}}</p>
          <p> Status: {{device.state}}</p>
      </v-toolbar>
      <component :is="deviceComponent" :device-id="device.id"></component>
    </v-layout>
    <loading v-else/>
  </v-layout>
</template>
