<script>

export default {
  props: {
    name: String,
    red: Boolean,
    green: Boolean,
    yellow: Boolean,
  },
  computed: {
    colorClass(){
      let colors = ['red', 'green', 'yellow']
      for(let color in colors){
        color = colors[color];
        if(this[color]){
          return color;
        }
      }
      return 'empty'
    }
  }
}

</script>

<style scoped>
.led {
  width:50px;
  height: 50px;
  margin: 5px;
  border-radius: 25px;
  background-image: radial-gradient(#ffffff4f 0%, transparent 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: black;
}

.red {
  background-color: red !important;
  box-shadow: 0 0 10px #ff0000;
}

.green {
  background-color: green !important;
  box-shadow: 0 0 10px #00ff007f;
}

.yellow {
  background-color: yellow !important;
  color: black;
  box-shadow: 0 0 10px #ffff007f;
}

</style>

<template>
  <div :class="colorClass" class="led"> {{name}} </div>
</template>
