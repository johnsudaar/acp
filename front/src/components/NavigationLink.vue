<script>
export default {
  props: [
    'icon',
    'title',
    'path',
  ]

}
</script>

<style scoped>
a {
  text-decoration: none;
  color: white;
}
</style>

<template>
  <router-link :to="path">
    <v-list-item>
      <v-list-item-action>
        <v-icon> {{ icon }}</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title> {{ title }} </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </router-link>
</template>
