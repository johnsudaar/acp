<template>
    <v-btn @click="open=true">
        <v-icon> mdi-plus </v-icon>
        <v-dialog
            :value="open"
            width="500"
            :persistent="true">
            <v-card>
                <v-card-title
                    class="headline"
                    primary-title>
                    Add a new position group
                </v-card-title>

                <v-card-text>
                    <v-form>
                        <v-text-field
                            v-model="name"
                            data-vv-name="name"
                            :error-messages="errors.collect('name')"
                            v-validate="'required'"
                            label="Group name"/>
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer />
                    <v-btn color="red darken-1" text @click="open=false">Close</v-btn>
                    <v-btn color="green darken-1" text :disabled="loading" :loading="loading" @click="submit()">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-btn>
</template>

<script>
export default {
    data() {
        return {
            name: null,
            open: false,
            error: null,
            loading: false,
        }
    },
    methods: {
        async submit() {
            await this.$validator.reset();
            let valid = await this.$validator.validateAll();
            if(!valid) {
                return
            }
            this.loading = true;
            let payload = {
                name: this.name,
            }

            try {
                await this.$store.dispatch('positiongroups/create', payload)
            } catch(e) {
                this.error = e.toString();
                this.loading = false;
                return
            }
            this.loading = false;
            this.open = false;
        }
    }
}
</script>