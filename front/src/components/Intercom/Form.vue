<template>
  <v-dialog v-model="activated" persistent max-width="600px">
    <v-card>
      <v-card-title>
        <span class="headline"> Create a new room </span>
      </v-card-title>
      <v-card-text>
        <v-form ref="form">
          <v-layout wrap>
            <v-flex xs12 sm12>
              <v-text-field label="Name" v-validate="'required'" data-vv-name="name" v-model="name" :error-messages="errors.collect('name')" required/>
            </v-flex>
            <v-flex xs12>
              <v-text-field label="Channel" v-validate="'required|numeric'" data-vv-name="channel" v-model="channel" :error-messages="errors.collect('channel')" required/>
            </v-flex>
            <v-flex xs12>
              <v-text-field label="Mix" v-validate="'required|numeric'" data-vv-name="mix" v-model="mix" :error-messages="errors.collect('mix')" required/>
            </v-flex>
          </v-layout>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer>
        </v-spacer>
        <v-btn color="blue darken-1" flat @click="$emit('close')"> Close </v-btn>
        <v-btn color="blue darken-1" flat @click="submit"> Save </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: ['activated', 'device'],
  data() {
    return {
      name: "",
      channel: 0,
      mix: 0,
    }
  },
  methods: {
    async submit() {
      let valid = await this.$validator.validateAll()
      if(!valid) {
        return
      }
      let params = {
        name: this.name,
        mix: parseInt(this.mix),
        channel: parseInt(this.channel),
      }
      try {
        await this.$store.state.config.apiClient.intercom.createRoom(this.device.id, params)
      } catch(e) {
        console.error(e)
        return
      }
      this.$emit('close')
    }
  }
}
</script>
