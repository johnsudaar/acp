<template>
  <v-layout class="row">
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
      // We want to check the key target
      // But if the target is not defined, we decide to abort
      if(!key || !key.target || !key.target.nodeName) {
        console.error("Error while processing key press: No target defined.")
        return
      }
      // If the user is in a form. Abort input processing
      if(key.target.nodeName === "INPUT") {
        return
      }
      key = key.keyCode
      // If the key is not a key a number.
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
