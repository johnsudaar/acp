<template>
  <v-container pa-0 style="min-height: 100%">
    <v-toolbar dense>
      <v-spacer/>
      <v-col class="d-flex" cols="2"> 
        <v-select :items="presets" 
                  item-text="name"
                  item-value="id"
                  label="Presets"
                  dense hide-details multiple/>
      </v-col> 
      <position-groups-add />
      <v-btn>
        <v-icon> mdi-delete </v-icon>
      </v-btn>
    </v-toolbar>
    <v-container>
      <group v-for="(device) in ptzDevices" :key="device.id" :title="device.name" class="col-12 mt-3" fill-width v-bind:class="tallyClass(device.id)">
        <v-layout>
          <ptz-positions :device="device"/>
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
    presets() {
      let res = [{
        name: "(default)",
        id: "",
      }];
      for(let group in this.$store.state.positiongroups.groups) {
        res.push(group)
      }
      return res;
    }
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
