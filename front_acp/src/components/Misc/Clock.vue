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
.clock {
  color: red;
  background-color: black;
  padding: 7px;
  font-size: 20px;
  font-family: sevenSeg;
}

</style>

<template>
  <seven-seg :text="curTime"/>
</template>
