<template>
<v-btn @click="open=true">
    <v-icon> mdi-delete </v-icon>
    <v-dialog
        :value="open"
        width="500"
        :persistent="true">
        <v-card>
            <v-card-title
                class="headline"
                primary-title>
                Delete a position group
            </v-card-title>
            <v-card-text>
                <v-select
                    :items="$store.getters['positiongroups/asArray']"
                    item-text="name"
                    item-value="id"
                    v-model="selectedPositionGroup"
                    label="Position Group to delete"
                    dense hide-details/>
                <div v-if="selectedPositionGroup !== null">
                    <p class="mt-3"> There is {{positionCount}} positions in this category. Do you want to delete them too? </p>
                    <p> If you choose not to delete your positions, they will be moved to the default group. </p>

                    <v-checkbox
                        v-model="destroyMembers"
                        label="Yes, destroy members of this group"/>
                </div>

            </v-card-text>
            <v-card-actions>
                <v-spacer />
                    <v-btn color="red darken-1" text @click="open=false">Close</v-btn>
                    <v-btn color="green darken-1" text :disabled="disabled" :loading="loading" @click="submit()"> Save </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</v-btn>
</template>


<script>
export default {
    data() {
        return {
            open: false,
            error: null,
            loading: false,
            selectedPositionGroup: null,
            destroyMembers: false,
        }
    },
    methods: {
        async submit() {
            this.loading = true;
            try {
                await this.$store.dispatch('positiongroups/destroy', {
                    id: this.selectedPositionGroup,
                    destroyMembers: this.destroyMembers,
                })
            } catch(e) {
                this.loading = false;
                return
            }

            let devices = this.$store.state.devices.devices;
            for(let deviceId in devices) {
                if(devices[deviceId].types.includes("ptz")) {
                    await this.$store.dispatch('ptzpositions/refresh', devices[deviceId].id);
                }
            }
            this.loading = false;
            this.open = false;

        }
    },
    computed: {
        disabled() {
            return this.selectedPositionGroup == null || this.loading
        },
        positionCount() {
            let positions = this.$store.state.ptzpositions.positions;

            let count = 0;
            for(let cam in positions) {
                for(let posIdx in positions[cam]) {
                    if(positions[cam][posIdx].position_group_id == this.selectedPositionGroup) {
                        count++
                    }
                }
            }
            return count;
        },
    }
}
</script>
