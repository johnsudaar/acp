<template>
  <v-dialog
    :value="open"
    width="500"
    :persistent="true"
    >
    <v-card>
      <v-card-title
        class="headline"
        primary-title
        >
        Edit a position for {{device.name}}
      </v-card-title>
      <v-card-text>
        <v-flex xs12 d-flex>
          <v-select :items="positions" label="Position to edit" item-text="name" item-value="id" @change="onPosChanged"/>
        </v-flex>
        <ptz-form v-if="pos" @changed="onChanged" :device="device" :pos="pos" :disabled="!open"/>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="red darken-1" text @click="close()">Close</v-btn>
        <v-btn color="green darken-1" text v-bind:disabled="submitDisabled" v-bind:loading="loading" @click="submit()">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import {PtzPositionEditBus} from '@/buses'

export default {
  props: {
    device: Object,
    open: Boolean,
  },
  created() {
    PtzPositionEditBus.$on("requestEditFormFor", (device, position) => {
      if(device === this.device.id) {
        this.posId = position;
      }
    })
    PtzPositionEditBus.$on("setActivePositionFor", (device, position) => {
      if(device === this.device.id) {
        this.posId = position;
      }
    })
  },
  data () {
    return {
      loading: false,
      params: {},
      posId: null,
    }
  },
  computed: {
    pos() {
      return this.$store.getters['ptzpositions/find'](this.device.id, this.posId)
    },
    positions() {
      return this.$store.getters['ptzpositions/forDevice'](this.device.id)
    },
    submitDisabled() {
      return this.loading
    }
  },
  methods: {
    onPosChanged(id) {
      this.posId = id
    },
    onChanged(params) {
      this.params = params
    },
    async submit() {
      var response
      try {
        response = await this.$store.state.config.apiClient.ptz.updatePosition(this.device.id, this.pos.id, this.params)
      } catch(error) {
        console.error(error)
        return
      }
      this.$store.commit('ptzpositions/setPosition', {
        cam: this.device.id,
        position: response,
      })
      this.close()
    },
    close() {
      this.posId = null
      this.$emit('close')
    }
  }
}
</script>
