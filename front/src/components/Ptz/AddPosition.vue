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
        Add PTZ position for {{device.name}}
      </v-card-title>
      <v-card-text>
        <ptz-form @changed="onChanged" :device="device"/>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="red darken-1" text @click="$emit('close')">Close</v-btn>
        <v-btn color="green darken-1" text v-bind:disabled="submitDisabled" v-bind:loading="loading" @click="submit()">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: {
    device: Object,
    open: Boolean,
  },
  data () {
    return {
      loading: false,
      params: {},
    }
  },
  computed: {
    submitDisabled() {
      return this.loading
    }
  },
  methods: {
    onChanged(params) {
      this.params = params
    },
    async submit() {
      var response
      try {
        response = await this.$store.state.config.apiClient.ptz.storePosition(this.device.id, this.params)
      } catch(error) {
        console.error(error)
        return
      }
      this.$store.commit('ptzpositions/addPosition', {
        cam: this.device.id,
        position: response,
      })
      this.$emit('close')
    }
  }
}
</script>
