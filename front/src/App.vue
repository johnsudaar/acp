<script>
import {mapState} from 'vuex'
import deviceFilters from './filters/deviceFilters'

export default {
  name: 'app',
  data: () => ({
    drawer: null,
  }),
  props: {
    source: String,
  },
  computed: {
    ...mapState('devices', ['devices']),
  },
  created() {
    // Load configuration from localStorage
    this.$store.dispatch('config/load')
  },
  methods: {
    showDevice(device) {
      switch(device.type) {
        case "JVC_REMOTE":
          return true
      }
      return false
    }
  },
  filters: deviceFilters,
}
</script>

<style>
@font-face {
  font-family: sevenSeg;
  src: url(../public/fonts/DSEG7Modern-Bold.woff);
}

body {
  margin: 0;
}

.fill-width {
  width: 100%;
}

.no-grow {
  flex-grow: 0;
}
</style>

<template>
  <v-app id="inspire" dark>
    <v-navigation-drawer
      v-model="drawer"
      clipped
      fixed
      app
    >
      <v-list dense>
        <navigation-link title="Network" icon="device_hub" path="/"/>
        <navigation-link v-for="(device) in devices" :icon="device.type | deviceTypeIcon" :title="device.name" :path="device.path()" v-if="showDevice(device)"/>
        <navigation-link title="CCU" icon="camera" path="/ccu"/>
        <navigation-link title="Rec control" icon="camera" path="/rec/control"/>
        <navigation-link title="PTZ Control" icon="gamepad" path="/cam/control"/>
        <navigation-link title="Cockpit" icon="gamepad" path="/cockpit"/>
        <v-spacer/>
        <navigation-link title="Configuration" icon="settings" path="/config" />
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app fixed clipped-left>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>ACP Control Board</v-toolbar-title>
      <v-spacer/>
      <notifications class="mr-2"/>
      <server-status/>
      <clock/>
    </v-toolbar>
    <v-content>
      <v-layout column v-if="$store.state.config.connected || $route.meta.offline" fill-height>
        <v-container fluid fill-height>
          <router-view> </router-view>
        </v-container>
      </v-layout>
      <v-layout fill-height align-center justify-center v-else>
        <loading />
      </v-layout>
    </v-content>
    <v-footer app fixed>
    </v-footer>
  </v-app>
</template>
