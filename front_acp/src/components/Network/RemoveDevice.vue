<template>
  <v-dialog v-model="modalOpen" persistent max-width="290" :disabled="noDeviceSelected">
    <v-btn
      slot="activator"
      color="error"
     :disabled="noDeviceSelected"
     >
      Remove Device
    </v-btn>

    <v-card>
      <v-card-title class="headline">Are you sure? </v-card-title>
      <v-card-text>This will delete a device and all his links. Continue?</v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="red darken-1" flat @click="modalOpen = false">No </v-btn>
        <v-btn color="green darken-1" flat @click="removeDevice">Yes </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ['device'],
  data() {
    return {
      modalOpen: false,
    }
  },
  computed: {
    noDeviceSelected() {
      return this.device === null
    }
  },
  methods: {
    async removeDevice() {
      await this.$store.dispatch('devices/destroy', this.device)
      this.$emit('delete-device', this.device)
      this.modalOpen = false
    }
  }
}
</script>
