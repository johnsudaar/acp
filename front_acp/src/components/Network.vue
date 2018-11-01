<template>
  <div id="graph_container">
  </div>
</template>

<script>
import network from '../api/network';

export default {
  data() {
    return {
      width: 0,
      height: 0,
    }
  },
  mounted() {
    this.graph = new joint.dia.Graph;
    this.paper = new joint.dia.Paper({
      el: document.getElementById("graph_container"),
      model: this.graph,
      gridSize: 1,
    });
    network.getNetwork()
      .then(response =>{
        console.log(response.data);
      })
      .catch(error => {
        console.log(error)
      })
  },
  methods: {
    updateWindowWidth(event) {
      this.width = document.documentElement.clientWidth;
      this.updatePaperSize();
    },
    updateWindowHeight(event) {
      this.height = document.documentElement.clientHeight;
      this.updatePaperSize();
    },
    updatePaperSize() {
      this.paper.setDimensions(this.width, this.height);
    }
  },
}
</script>

<style scoped>

#graph_container {
  overflow: 'hidden';
  maw_height: 100%;
}
</style>
