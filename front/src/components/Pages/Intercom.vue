<template>
  <v-layout row wrap>
    <v-flex xs12 v-if="x32 === null">
      <v-alert value="true" type="error" transition="scale-transition" >
        No X32 found
      </v-alert>
    </v-flex>

    <v-container v-else>
      <group title="Control" class="col-12 mt-3" fill-width>
        <v-layout>
          <v-select v-if="rooms"
            :items="rooms"
            label="My Room"
            item-text="name"
            item-value="id"
            v-model="myRoomID"
            />
          <v-btn dark @click="createDevice" inline>
            Create Room
          </v-btn>
        </v-layout>
      </group>
      <group title="Buttons" class="col-12 mt-3" fill-width>
      <atem-btn v-for="(room) in otherRooms" :name="room.name" @mouseup="stop(room)" @mousedown="start(room)"/>
      </group>
      <intercom-form
        :v-if="showModal"
        :activated="showModal"
        :device="x32"
        v-on:close="closeModal"/>
    </v-container>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      showModal: false,
      rooms: null,
      myRoomID: null,
    }
  },
  computed: {
    x32() {
      let x32 = this.$store.state.devices.devices.filter((device) => {
        return device.type == "X32"
      })
      if(x32.length === 0) {
        return null
      }
      this.fetchRooms(x32[0])
      return x32[0]
    },
    otherRooms() {
      if(this.rooms === null) {
        return null
      }

      return this.rooms.filter((room) => {
        return room.id != this.myRoomID
      })
    }
  },
  created() {
    document.addEventListener("keyup", this.onKeyPressed)
    document.addEventListener("keydown", this.onKeyPressed)
  },
  destroyed() {
    document.removeEventListener("keyup", this.onKeyPressed)
    document.removeEventListener("keydown", this.onKeyPressed)
  },
  methods: {
    async fetchRooms(device) {
      let rooms = null;
      try {
        rooms = await this.$store.state.config.apiClient.intercom.listRooms(device.id)
      } catch(e) {
        console.error(e)
        return
      }
      this.rooms = rooms
    },
    createDevice() {
      this.showModal = true
    },
    closeModal() {
      this.showModal = false
    },
    start(room) {
      console.log("Start", room)
    },
    stop(room) {
      console.log("Stop", room)
    },
    onKeyPressed(key) {
      console.log("pressed", key)
    },
    onKeyReleased(key) {
      console.log("released", key)
    }
  }
}
</script>
