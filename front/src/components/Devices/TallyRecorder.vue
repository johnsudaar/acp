<template>
  <v-data-table
    :headers="headers"
    :items="results"
    class="elevation-1"
    >
    <v-progress-linear v-slot:progress color="blue" indeterminate></v-progress-linear>
    <template v-slot:items="props">
      <td> aaaa </td>
    </template>
  </v-data-table>
</template>

<script>
export default {
  props: {
    deviceId: String
  },
  data() {
    return {
      headers: [{text: 'Created At', value: 'created_at', sortable: false}],
      results: [],
      error: null,
      loading: true,
    }
  },
  mounted() {
    this.startSearch()
  },
  watch: {
    '$store.state.config.connected': async function() {
      // If we weren't connected on page load
      if(this.$store.state.config.connected) {
        // If we are connected now
        await Vue.nextTick()
        // Wait for next tick (wait for template evaluation)

        this.startSearch()
      }
    }
  },
  methods: {
    async startSearch() {
      this.loading = true
      if(!this.$store.state.config.connected) {
        return
      }
      let client = this.$store.state.config.apiClient
      try {
        this.results = await client.tally_recorder.search(this.deviceId)
      } catch(error) {
        this.error = error
        console.log(error)
      }
      this.loading = false
    }
  }
}
</script>
