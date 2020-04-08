<template>
  <v-container>
    <group v-for="(device) in ptzDevices" :key="device.id" :title="device.name" class="col-12 mt-3" fill-width>
      <v-layout row>
        <ptz-positions :device="device" />
        <v-spacer/>
        <ptz-edit :device="device"/>
      </v-layout>
    </group>

    <group v-for="(device) in switchers" :key="device.id" :title="device.name" class="col-12 mt-3" fill-width>
      <v-layout row>
        <switcher :device="device"/>
      </v-layout>
    </group>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      fab: false,
    }
  },
  computed: {
    ptzDevices() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("ptz")
      }).sort((a, b) => {
        return a.name > b.name ? 1 : -1
      })
    },
    switchers() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("switcher")
      }).sort((a, b) => {
        return a.name > b.name ? 1 : -1
      })
    }
  }
}
</script>
