<template>
  <v-card height="100%">
    <v-btn
      height="100%"
      width="100%"
      @click="open = true"
      text>
      <v-icon> mdi-plus </v-icon>
    </v-btn>
    <v-dialog
      :value="open"
      width="500"
      :persistent="true"
      >
      <timer-form @close="onClose"/>
    </v-dialog>
  </v-card>
</template>
test
<script>
export default {
  props: {
    sourceTimer: Object,
  },
  data() {
    return {
      open: false,
    }
  },
  methods: {
    onClose(event) {
      this.open = false;
    },
  }
}
</script>
