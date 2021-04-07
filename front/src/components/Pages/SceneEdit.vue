<template>
  <v-container>
    <v-row>
      <v-col cols="6">
        <div class="preview">
          <scene-view :elems="elems"/>
        </div>
      </v-col>
      <v-col cols="6">
        <v-card>
          <v-card-title> Edit </v-card-title>
          <v-alert :value="error" type="error" transition="scale-transition">
            {{ error }}
          </v-alert>

          <v-card-text>
            <v-flex>
              <v-expansion-panels accordion>
                <v-text-field
                  v-model="scene.name"
                  data-vv-name="name"
                  :error-messages="errors.collect('name')"
                  v-validate="'required'"
                  label="Name"
                  />
                <v-expansion-panel
                  v-for="(elem, i) in elems"
                  :key="elem.id"
                >
                  <v-expansion-panel-header>{{elem.name}}</v-expansion-panel-header>
                  <v-expansion-panel-content>
                    <v-form>
                      <v-flex xs12>
                        <scene-elem-form :item="elem"/>
                        <v-btn text color="red" @click="removeItem(i)"> Destroy </v-btn>
                      </v-flex>
                    </v-form>
                  </v-expansion-panel-content>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-flex>
            <v-btn text width="100%" @click="addItem"> <v-icon> mdi-plus </v-icon> </v-btn>
          </v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn text color="green" @click="submit()" :disabled="submitDisabled"> Save </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style>
.preview {
  position: fixed;
  margin: auto;
  width: 40vw;
  height: 40vh;
}
</style>

<script>
import TimersMixin from '@/mixins/timers'
import ChatMixin from '@/mixins/chat'

export default {
  mixins: [TimersMixin, ChatMixin],
  data() {
    return {
      scene: {},
      loading: false,
      error: null,
      elems: [
      ]
    }
  },
  async mounted() {
    this.$store.dispatch('timers/load')
    await this.$store.dispatch('scenes/load')
    let scene = this.$store.state.scenes.scenes[this.$route.params.id]
    this.scene = scene || {};
    this.elems = this.scene.elements || [];
  },
  computed: {
    submitDisabled() {
      return this.loading;
    }
  },
  methods: {
    addItem() {
      this.elems.push({
        id: this.$uuid.v4(),
        name: "NEW PANEL",
        type: "text",
        text: "NEW PANEL",
        startX: 1,
        startY: 1,
        sizeX: 1,
        sizeY: 1,
      })
    },
    removeItem(i) {
      this.elems.splice(i, 1)
    },
    async submit() {
      await this.$validator.reset();
      let valid = await this.$validator.validateAll();
      if(!valid){
        return
      }
      this.loading = true
      this.scene.elements = this.elems;

      try {
        if(!this.scene.id) {
          await this.$store.dispatch('scenes/create', this.scene)
        } else {
          await this.$store.dispatch('scenes/updateScene', { id: this.scene.id, params: this.scene })
        }
      } catch(e) {
        this.error = e.toString();
        this.loading = false;
        return
      }
      this.loading = false;
    }
  }
}
</script>
