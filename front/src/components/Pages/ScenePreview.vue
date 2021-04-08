<template>
<div class="scene-preview">
  <scene-view :elems="scene.elements"/>
</div>
</template>

<style scoped>
.scene-preview {
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
      scene: Object,
    }
  },
  async mounted() {
    this.$store.dispatch('timers/load')
    await this.$store.dispatch('scenes/load')
    let scene = this.$store.state.scenes.scenes[this.$route.params.id]
    this.scene = scene || {};
  }
}
</script>
