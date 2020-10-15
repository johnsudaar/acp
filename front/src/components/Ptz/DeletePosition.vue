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
        Delete a position for {{device.name}}
      </v-card-title>
        <v-card-text>
          <v-flex xs12 d-flex>
            <v-select :items="positions" label="Position to delete" item-text="name" item-value="id" @change="onPosChanged"/>
          </v-flex>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red darken-1" flat @click="$emit('close')">Close</v-btn>
          <v-btn color="green darken-1" flat v-bind:disabled="submitDisabled" v-bind:loading="loading" @click="submit()">Delete</v-btn>
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
      return this.loading || !this.pos
    }
  },
  methods: {
    onPosChanged(id) {
      this.posId = id
    },
    async submit() {
      try {
        await this.$store.state.config.apiClient.ptz.deletePosition(this.device.id, this.pos.id)
      } catch(error) {
        console.error(error)
        return
      }
      this.$store.commit('ptzpositions/removePosition', {
        cam: this.device.id,
        id: this.pos.id,
      })
      this.$emit('close')
    }
  }
}
</script>
