<template>
  <group :title="name" :bg="color">
    <table>
      <tr>
        <td> Iris</td>
        <td> 1.5 (Manual) </td>
      </tr>
      <tr>
        <td> Gain </td>
        <td> 6db (Manual) </td>
      </tr>
      <tr>
        <td> Focus </td>
        <td> 2m (Manual) </td>
      </tr>
      <tr>
        <td> Black Level </td>
        <td> -2 </td>
      </tr>
      <tr>
        <td> WB </td>
        <td> 3200k (Manual) </td>
      </tr>
    </table>
  </group>
</template>

<script>
export default {
  props: {
    name: { type: String },
    color: { type: String },
  }
}
</script>

<style scoped>
</style>
