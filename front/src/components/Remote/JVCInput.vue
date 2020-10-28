<template>
  <v-layout>
    <v-flex xs12 sm4 offset-sm2>
      <v-text-field
        label="IP"
        v-validate="'required|ip'"
        :error-messages="errors.collect('ip')"
        data-vv-name="ip"
        v-model="input.ip"
        :disabled="loading"/>
    </v-flex>
    <v-flex xs12 sm3 offset-sm1 >
      <v-text-field
        label="Matrix Input"
        v-validate="'required|numeric'"
        :error-messages="errors.collect('input')"
        data-vv-name="input"
        v-model="input.input"
        :disabled="loading"/>
    </v-flex>
  </v-layout>
</template>

<script>

export default {
  props: {
    input: { type: Object },
    loading: { type: Boolean }
  }
}
</script>
