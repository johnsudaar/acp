<template>
  <v-dialog v-model="activated" persistent max-width="600px">
    <v-card>
      <v-card-title>
        <span class="headline">Add {{deviceType}}</span>
      </v-card-title>
      <v-card-text>
        <v-form v-on:input="inputChanged" ref="form">
          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12>
                <v-text-field
                  label="Name"
                  v-validate="'required'"
                  :error-messages="errors.collect('name')"
                  data-vv-name="name"
                  v-model="name"
                  :disabled="loading"
                  required/>
              </v-flex>
              <v-flex xs12>
                <component :is="deviceFormComponent" ref="deviceForm" :disabled="loading"/>
              </v-flex>
            </v-layout>
          </v-container>
        </v-form>
        <v-alert :value="error" dismissible type="error" transition="scale-transition">
          {{ error }}
        </v-alert>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" flat @click="$emit('close')">Close</v-btn>
        <v-btn color="blue darken-1" flat v-bind:disabled="submitDisabled" v-bind:loading="loading" @click="submit">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ['activated', 'deviceType'],
  data() {
    return {
      invalidForm: true,
      loading: false,
      name: "",
      error: null,
    }
  },
  computed: {
    // Get the component name to get a form for the current component type
    deviceFormComponent() {
      let type = this.deviceType
      if(type === null) {
        return "loading"
      }
      type = type.toLowerCase().replace(/_/g, '-')
      return `network-form-${type}`
    },
    submitDisabled() {
      return this.loading || this.invalidForm
    }
  },
  methods: {
    inputChanged(valid) {
      this.invalidForm = !valid;
    },
    async submit() {
    this.loading = true
      this.error = null
      // Submit our part of the form
      let valid = await this.$validator.validateAll()
      // Get the device data
      let data = await this.$refs.deviceForm.submit()
      // If the device params where not valid or our own form was not valid
      if(data == null || !valid) {
        // Abort
        this.loading = false
        return
      }
      let device = null;
      try {
        device = await this.$store.dispatch('devices/create',{
          name: this.name,
          type: this.deviceType,
          params: data
        })
      } catch(e) {
        this.error = e.toString()
        this.loading = false
        return
      }
      this.loading = false
      this.$emit('add-device', device.id)
      this.$emit('close')
    }
  }
}
</script>
