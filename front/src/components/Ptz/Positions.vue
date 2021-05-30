<template>
  <v-layout>
    <atem-btn v-for="(pos) in positions" :key="pos.id" :name="pos.name" big @click="positionClicked(pos.id)" @contextmenu.native.prevent="positionRightClicked(pos.id)" :red="pos.id === activePosition"/>
  </v-layout>
</template>

<script>
import {PtzPositionEditBus} from '@/buses'

export default {
  props: {
    device: Object,
  },
  data() {
    return {
      activePosition: null,
    }
  },
  computed: {
    positions() {
      return this.$store.getters['ptzpositions/forDevice'](this.device.id)
    }
  },
  mounted() {
    this.$store.dispatch('ptzpositions/ensure', this.device.id)
  },
  methods: {
    positionRightClicked(positionId) {
      PtzPositionEditBus.$emit("requestEditFormFor", this.device.id, positionId)
    },
    async positionClicked(id) {
      let pos = this.$store.getters['ptzpositions/find'](this.device.id, id)
      if(!pos) {
        return
      }
      try {
        await this.$store.state.config.apiClient.ptz.position(this.device.id, {
          position_id: id,
        })
      } catch(error) {
        console.log(error)
        return
      }
      PtzPositionEditBus.$emit("setActivePositionFor", this.device.id, id)
      this.activePosition = id
    }
  }
}
</script>
