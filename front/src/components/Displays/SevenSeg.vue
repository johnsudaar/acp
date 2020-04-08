<script>

export default {
  props: {
    text: String,
    color: {
      type: String,
      default: "red"
    }
  }
}

</script>

<style scoped>
.seven-seg {
  background-color: black !important;
  padding: 7px;
  font-size: 20px;
  font-family: sevenSeg;
  border-radius: 3px;
  width: fit-content;
}

.seven-seg.red{
  color: red;
}

.seven-seg.yellow {
  color: yellow;
}

.seven-seg.green {
  color: green;
}

</style>

<template>
  <div :class="['seven-seg', color]">
    {{text}}
  </div>
</template>
