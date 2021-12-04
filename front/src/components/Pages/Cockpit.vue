<template>
  <v-container pa-0 style="min-height: 100%">
    <v-toolbar dense>
      <v-spacer/>
      <v-col class="d-flex" cols="4">
        <v-select :items="availablePositionGroups"
                  v-model="positionGroups"
                  item-text="name"
                  item-value="id"
                  label="Position Groups"
                  dense hide-details multiple/>
      </v-col>
      <position-groups-add />
      <position-groups-destroy />
    </v-toolbar>
    <v-container>
      <group v-for="(device) in ptzDevices" :key="device.id" :title="device.name" class="col-12 mt-3" fill-width v-bind:class="tallyClass(device.id)">
        <v-layout>
          <ptz-positions :device="device" :positionGroups="positionGroups"/>
          <v-spacer/>
          <ptz-edit :device="device"/>
        </v-layout>
      </group>

      <group v-for="(device) in switchers" :key="device.id" :title="device.name" class="col-12 mt-3" fill-width>
        <v-layout>
          <switcher :device="device"/>
        </v-layout>
      </group>
    </v-container>
  </v-container>
</template>

<script>
import TallyMixin from '@/mixins/tally'
export default {
  mixins: [TallyMixin],
  data() {
    return {
      fab: false,
      positionGroups: ['*'],
    }
  },
  computed: {
        ptzDevices() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("ptz")
      }).sort((a, b) => {
        return a.name > b.name ? 1 : -1
      })
    },
    switchers() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("switcher")
      }).sort((a, b) => {
        return a.name > b.name ? 1 : -1
      })
    },
    availablePositionGroups() {
      let res = [{
        name: "(default)",
        id: "*",
      }];
      for(let group in this.$store.state.positiongroups.groups) {
        res.push(this.$store.state.positiongroups.groups[group])
      }
      return res;
    }
  },
  mounted() {
    this.$store.dispatch('positiongroups/load')
  },

  methods: {
    tallyClass(deviceId) {
      if(this.isPreview(deviceId)) {
        return 'tally-pvw'
      }
      if(this.isProgram(deviceId)) {
        return 'tally-pgm'
      }
    }
  }
}
</script>

<style scoped>
.tally-pgm {
  background-color: #ff000030;
}

.tally-pvw {
  background-color: #00ff0030;
}
</style>
