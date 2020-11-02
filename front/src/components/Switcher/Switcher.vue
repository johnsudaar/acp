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
  mounted() {
    window.addEventListener('keydown', this.onKeyDown);
  },
  beforeDestroy() {
    window.removeEventListener('keydown', this.onKeyDown);
  },
  methods: {
    onKeyDown(key) {
      key = key.keyCode
      if(key < 97 || key > 105) {
        return
      }
      let input = key - 97;
      this.changeInput(this.device.input_ports[input])
    },
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
