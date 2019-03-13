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
    <v-list-tile>
      <v-list-tile-action>
        <v-icon> {{ icon }}</v-icon>
      </v-list-tile-action>
      <v-list-tile-content>
        <v-list-tile-title> {{ title }} </v-list-tile-title>
      </v-list-tile-content>
    </v-list-tile>
  </router-link>
</template>
