<template>
  <v-form>
    <v-row>
      <v-col cols="6">
        <v-text-field
          v-model="elem.name"
          data-vv-name="name"
          :error-messages="errors.collect('name')"
          v-validate="'required'"
          label="Element name"
          />
      </v-col>
      <v-col cols="6">
        <v-select
          v-model="elem.type"
          :items="elemTypes"
          item-text="name"
          item-value="value"
          label="Element type"/>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="3">
        <v-text-field
          v-model="elem.startX"
          label="Start X"
          type="number"
          :min="1"
          :max="6"
          />
      </v-col>
      <v-col cols="3">
        <v-text-field
          v-model="elem.startY"
          label="Start Y"
          type="number"
          :min="1"
          :max="6"
          />
      </v-col>
      <v-col cols="3">
        <v-text-field
          v-model="elem.sizeX"
          label="Width"
          type="number"
          :min="1"
          :max="6"
          />
      </v-col>
      <v-col cols="3">
        <v-text-field
          v-model="elem.sizeY"
          label="Height"
          type="number"
          :min="1"
          :max="6"
          />
      </v-col>
    </v-row>

    <v-row
      v-if="elem.type == 'text'">
      <v-col cols="12">
        <v-textarea
          v-model="elem.text"
          rows=2
          label="Text"
          />
      </v-col>
    </v-row>
    <v-row
      v-if="elem.type == 'text'">
      <v-col cols="6">
        <v-text-field
          v-model="elem.color"
          label="Color"
          />
      </v-col>
      <v-col cols="6">
        <v-text-field
          v-model="elem.background_color"
          label="Background"
          />
      </v-col>
    </v-row>
    <v-row
      v-if="elem.type == 'timer'">
      <v-col cols="12">
        <v-select
          v-model="elem.timer_id"
          :items="timers"
          item-text="name"
          item-value="id"
          label="Timer"/>
      </v-col>
      <v-col cols="12">
        <v-text-field
          v-model="elem.color"
          label="Color"
          />
      </v-col>
    </v-row>
    <v-row
      v-if="elem.type == 'chat'">
      <v-col cols="6">
        <v-select
          v-model="elem.chatId"
          :items="chats"
          item-text="name"
          item-value="id"
          label="Chat"/>
      </v-col>
      <v-col cols="6">
        <v-text-field
          v-model="elem.chatFontSize"
          label="Chat Font Size"
          />
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  props: {
    item: Object,
  },
  data() {
    return {
      elem: {
        type: "text",
      },

      elemTypes: [
        {name: "Text", value: "text"},
        {name: "Timer", value: "timer"},
        {name: "Chat", value: "chat"},
      ],
    }
  },
  computed: {
    timers() {
      return Object.values(this.$store.state.timers.timers);
    },
    chats() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("chat")
      }).sort((a,b) => {
        return a.name > b.name ? 1 : -1
      })
    }
  },
  mounted() {
    if(this.item) {
      this.elem = this.item;
    }
  }
}
</script>
