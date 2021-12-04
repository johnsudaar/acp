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
  font-family: sevenSeg;
  border-radius: 3px;
  width: 100%;
}
</style>

<template>
  <svg viewBox="0 0 100 22" class="seven-seg">
    <text text-anchor="middle" x="50%" y="18" :fill="color">{{text}}</text>
  </svg>
</template>
