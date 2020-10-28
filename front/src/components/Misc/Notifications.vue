<template>
  <div class="text">
    <v-menu offset-y>
      <template v-slot:activator="{ on, attrs }">
        <v-badge left color="red" overlap :value="notificationCount != 0" v-on="on" v-bind="attrs">
          <span slot="badge">{{notificationCount}}</span>
          <v-icon large>
            notification_important
          </v-icon>
        </v-badge>
      </template>
      <v-list two-line>
        <v-list-tile v-if="$store.state.config.nextRelease">
          <v-list-tile-content>
            <v-list-tile-title>New Version Available</v-list-tile-title>
            <v-list-tile-sub-title><strong>{{$store.state.config.nextRelease.tag_name}} - </strong>{{$store.state.config.nextRelease.name}}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-menu>
  </div>
</template>

<script>
export default {
  computed: {
    notificationCount() {
      let count = 0
      if(this.$store.state.config && this.$store.state.config.updateAvailable) {
        count ++
      }
      return count
    }
  }
}
</script>
