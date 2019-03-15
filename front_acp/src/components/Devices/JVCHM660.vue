<template>
  <v-layout column fill-width>
    <v-layout row class="no-grow" justify-center>
      <group title="Rec control">
        <v-layout align-end>
          <atem-btn class="mr-2" name="Rec"/>
          <div class="mr-1">
            <small> Record time: </small>
            <seven-seg color="red" text="00:00:00"/>
          </div>
        </v-layout>
      </group>
      <group title="Cards" class="ml-2" fill-height>
        <v-layout column fill-height justify-center>
          <div>
            Slot A: Not inserted
          </div>
          <div>
            Slot B: 120m remaining
          </div>
        </v-layout>
      </group>

      <group title="Tally" class="ml-2">
        <v-layout fill-height align-center>
          <led name="Off"/>
        </v-layout>
      </group>
    </v-layout>
    <v-layout fill-height>
      <h1> TODO ... </h1>
    </v-layout>
  </v-layout>
</template>

<script>

</script>
