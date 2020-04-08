<script>
import {mapState} from 'vuex'

export default {
  computed: {
    ...mapState('config', ['connected'])
  }
}
</script>

<style scoped>
.status-container {
  display: flex;
  height: 100%;
  align-items: center;
  margin-right: 20px;
}

.status-container p {
  margin: 0;
}

.status-led {
  width: 10px;
  height: 10px;
  background-color: red;
  margin-right: 10px;
  border-radius: 5px;
  box-shadow: 0 0 10px red;
}

.status-led.connected{
  background-color: green;
  box-shadow: 0 0 10px green;
}
</style>

<template>
  <div class="status-container">
    <div class="status-led" v-bind:class="{connected: connected}"/>
    <p v-if="connected"> Connected </p>
    <p v-else> Disconnected </p>
  </div>
</template>
