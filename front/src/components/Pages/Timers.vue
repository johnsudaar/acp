<template>
  <v-row align="stretch">
    <v-col sm="6" md="4" v-for="(timer) in $store.state.timers.timers" :key="timer.id">
      <timer-card :timer="timer"/>
    </v-col>
    <v-col sm="6" md="4">
      <add-timer />
    </v-col>
  </v-row>
</template>

<script>
import TimersMixin from '@/mixins/timers'

export default {
  mixins: [TimersMixin],
  data() {
    return {
    }
  },
  mounted() {
    this.$store.dispatch('timers/load')
  }
}

</script>
