<template>
  <v-layout wrap>
    <v-flex xs12 sm12>
      <v-text-field
        label="Shoot name"
        v-validate="'required'"
        :error-messages="errors.collect('shoot')"
        data-vv-name="shoot"
        v-model="shoot"
        :disabled="disabled"
        required/>
    </v-flex>
  </v-layout>
</template>

<script>

export default {
  data() {
    return {
      shoot: null,
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
        shoot: this.shoot,
      }
    }
  }
}
</script>
