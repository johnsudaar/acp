<template>
  <v-layout>
    <atem-btn v-for="(pos) in positions" :key="pos.id" :name="pos.name" big @click="positionClicked(pos.id)" :red="selectedPos == pos.id"/>
  </v-layout>
</template>

<script>
export default {
  props: {
    device: Object,
  },
  data() {
    return {
      selectedPos: null,
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
    async positionClicked(id) {
      let pos = this.$store.getters['ptzpositions/find'](this.device.id, id)
      if(!pos) {
        return
      }
      try {
        await this.$store.state.config.apiClient.ptz.position(this.device.id, {
          pan: pos.pan,
          tilt: pos.tilt,
          zoom: pos.zoom,
          focus: 0,
        })
        this.selectedPos = pos.id
      } catch(error) {
        console.log(error)
      }
    }
  }
}
</script>
