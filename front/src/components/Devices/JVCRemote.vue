<template>
  <v-container>
    <v-form
      @submit="submitForm">
      <v-layout row>
        <v-spacer/>
        <v-btn
          :loading="loading"
          type="submit"
          color="primary">
          Save
        </v-btn>
      </v-layout>
      <v-layout row>
        <v-flex sm2 xs12 offset-sm4 class="pad-right">
          <v-select
            :loading="loading"
            :items="$store.state.devices.devices"
            @change="switchSource = $event"
            label="Switch source"
            item-text="name"
            item-value="id"
            />

        </v-flex>
        <v-flex sm2 xs12>
          <v-select
            :loading="loading"
            :items="switchPorts"
            :disabled="switchPorts == null"
            @change="switchPort = $event"
            label="Output Port"
            item-text="name"
            item-value="id"
            />
        </v-flex>
      </v-layout>
      <jvc-remote-input :input="input" v-for="input in inputs"/>
      <v-layout row>
        <v-btn
          @click="addInput"
          color="pink"
          dark
          small
          fab>
          <v-icon>add</v-icon>
        </v-btn>
      </v-layout>
    </v-form>
  </v-container>
</template>

<script>

export default {
  data() {
    return {
      inputs: [],
      loading: false,
      switchPorts: null,
    }
  },
  methods: {
    addInput() {
      this.inputs.push({ip: "192.168.20.101", input: 10})
    },
    submitForm(e) {
      e.preventDefault()
      this.loading = true
    }
  }
}
</script>

<style scoped>
.pad-right {
  padding-right: 10px;
}
</style>
