<template>
  <group title="Cam 1" class="group">
    <div class="group-container">
      <v-layout column fill-height fluid>
        <v-flex xs8>
          <v-layout>
            <v-flex xs4>
              <group title="Iris">
                <center> 1.5 (manual) </center>
                <div class="slider">
                  <v-slider :min="-3" :max="3" v-model="irisSlider" ticks="always" track-color="transparent" always-dirty v-on:end="irisSliderEnd"/>
                </div>
              </group>
            </v-flex>
            <v-flex xs4>
              <group title="Gain">
                <center> 0db (manual) </center>
                <v-slider :min="-3" :max="3" v-model="irisSlider" ticks="always" track-color="transparent" always-dirty/>
              </group>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex xs4 d-flex>
          <v-layout>
            <v-flex xs4>
              <img src="http://wwp.acontreplongee.com/wp-content/uploads/2019/02/Structures.jpg"/>
            </v-flex>
            <v-flex xs4>
              <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmgIjjchem26nISlBiaIP8OVio0Jr6SYV4bJBaSmejZs7pqYdi"/>
            </v-flex>
            <v-flex xs4>
              <img src="https://programminghistorian.org/images/introduction-to-ffmpeg/vectorscope.png"/>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </div>
  </group>
</template>

<style >
.v-slider__track-fill.primary{
  background-color: transparent !important;
}
.slider {
  margin: 0 1rem;
}

.group{
  width: 100% !important;
  height: 100%;
}

.group-container {
  width: 100%;
  height: 100%;
  border: solid .5rem green;
  padding: .5rem;
}

img {
  width: 270px;
  height: 151px;
}
</style>

<script>
export default {
  data() {
    return {
      irisSlider: 0,
      gainSlider: 0,
    }
  },
  methods: {
    irisSliderEnd() {
      this.irisSlider = 0
    }
  }

}
</script>
