<template>
  <v-layout wrap>
    <v-flex xs12>
      <v-text-field
        v-model="name"
        label="Position name"
        />
    </v-flex>
    <v-flex xs12>
      <v-select :items="availablePositionGroups"
                v-model="position_group_id"
                item-text="name"
                item-value="id"
                label="Position Group"
                hide-details/>
    </v-flex>
    <v-flex xs9>
      <v-slider
        v-model="pan"
        step="0.1"
        :max="100"
        label="Pan"
        ></v-slider>
    </v-flex>

    <v-flex xs3>
      <v-text-field
        step=".1"
        v-model="pan"
        class="mt-0 ml-2"
        type="number"
        ></v-text-field>
    </v-flex>
    <v-flex xs9>
      <v-slider
        v-model="tilt"
        step="0.1"
        :max="100"
        label="Tilt"
        ></v-slider>
    </v-flex>

    <v-flex xs3>
      <v-text-field
        step=".1"
        v-model="tilt"
        class="mt-0 ml-2"
        type="number"
        ></v-text-field>
    </v-flex>
    <v-flex xs9>
      <v-slider
        v-model="zoom"
        step="0.1"
        :max="100"
        label="Zoom"
        ></v-slider>
    </v-flex>

    <v-flex xs3>
      <v-text-field
        step=".1"
        v-model="zoom"
        class="mt-0 ml-2"
        type="number"
        ></v-text-field>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  props: {
    pos: Object,
    device: Object,
    disabled: Boolean,
  },
  data () {
    return {
      pan: 0.0,
      tilt: 0.0,
      zoom: 0.0,
      name: "",
      position_group_id: null,
      updateTimeout: null,
    }
  },
  mounted() {
    this.copyPosParams()
  },
  computed: {
    availablePositionGroups() {
      let res = [{
        name: "(None)",
        id: null,
      }];
      for(let group in this.$store.state.positiongroups.groups) {
        res.push(this.$store.state.positiongroups.groups[group])
      }
      return res;
    },
  },
  watch: {
    pos: function() {
      this.copyPosParams()
    },
    name: function() {
      this.onInputChanged()
    },
    pan: function() {
      this.onInputChanged()
    },
    tilt: function() {
      this.onInputChanged()
    },
    zoom: function() {
      this.onInputChanged()
    },
    position_group_id: function() {
      this.onInputChanged()
    },
  },
  methods: {
    copyPosParams() {
      if(!this.pos) {
        return
      }
      this.pan = this.pos.pan
      this.tilt = this.pos.tilt
      this.zoom = this.pos.zoom
      this.name = this.pos.name
      this.position_group_id = this.pos.position_group_id
      this.onInputChanged()
    },
    previewPosition() {
      if(this.disabled) {
        return
      }

      this.$store.state.config.apiClient.ptz.position(this.device.id, {
        pan: this.pan,
        tilt: this.tilt,
        zoom: this.zoom,
        focus: 0,
      })
      this.updateTimeout = null
    },
    onInputChanged() {
      this.$emit('changed', {
        pan: this.pan,
        tilt: this.tilt,
        zoom: this.zoom,
        name: this.name,
        position_group_id: this.position_group_id,
      })
      if(this.updateTimeout != null) {
        return
      }
      this.updateTimeout = setTimeout(()=> {
        this.previewPosition()
      }, 100)
    },
  }
}
</script>
