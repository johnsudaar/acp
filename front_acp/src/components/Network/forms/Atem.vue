<template>
  <v-layout wrap>
    <v-flex xs12 sm6>
      <v-text-field
        label="IP"
        v-validate="'required|ip'"
        :error-messages="errors.collect('ip')"
        data-vv-name="ip"
        v-model="ip"
        :disabled="disabled"
        required/>
    </v-flex>
    <v-flex xs12 sm6>
      <v-text-field
        label="Port"
        v-validate="'required|between:1,65535'"
        :error-messages="errors.collect('port')"
        data-vv-name="port"
        v-model="port"
        :disabled="disabled"
        required/>
    </v-flex>
    <v-flex xs12>
      <v-text-field
        label="My Port"
        v-validate="'required|between:1,65535'"
        :error-messages="errors.collect('myPort')"
        data-vv-name="myPort"
        v-model="myPort"
        :disabled="disabled"
        required/>
    </v-flex>
  </v-layout>
</template>

<script>

export default {
  data() {
    return {
      ip: null,
      port: null,
      myPort: null,
    }
  },
  props: ['disabled'],
  methods: {
    async submit() {
      // Validate our form
      let valid = await this.$validator.validateAll()
      if(!valid) {
        // If it was not valid return null
        return null
      }
      // Data was valid, return user data
      return {
        ip: this.ip,
        port: parseInt(this.port),
        my_port: parseInt(this.myPort)
      }
    }
  }
}

</script>
