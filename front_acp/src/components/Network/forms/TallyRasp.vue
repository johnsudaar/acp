<template>
  <v-layout wrap>
    <v-flex xs12 sm12>
      <v-text-field
        label="IP"
        v-validate="'required|ip'"
        :error-messages="errors.collect('ip')"
        data-vv-name="ip"
        v-model="ip"
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
      }
    }
  }
}
</script>
