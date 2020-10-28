<template>
  <v-card>
    <v-card-title
      class="headline"
      primary-title>
      <div v-if="timer == null">
        Add new timer
      </div>
      <div v-else>
        Edit time {{timer.name}}
      </div>
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-layout wrap>
          <v-flex xs12>
            <v-text-field
              v-model="name"
              data-vv-name="name"
              :error-messages="errors.collect('name')"
              v-validate="'required'"
              label="Timer name"/>
          </v-flex>

          <v-flex xs12>
             <v-select
                v-model="type"
                :items="clockTypes"
                item-text="name"
                item-value="type"
                label="Clock type"
                single-line/>
          </v-flex>

          <v-flex xs12
            v-if="type != 'external' && type !='clock'"
            >
            <v-text-field
              v-model="duration"
              data-vv-name="duration"
              :error-messages="errors.collect('duration')"
              label="Duration"
              v-validate="'required|regex:^([0-9]+h)?([0-9]+m)?([0-9]+s)?$'"
              />
          </v-flex>

          <v-flex xs12
            v-if="type == 'external'"
            >
            <v-flex xs12>
                <v-select
                  v-model="external_device"
                  data-vv-name="external_device"
                  v-validate="'required'"
                  :error-messages="errors.collect('external_device')"
                  :items="timeable"
                  item-text="name"
                  item-value="id"
                  label="External Device"
                  single-line/>
            </v-flex>
            <v-flex xs12>
                <v-select
                  v-model="external_source"
                  data-vv-name="external_source"
                  v-validate="'required'"
                  :error-messages="errors.collect('external_source')"
                  :items="timerSources"
                  item-text="name"
                  item-value="id"
                  label="External Source"
                  single-line/>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-form>
     <v-alert :value="error" type="error" transition="scale-transition">
        {{ error }}
      </v-alert>

    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="red darken-1" text @click="close()">Close</v-btn>
      <v-btn color="green darken-1" text :disabled="submitDisabled" :loading="loading" @click="submit()">Save</v-btn>
    </v-card-actions>
  </v-card>

</template>

<script>
export default {
  props: {
    timer: Object,
  },
  data() {
    return {
      clockTypes: [
        { name: "Clock", type: "clock" },
        { name: "Count Down", type: "countdown" },
        { name: "Count Up", type: "countup" },
        { name: "External", type: "external" },
      ],
      type: 'clock',
      name: null,
      duration: null,
      external_device: null,
      external_source: null,
      timerSources: [],
      loading: false,
      error: null,
    }
  },
  async mounted() {
    if(this.timer) {
      this.name = this.timer.name
      this.type = this.timer.type
      this.duration = this.timer.duration
      this.external_device = this.timer.external_device
      await this.loadTimerSources()
      this.external_source = this.timer.external_source
    }
  },
  watch: {
   async external_device() {
      this.loadTimerSources()
    },
  },
  computed: {
    timeable() {
      return this.$store.state.devices.devices.filter((device) => {
        return device.types.includes("timer")
      })
    },
    submitDisabled() {
      return this.loading;
    }
  },

  methods: {
    async loadTimerSources() {
      this.timerSources = []
      let result = []
      let res = await this.$store.state.config.apiClient.timers.deviceSources(this.external_device)
      this.timerSources = res.sources.map((name) => {
        return { name: name, id: name }
      })
    },
    copyParams() {
      if(!this.timer) {
        return
      }

      this.name = this.timer.name
      this.type = this.timer.type
      this.duration = this.timer.duration
    },
    async submit() {
      await this.$validator.reset();
      let valid = await this.$validator.validateAll();
      if(!valid){
        return
      }

      this.loading = true;
      let payload = {
        name: this.name,
        type: this.type,
        duration: this.duration,
        external_device: this.external_device,
        external_source: this.external_source,
      }
      try {
        if(!this.timer) {
          await this.$store.dispatch('timers/create', payload)
        } else {
          await this.$store.dispatch('timers/updateTimer', {id: this.timer.id, params: payload})
        }
      } catch(e) {
        this.error = e.toString();
        this.loading = false;
        return
      }
      this.loading = false;
      this.$emit('close')
    },
    close() {
      this.$emit('close')
    }
  }
}
</script>
