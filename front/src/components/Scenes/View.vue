<template>
  <div id="grid" >
    <div
      v-for="(elem) in elems"
      :key="elem.id"
      :style="styleFor(elem)"
      :class="elem.type"
      >
      <auto-text :size=256 :minSize="1" :text="elem.text" v-if="elem.type == 'text'" class="auto-text"/>
        <seven-seg :color="elem.color" :text="$store.getters['timers/timerValue'](elem.timer_id)" v-if="elem.type == 'timer'"/>
      <chat-widget :elem="elem" v-if="elem.type == 'chat'"/>
    </div>
  </div>
</template>

<style>
#grid {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
  grid-template-rows: 1fr 1fr 1fr 1fr 1fr 1fr;
  gap: 0px 0px;
  background-color: black;
}

.text {
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.auto-text {
  text-align: center;
  width: 100%;
}

.timer {
  display: flex;
  overflow: hidden;
  background-color: black !important;
}
</style>

<script>
export default {
  props: {
    elems: Array
  },
  methods: {
    styleFor(elem) {
      let style = `grid-area: ${elem.startY} / ${elem.startX} / span ${elem.sizeY} / span ${elem.sizeX}; `
      style += `background-color: ${elem.background_color}; `
      style += `color: ${elem.color}; `
      return style;
    }
  }
}
</script>
