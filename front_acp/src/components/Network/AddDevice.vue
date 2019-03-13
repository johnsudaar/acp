<template>
  <div>
    <v-menu offset-y>
      <v-btn slot="activator" color="primary" dark>
      Add Device
      </v-btn>
      <v-list>
        <v-list-tile
          v-for="(type) in $store.state.devices.types"
          @click="chooseDeviceType(type)"
        >
          <v-list-tile-title>{{ type }}</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-menu>
    <network-add-device-modal
      v-bind:device-type="currentDeviceType"
      v-bind:activated="showModal"
      v-on:close="closeModal"
      v-on:add-device="addDevice"
      />
  </div>
</template>

<script>
import deviceFilters from '../../filters/deviceFilters'

export default {
  data() {
    return {
      showModal: false,
      currentDeviceType: null,
    }
  },
  created() {
    this.$store.dispatch('devices/refreshTypes')
  },
  filters: deviceFilters,
  methods: {
    // This is called when a device type has been choosed from the bottom sheet
    chooseDeviceType(device) {
      this.currentDeviceType = device
      this.showModal = true
    },
    // This is called when the Close button is clicked on the device modal
    closeModal() {
      this.showModal = false
    },
    addDevice(id) {
      // Send the add device event to the parent
      this.$emit('add-device', id)
    }
  }
}
</script>
