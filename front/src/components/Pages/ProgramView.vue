<template>
  <v-layout column>
    <v-card>
      <v-card-title>
        Configuration
      </v-card-title>
      <v-card-text>
        <v-layout column>
          <v-text-field label="Title"/>
          <v-textarea label="Description" rows=2 />
        </v-layout>
      </v-card-text>
      <v-card-actions>
        <v-spacer/>
        <v-btn text> Save </v-btn>
      </v-card-actions>
    </v-card>

    <v-card class="mt-4">
      <v-card-title>
        Timeline
      </v-card-title>

      <v-card-text>
        <v-data-table
          :headers="headers"
          :items="test"
          disable-sort
          hide-default-footer
          >

          <template v-slot:item.name="props">
            <v-edit-dialog
              :return-value.sync="props.item.name"
              @save="save"
              @cancel="cancel"
              @open="open"
              @close="close"
            >
              {{ props.item.name }}
              <template v-slot:input>
                <v-text-field
                  v-model="props.item.name"
                  label="Edit"
                  single-line
                />
              </template>
            </v-edit-dialog>
          </template>
          <template v-slot:item.sort="{ items }">
            Toto
          </template>

          <template v-slot:item.actions="{ items }">
            Toto
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </v-layout>
</template>

<script>
  export default {
    data () {
      return {
        snack: false,
        snackColor: '',
        snackText: '',
        pagination: {},
        headers: [
          { text: "", value: "sort", sortable: false },
          { text: "Name", value: "name" },
          { text: "Start", value: "start" },
          { text: "Duration", value: "duration" },
          { text: "Notes", value: "notes" },
          { text: "Actions", value: "actions", sortable: false }
        ],
        defaultItem: {
          sort: '',
          actions: ''
        },
        test: [
          { name: "Preroll", start: "19:30", duration: "30mn", notes: "Media" },
          { name: "Intro", start: "20:00", duration: "5mn", notes: "Studio 1" },
          { name: "Video", start: "20:05", duration: "1mn", notes: "Media" },
        ]
      }
    },
    methods: {
      save () {
        this.snack = true
        this.snackColor = 'success'
        this.snackText = 'Data saved'
      },
      cancel () {
        this.snack = true
        this.snackColor = 'error'
        this.snackText = 'Canceled'
      },
      open () {
        this.snack = true
        this.snackColor = 'info'
        this.snackText = 'Dialog opened'
      },
      close () {
        console.log('Dialog closed')
      }
    }
  }
</script>
