<template>
  <div class="scene-active">
    <scene-view :elems="scene.elements" />
  </div>
</template>

<style scoped>
.scene-active {
  width: 100vw;
  height: 100vh;
}
</style>

<script>
import TimersMixin from '@/mixins/timers'
import ChatMixin from '@/mixins/chat'

export default {
  mixins: [TimersMixin, ChatMixin],
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
