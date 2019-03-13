<script>
export default {
  name: 'atem-btn',
  props: {
    name: String,
    red: Boolean,
    green: Boolean,
    yellow: Boolean,
  },
  computed: {
    btnClass(){
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
.atem.btn .btn{
  height: 50px;
  width: 50px;
  background-image: url(/images/ATEM/Empty.png);
  background-color: transparent !important;
  background-size: cover;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: bold;
  color: black;
  cursor: pointer;
  z-index: 2;
  position: absolute;
}

.atem.btn .btn.red{
  background-image: url(/images/ATEM/Red.png);
}

.atem.btn .btn.green{
  background-image: url(/images/ATEM/Green.png);
}

.atem.btn .btn.yellow{
  background-image: url(/images/ATEM/Yellow.png);
}


.atem.btn .btn-shadow {
  z-index: 1;
  position: absolute;
  margin: 2px;
  height: 45px;
  width: 45px;
}

.atem.btn .btn-shadow.red{
  box-shadow: 0 0 10px red;
}

.atem.btn .btn-shadow.green{
  box-shadow: 0 0 10px green;
}

.atem.btn .btn-shadow.yellow{
  box-shadow: 0 0 10px yellow;
}

.atem.btn .label {
  text-align: center;
  margin: 0;
  width: 40px;
  max-height: 40px;
  overflow: hidden;
  text-overflow: ellipsis;
}

.atem.btn{
  margin: 2px;
  height: 50px;
  width: 50px;
}
</style>

<template>
  <div class="atem btn">
    <div v-bind:class="btnClass" class="btn-shadow" />
    <div v-bind:class="btnClass" class="btn">
      <div class="label">
        {{ name }}
      </div>
    </div>
  </div>
</template>
