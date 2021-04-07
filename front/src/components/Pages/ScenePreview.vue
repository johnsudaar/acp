<template>
  <scene-view :elems="scene.elements"/>
</template>

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
