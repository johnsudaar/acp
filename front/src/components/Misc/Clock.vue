// Clock.vue
//
// Display a clock
// When created this component will add a interval with id intervalID.
// When destroyed the interval will be cleared to prevent interval leaking.
<script>
import format from 'date-fns/format'
export default {
  data() {
    return {
      intervalID: null,
      curTime: "00:00:00",
    }
  },
  methods: {
    startClock() {
      clearInterval(this.intervalID)
      this.intervalID = setInterval(()=>{
        this.curTime = format(
          new Date(),
          "HH:mm:ss"
        )
      }, 100)
    }
  },
  created() {
    this.startClock()
  },
  beforeDestroy() {
    clearInterval(this.intervalID)
  }
}
</script>

<style scoped>
.clock{
  width: 10rem;
}
</style>

<template>
  <div class="clock">
    <seven-seg :text="curTime"/>
  </div>
</template>
