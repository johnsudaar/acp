<script>
export default {
  name: "atem-btn-grp",
  props: {
    name: String
  }
}
</script>

<style scoped>
.atem.btn-grp {
  background-color: rgba(0,0,0,.5);
  border-radius: 4px;
  display: flex;
  padding: 5px;
}
</style>

<template>
  <group :title="name" class="mt-3">
    <div class="atem btn-grp">
      <atem-btn name="Cam 1" red/>
      <atem-btn name="Cam 2" green/>
      <atem-btn name="Cam 3" yellow/>
      <atem-btn name="Cam 4"/>
      <atem-btn name="Cam 5"/>
      <atem-btn name="Cam 6"/>
      <atem-btn name="Cam 7"/>
      <atem-btn name="Cam 8"/>
      <atem-btn name="Cam 9"/>
      <atem-btn name="Cam 10"/>
      <atem-btn name="Cam 11"/>
      <atem-btn name="Cam 12"/>
    </div>
  </group>
</template>
