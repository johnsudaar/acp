<template>
  <div>
    <v-menu offset-y>
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" color="primary" dark>
          Add Device
        </v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="(type) in $store.state.devices.types"
          @click="chooseDeviceType(type)"
        >
          <v-list-item-title>{{ type }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
    <network-add-device-modal
      :v-if="showModal"
      :device-type="currentDeviceType"
      :activated="showModal"
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
    async chooseDeviceType(device) {
      // Reset device form
      this.currentDeviceType = null
      await this.$nextTick()
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
