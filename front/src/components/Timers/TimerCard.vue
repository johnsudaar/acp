<template>
  <v-card height="100%">
    <v-card-title>
      <v-layout row>
        <v-flex>
          {{ timer.name }}
        </v-flex>
        <v-spacer/>
        <v-btn color="green" text @click="edit"> Edit </v-btn>
        <v-btn color="red" text @click="destroy"> Delete </v-btn>
      </v-layout>
    </v-card-title>
    <v-card-text>
      <v-layout row>
        <seven-seg color="red" :text="$store.getters['timers/timerValue'](timer.id)"/>
      </v-layout>
    </v-card-text>
    <v-card-actions v-if="timer.type == 'countdown' || timer.type == 'countup'">
      <v-btn @click="reset" text> Reset </v-btn>
      <v-btn @click="start" text> Start </v-btn>
      <v-btn @click="pause" text> Pause </v-btn>
    </v-card-actions>
   <v-dialog
      :value="openEdit"
      width="500"
      :persistent="true"
      >
      <timer-form @close="onCloseEdit" :timer="timer"/>
    </v-dialog>
     <v-dialog
       :value="openDelete"
       width="500"
       :persistent="true"
       >
       <v-card>
         <v-card-title> Delete {{timer.name}} ?</v-card-title>
         <v-card-text> Do you really want to destroy {{timer.name}} ? </v-card-text>
         <v-card-actions>
           <v-btn text @click="onCloseDelete" color="green"> Abort </v-btn>
           <v-btn text @click="reallyDestroy" color="red"> Destroy </v-btn>
         </v-card-actions>
       </v-card>
     </v-dialog>
  </v-card>
</template>

<script>
export default {
  props: {
    timer: Object,
  },
  data() {
    return {
      openEdit: false,
      openDelete: false,
    }
  },
  methods: {
    reset() {
      this.action("reset")
    },
    start() {
      this.action("start")
    },
    pause() {
      this.action("pause")
    },
    async action(name, params) {
      this.$store.state.config.apiClient.timers.action(this.timer.id, name, params);
    },
    edit() {
      this.openEdit = true;
    },
    onCloseEdit() {
      this.openEdit = false;
    },
    destroy() {
      this.openDelete = true;
    },
    onCloseDelete() {
      this.openDelete = false;
    },
    async reallyDestroy() {
      await this.$store.dispatch('timers/destroy', this.timer.id)
      this.openDelete = false;
    }
  }
}
</script>
