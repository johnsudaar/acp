<script>
export default {
  props: {
    title: { type: String },
    'fill-height': {type: Boolean},
    'fill-width': {type: Boolean},
    'bg': {type: String, default: ""},
  },
  computed: {
    containerClasses() {
      let result = this.bg
      if(this.fillWidth) {
        result += " fill-width"
      }
      return result
    }
  }
}

</script>
<style scoped>
.grp-container {
  border: 1px solid gray;
  padding: 5px;
  border-radius: 3px;
  width: fit-content;
  height: fit-parent;
}

.grp-title {
  width: fit-content;
  padding: 0 5px;
  margin-top: -15px;
  background-color: #303030;
  height: 20px;
}

.grp-container.fill-width {
  width: 100%;
}

.fill-height {
  height: 100%;
}
</style>

<template>
  <div class="grp-container" :class="containerClasses">
    <div class="grp-title"> {{title}} </div>
    <div class="grp-content" v-bind:class="{ 'fill-height': 'fill-height' }">
      <slot> </slot>
    </div>
  </div>
</template>
