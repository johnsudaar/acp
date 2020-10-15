<script>
import API from '../../api'

export default {
  data() {
    return {
      ip: null, // IP of the server
      port: null, // Port
      error: null, // Was there an error when trying to contact the remote server?
      success: null, // Were we successful when trying to contact the remote server?
      loading: false, // The ping request is pending
    }
  },
  methods: {
    async connect() {
      this.error = null
      this.success = false
      this.loading = true

      // Form validation
      let valid = await this.$validator.validateAll();
      if(!valid) {
        this.loading = false
        return
      }

      let ip = this.ip || this.$store.state.config.ip;
      let port = this.port || this.$store.state.config.port;

      this.$store.commit('config/setServerEndpoint', {ip, port})

      // Temp API Client (to check IP and port)
      let api = new API(ip, port)
      try {
        // Try to ping
        await api.ping()
      } catch(e) {
        // If it fails
        this.loading = false
        this.error = e.toString()
        return
      }

      // It succeeded \o/
      this.$store.dispatch('config/connected', api)
      this.$store.dispatch('config/save')

      this.loading = false
      this.success = true
    }
  },
  created() {
    this.port = this.$store.state.config.port
    this.ip = this.$store.state.config.ip
  }
}
</script>
<template>
  <v-form>
    <v-container>
      <v-layout row wrap>
        <v-flex xs12 md6>
          <v-text-field
            label="Host"
            v-validate="'required'"
            data-vv-name="ip"
            :error-messages="errors.collect('ip')"
            v-model="ip"
            :disabled="loading"
            required
          ></v-text-field>
        </v-flex>
        <v-flex xs12 md4>
          <v-text-field
            label="Port"
            v-validate="'required|between:1,65535'"
            data-vv-name="port"
            :error-messages="errors.collect('port')"
            v-model="port"
            :disabled="loading"
            required
          ></v-text-field>
        </v-flex>
        <v-flex xs12 md2>
          <v-spacer />
          <v-btn large :disabled="loading" :loading="loading" @click="connect" type="submit">Connect</v-btn>
        </v-flex>
        <v-alert :value="success" dismissible type="success" transition="scale-transition">
          Successfully connected to server.
        </v-alert>
        <v-alert :value="error" dismissible type="error" transition="scale-transition">
          {{ error }}
        </v-alert>
      </v-layout>
    </v-container>
  </v-form>
</template>
