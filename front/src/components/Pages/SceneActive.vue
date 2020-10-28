<template>
  <scene-view :elems="scene.elements"/>
</template>

<script>
import TimersMixin from '@/mixins/timers'
export default {
  mixins: [TimersMixin],
  data() {
    return {
      scene: null,
      subscription: null,
    }
  },
  async mounted() {
    this.$store.dispatch('timers/load')
    this.subscription = this.$store.state.config.apiClient.realtime.subscribe("scene_active", this.onActiveSceneMessage)
    let active = await this.$store.state.config.apiClient.scenes.active()
    if(active.scene_id !== "") {
      this.onActiveSceneMessage({data: {data: active}})
    }
  },
  beforeDestroy() {
    if(this.subscription) {
      this.subscription.unsubscribe();
    }
  },
  methods: {
    async onActiveSceneMessage(message) {
      console.log(message)
      try {
        this.scene = await this.$store.state.config.apiClient.scenes.show(message.data.data.scene_id)
      } catch {

      }
    }
  }
}
</script>
