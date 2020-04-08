<template>
  <v-layout fill-height align-center justify-center>
    <v-layout column v-if="$store.state.config.connected" fill-height>
      <group title="Tools" fill-width>
        <v-layout>
          <network-add-device v-on:add-device="addDevice"/>
          <v-spacer />
          <network-remove-device :device="device" v-on:delete-device="deleteDevice"/>
        </v-layout>
      </group>
      <div id="graph_container">
      </div>
    </v-layout>
    <loading v-else/>

    <v-dialog
      v-model="openErrorModal"
      width="500"
    >
      <v-card>
        <v-card-title class="headline red lighten-2" primary-title>
          Error
        </v-card-title>
        <v-card-text>
          {{errorModalText}}
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" flat @click="openErrorModal= false">
            Ok
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
import network from '../../network/builder'
import Vue from 'vue'

export default {
  data() {
    return {
      paper: null,
      device: null,
      openErrorModal: false,
      errorModalText: "",
    }
  },
  mounted() {
    this.initGraph()
  },
  watch: {
    '$store.state.config.connected': async function() {
      // If we're connected now
      if(this.$store.state.config.connected) {
        // Wait for the next tick (wait for template evaluation)
        await Vue.nextTick()
        // Init graph
        this.initGraph()
      }
    }
  },
  computed: {
    noDeviceSelected() {
      return this.device === null
    }
  },
  methods: {
    async initGraph() {
      if(!this.$store.state.config.connected) {
        return
      }
      let paper = await network.build(
        "graph_container",
        this.$store.state.config.apiClient,
      )

      paper.on('cell:move', this.onDeviceMove)
      paper.on('cell:selected', this.onDeviceSelected)
      paper.on('error', this.showError)
      this.paper = paper
    },
    onDeviceMove({id, position}) {
      let client = this.$store.state.config.apiClient;
      client.devices.update(id, {"display_opts": {position}})
    },
    onDeviceSelected(id) {
      this.device = id
    },
    addDevice(id) {
      this.paper.deviceManager.add(this.$store.state.config.apiClient, id)
    },
    deleteDevice(id) {
      this.paper.deviceManager.delete(id)
    },
    showError(error) {
      this.errorModalText = error.toString()
      this.openErrorModal = true
    }
  }
}
</script>

<style scoped>
#graph_container {
  margin-top: 10px;
  width: 100%;
  height: 100%;
}

.toolbar {
  width: 100%;
}
</style>
