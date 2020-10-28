<template>
  <v-row align="stretch">
    <v-col sm="6" md="4" v-for="(scene) in $store.state.scenes.scenes" :key="scene.id">
      <v-card>
        <v-card-title> {{scene.name}} </v-card-title>
        <v-card-actions>
          <v-flex xs12>
            <v-btn
              @click="launch(scene.id)"
              text>
              Launch
            </v-btn>
            <v-btn
              :to="{ name:'scene_preview', params: { id: scene.id } }"
              color="blue"
              text>
              Preview
            </v-btn>

            <v-btn
              :to="{ name:'scene_edit', params: { id: scene.id } }"
              color="green"
              text>
              Edit
            </v-btn>
            <v-btn
              color="red"
              @click="destroy(scene)"
              text>
              Destroy
            </v-btn>
          </v-flex>
        </v-card-actions>
      </v-card>
    </v-col>
    <v-col sm="6" md="4">
      <v-card height="100%">
        <v-btn
          height="100%"
          width="100%"
          :to="{ name:'scene_edit', params: { id: 'new' } }"
          text>
          <v-icon> mdi-plus </v-icon>
        </v-btn>
      </v-card>
    </v-col>
    <v-dialog
      :value="openDelete"
      width="500"
      :persistent="true"
      >
      <v-card v-if="scene">
        <v-card-title> Delete {{scene.name}} ?</v-card-title>
        <v-card-text> Do you really want to destroy {{scene.name}} ? </v-card-text>
        <v-card-actions>
          <v-btn text @click="onCloseDelete" color="green"> Abort </v-btn>
          <v-btn text @click="reallyDestroy" color="red"> Destroy </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import TimersMixin from '@/mixins/timers'

export default {
  mixins: [TimersMixin],
  data() {
    return {
      scene: null,
      openDelete: false,
    }
  },
  mounted() {
    this.$store.dispatch('timers/load')
    this.$store.dispatch('scenes/load')
  },
  methods: {
    destroy(scene) {
      this.scene = scene;
      this.openDelete = true;
    },
    onCloseDelete() {
      this.scene = null;
      this.openDelete = false;
    },
    async reallyDestroy() {
      await this.$store.dispatch('scenes/destroy', this.scene.id)
      this.openDelete = false;
      this.scene=  null;
    },
    async launch(id) {
      await this.$store.state.config.apiClient.scenes.launch(id)
    }
  }
}

</script>
