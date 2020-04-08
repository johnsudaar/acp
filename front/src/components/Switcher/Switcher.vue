<template>
  <v-layout row>
    <atem-btn v-for="(input) in device.input_ports" :key="input" :name="input" big @click="changeInput(input)" :red="selectedInput == input"/>
  </v-layout>

</template>

<script>
export default {
  props: {
    device: Object,
  },
  data() {
    return {
      selectedInput: null,
    }
  },
  methods: {
    async changeInput(input) {
      try {
        await this.$store.state.config.apiClient.switcher.switchOutput(this.device.id, "PGM", input)
        this.selectedInput = input
      } catch(error) {
        console.log(error)
      }
    }
  }
}
</script>
