<script>
import {mapState} from 'vuex'
import deviceFilters from './filters/deviceFilters'

export default {
  name: 'app',
  data: () => ({
    drawer: null,
  }),
  mounted() {
    this.updateFullscreen()
  },
  updated() {
    this.updateFullscreen()
  },
  props: {
    source: String,
  },
  computed: {
    ...mapState('devices', ['devices']),
  },
  created() {
    let location = new URL(window.location.href)
    let ip = localStorage.getItem('config/serverEndpoint/ip')
    let port = localStorage.getItem('config/serverEndpoint/port')
    // If the IP or the port does not exist, store try to autoguess them
    if(!ip) {
      localStorage.setItem('config/serverEndpoint/ip', location.hostname)
    }
    if(!port) {
      localStorage.setItem('config/serverEndpoint/port', location.port || "80")
    }
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
    },
    updateFullscreen() {
      if(this.$route.meta.fullscreen) {
        this.$refs['fullscreen'].requestFullscreen();
      }
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
html {
  overflow-y: auto;
}

.fill-width {
  width: 100%;
}

.no-grow {
  flex-grow: 0;
}

.fullscreen {
  background-color: black;
  width: 100vw;
  height: 100vh;
}
</style>

<template>
  <div ref="fullscreen" class="fullscreen" v-if="$route.meta.fullscreen ">
    <v-app>
      <v-main>
        <router-view v-if="$store.state.config.connected"/>
      </v-main>
    </v-app>
  </div>
  <v-app v-else>
    <v-navigation-drawer
      v-model="drawer"
      clipped
      fixed
      app
    >
      <v-list dense >
        <navigation-link title="Network" icon="device_hub" path="/"/>
        <navigation-link v-for="(device) in devices" :icon="device.type | deviceTypeIcon" :title="device.name" :path="device.path()" v-if="showDevice(device)"/>
        <!-- <navigation-link title="CCU" icon="camera" path="/ccu"/>
        <navigation-link title="Rec control" icon="camera" path="/rec/control"/> -->
        <navigation-link title="PTZ Control" icon="gamepad" path="/cam/control"/>
        <navigation-link title="Cockpit" icon="gamepad" path="/cockpit"/>
        <navigation-link title="Scenes" icon="tv" path="/scenes"/>
        <navigation-link title="Timers" icon="schedule" path="/timers"/>
        <navigation-link title="Program" icon="event_note" path="/programs"/>
        <v-spacer/>
        <navigation-link title="Configuration" icon="settings" path="/config" />
        <navigation-link title="Screen" icon="tv" path="/scenes/_active"/>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      app
      clipped-left>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>ACP Control Board</v-toolbar-title>
      <v-spacer/>
      <notifications class="mr-2"/>
      <server-status/>
      <clock/>
    </v-app-bar>
    <v-main>
      <v-layout column v-if="$store.state.config.connected || $route.meta.offline" fill-height>
        <v-container fluid fill-height fill-width :pa-0="$route.meta.doubleMenu">
          <router-view> </router-view>
        </v-container>
      </v-layout>
      <v-layout fill-height align-center justify-center v-else>
        <loading />
      </v-layout>
    </v-main>
    <v-footer app fixed>
    </v-footer>
  </v-app>
</template>
