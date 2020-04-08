<template>
  <v-layout row>
    <v-flex xs9>
      <ccu-single />
    </v-flex>
    <v-flex xs3>
      <ccu-overview name="Cam 1" color="green"/>
      <ccu-overview name="Cam 2"/>
      <ccu-overview name="Cam 3" color="red"/>
      <ccu-overview name="Cam 4"/>
    </v-flex>
  </v-layout>
</template>
