<template>
  <v-container grid-list-md fill-height>
    <v-layout row wrap>
      <v-flex xs3>
        <pad :pad="pads[0]" index="0"/>
      </v-flex>
      <v-flex xs3>
        <pad :pad="pads[1]" index="1"/>
      </v-flex>
      <v-flex xs3>
        <pad :pad="pads[2]" index="2"/>
      </v-flex>
      <v-flex xs3>
        <pad :pad="pads[3]" index="3"/>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import XBoxController from '@/lib/gamepads/XBoxController'
export default {
  data() {
    return {
      timer: null,
      pads: [
        null,
        null,
        null,
        null,
      ],
      controllers: [
        null,
        null,
        null,
        null
      ],
    }
  },
  mounted() {
    this.timer = setInterval(this.refreshGamepads, 20)
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  methods: {
    refreshGamepads() {
      this.pads = []
      let gamepads = navigator.getGamepads()
      for(let i = 0; i < gamepads.length; i ++) {
        let gamepad = gamepads[i];
        if(gamepad !== null) {
          let controller = this.controllers[i]
          if(controller === null) {
            console.log("NEW")
            this.controllers[i] = new XBoxController()
            controller = this.controllers[i]
          }
          let action = controller.toActions(gamepad)
          this.pads[action.id] = action;
        } else {
          this.controllers[i] = null
        }
      }
    },
  }
}
</script>
