<template>
  <v-layout fill-width>
    <div v-for="(device) in recordingDevices" class="mr-2">
      <group :title="device.name">
        <rec-status :device-id="device.id"/>
      </group>
    </div>
  </v-layout>
</template>

<script>

export default {
  computed: {
    recordingDevices() {
      return this.$store.state.devices.devices.filter((device) => {
        return ["JVC_HM_660"].includes(device.type)
      })
    }
  }
}
</script>
