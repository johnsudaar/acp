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
    positionGroups: Array,
  },
  data() {
    return {
      activePosition: null,
    }
  },
  computed: {
    positions() {
      let positions = this.$store.getters['ptzpositions/forDevice'](this.device.id);
      if(!positions) {
        return positions;
      }

      return positions.filter((position) => {
        let group = position.position_group_id;
        if(group === "" || group === null) {
          group = "*"
        }
        return this.positionGroups.includes(group);
      })
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
