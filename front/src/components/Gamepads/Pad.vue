<template>
  <v-container>
    <v-layout row>
      <led :green="pad != null" class="center"/>
    </v-layout>
    <v-layout v-if="pad" class="margin-top">
      <v-flex xs12>
        <v-select label="Camera" :items="items" item-text="name" item-value="id" v-on:change="onSelect"/>
      </v-flex>
      <br/>
      <v-flex xs12>
      {{pad.cam.x}}
      {{pad.cam.y}}
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  props: ['pad'],
  data() {
    return {
      camId: null,
      timer: null,
      isSending: false,
      wasMoving: false,
      resetSendingTimeout: 0,
      lastSendAt: null,
      leftHandedMode: false,
      lastPayload: null,
    }
  },
  mounted() {
    this.timer = setInterval(this.sendValues, 50)
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  computed: {
    items() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes('ptz')
      })
    }
  },
  methods: {
    onSelect(id) {
      this.camId = id
    },
    resetSending() {
      this.isSending = false
    },
    async sendValues() {
      if(!this.camId) {
        return
      }

      if(!this.pad) {
        return
      }

      let camX = this.pad.cam.x;
      let camY = this.pad.cam.y;
      let isMoving = false;

      camY *= 2/3;

      if(camX != 0 || camY != 0) {
        isMoving = true;
      }

      if(this.isSending && isMoving == this.wasMoving) {
        return
      }
      this.isSending = true
      clearTimeout(this.resetSendingTimeout)
      this.resetSendingTimeout = setTimeout(this.resetSending, 300)
      this.wasMoving = isMoving;

      let payload = {
        pan: camX,
        tilt: camY,
        zoom: this.pad.zoom,
        focus: this.pad.focus,
        buttons: this.pad.buttons,
      }
      if(JSON.stringify(this.lastPayload) === JSON.stringify(payload)) {
        return
      }

      this.lastPayload = payload
      try {
        await this.$store.state.config.apiClient.realtime.ptzJoystick(this.camId, payload)
        this.lastPayload = payload
      } catch(e) {
        console.error(e)
      }
    }
  }
}
</script>

<style scoped>
.center {
  margin: auto;
}

.margin-top {
  margin-top: 30px !important;
}
</style>
