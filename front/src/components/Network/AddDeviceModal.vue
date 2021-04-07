<template>
  <v-dialog v-model="activated" persistent max-width="600px">
    <v-card>
      <v-card-title>
        <span class="headline">Add {{deviceType}}</span>
      </v-card-title>
      <v-card-text>
        <v-form v-on:input="inputChanged" ref="form">
          <v-container grid-list-md>
            <center v-if="!params">
              <loading/>
            </center>
            <v-layout wrap v-else>
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
              <v-flex xs12 sm12 v-for="(meta, name) in params" :key="name">
                <v-select
                  item-text="name"
                  item-value="value"
                  :label="meta.description"
                  :value="values[name]"
                  :items="meta.options"
                  v-on:input="values[name] = $event"
                  :error-messages="errors.collect(name)"
                  :data-vv-name="name"
                  v-if="meta.type == 'select'"
                  />
                <v-text-field
                  :value="values[name]"
                  v-on:input="values[name] = $event"
                  :label="meta.description"
                  v-validate="validatorFor(meta)"
                  :error-messages="errors.collect(name)"
                  :data-vv-name="name"
                  v-else
                  />
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
        <v-btn color="blue darken-1" text @click="$emit('close')">Close</v-btn>
        <v-btn color="blue darken-1" text v-bind:disabled="submitDisabled" v-bind:loading="loading" @click="submit">Save</v-btn>
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
      params: null,
      name: "",
      error: null,
      values: {},
    }
  },
  computed: {
    submitDisabled() {
      return this.loading || this.invalidForm
    }
  },
  mounted() {
    this.fetchParams()
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
      this.loading = false
      let data = this.values;
      // If the device params where not valid or our own form was not valid
      if(data == null || !valid) {
        // Abort
        this.loading = false
        return
      }
      for(let name of Object.keys(data)) {
        if(this.params[name].type == "number") {
          try{
            data[name] = parseInt(data[name])
          } catch(e) {
            this.loading = false
            this.errors.add({
              field: name,
              msg: "must be a number"
            })
            return
          }
        }
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
    },
    async fetchParams() {
      if(!this.deviceType) {
        this.params = null
        return
      }
      let client = this.$store.state.config.apiClient
      try{
        this.params = await client.devices.typeParams(this.deviceType)
      } catch(e) {
        this.error = `Fail to get device params ${e.toString()}`
        return
      }

      let values = {}
      for(let name of Object.keys(this.params)) {
        values[name] = null
        if(this.params[name].default !== undefined) {
          values[name] = this.params[name].default
        }
      }
      this.values = values
    },
    validatorFor(meta) {
      let validators = []
      if(meta.required) {
        validators.push("required")
      }
      if(meta.type == "ip") {
        validators.push("ip")
      }
      if(meta.type == "number") {
        validators.push("numeric")
      }
      if(meta.min !== undefined) {
        validators.push(`min_value:${meta.min}`)
      }
      if(meta.max !== undefined) {
        validators.push(`max_value:${meta.max}`)
      }

      return validators.join("|")
    }
  },
  watch:{
    deviceType() {
      this.fetchParams()
    }
  }
}
</script>
