<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" class="mb-4">
        <h2 class="display-2 font-weight-bold mb-3">
          Create Game
        </h2>
        <form>
          <v-text-field
            v-model="game.name"
            :error-messages="nameErrors"
            :counter="255"
            label="Name"
            required
            @input="$v.game.name.$touch()"
            @blur="$v.game.name.$touch()"
          ></v-text-field>
          <v-select
            v-model="game.settings.version"
            :error-messages="versionErrors"
            :items="versions"
            item-text="name"
            item-value="value"
            label="Game version"
            required
            @change="updateTaxFlag"
          ></v-select>
          <v-select
            v-model="game.settings.startingMode"
            :error-messages="startingModeErrors"
            :items="startingMode"
            item-text="name"
            item-value="value"
            label="Starting Mode"
            required
          ></v-select>
          <v-select
            v-model="game.numberOfPlayers"
            :items="[2, 3, 4, 5, 6]"
            label="Number of Players"
            required
            @change="$v.game.players.$touch()"
          ></v-select>
          <v-autocomplete
            v-model="game.players"
            :error-messages="playerErrors"
            :items="items"
            :loading="isLoading"
            :search-input.sync="search"
            item-text="nickname"
            :item-value="function(item){ return { user_id: item.user_id } }"
            label="Players"
            placeholder="Start typing to search"
            clearable
            multiple
            hide-no-data
            hide-selected
            prepend-icon="mdi-account-search"
            @change="$v.game.players.$touch()"
          >
            <template v-slot:selection="data">
              <v-chip
                v-bind="data.attrs"
                :input-value="data.selected"
                close
                @click="data.select"
                @click:close="remove(data.item)"
              >
                <v-avatar left>
                  <v-img :src="data.item.gravatar" />
                </v-avatar>
                {{ data.item.nickname }}
              </v-chip>
            </template>
            <template v-slot:item="data">
                <v-list-item-avatar>
                  <img :src="data.item.gravatar">
                </v-list-item-avatar>
                <v-list-item-content>
                  <v-list-item-title v-html="data.item.nickname"></v-list-item-title>
                  <v-list-item-subtitle v-html="data.item.email"></v-list-item-subtitle>
                </v-list-item-content>
              </template>
          </v-autocomplete>
          <h2>
            Options
          </h2>
            <v-checkbox
              v-model="game.settings.random"
              label="Random order"
            ></v-checkbox>
            <v-checkbox
              v-model="game.settings.investorCard"
              label="Investor card"
            ></v-checkbox>
            <v-checkbox
              v-model="game.settings.taxIncreaseBonus"
              label="Tax bonus only on taxation increase"
            ></v-checkbox>

          <v-btn @click="createGame" color="success" x-large dark>Create</v-btn>
        </form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ImperialApi from "@/api";
import util from "@/util";
import { validationMixin } from "vuelidate";
import { required, maxLength, minLength } from "vuelidate/lib/validators";

export default {
  name: "Dashboard",
  mixins: [validationMixin],
  data: () => ({
    game: {
      name: "",
      numberOfPlayers: 4,
      settings: {
        investorCard: true,
        taxIncreaseBonus: false,
      },
      players: [],
    },
    items: [],
    isLoading: false,
    search: null,
    valid: null,
    versions: [{name: "Imperial Original", value: "Original"},{name:"Imperial 2030", value:"Imperial 2030"}],
    startingMode: [{name: "Countries distributed randomly", value: "Random"},{name:"Auction", value:"Experienced"}]
  }),
  validations() {
    return {
      game: {
        name: { required , minLength: minLength(5), maxLength: maxLength(255), },
        players: { required , maxLength: maxLength(this.game.numberOfPlayers) },
        settings: { 
          startingMode: { required },
          version: {required}
        }
      }
    }
  },
  computed: {
    nameErrors() {
      const errors = [];
      if (!this.$v.game.name.$dirty) return errors;
      !this.$v.game.name.maxLength &&
        errors.push("Name must be at most 255 characters long");
      !this.$v.game.name.minLength &&
        errors.push("Name must be at least 5 characters long");
      !this.$v.game.name.required && errors.push("Name is required.");
      return errors;
    },
    playerErrors() {
      const errors = [];
      !this.$v.game.players.maxLength &&
        errors.push(`Can't exceed the number of players`);
      return errors;
    },
    versionErrors() {
      const errors = [];
      if (!this.$v.game.settings.version.$dirty) return errors;
      !this.$v.game.settings.version.required &&
        errors.push('Version is required.');
      return errors;
    },
    startingModeErrors() {
      const errors = [];
      if (!this.$v.game.settings.startingMode.$dirty) return errors;
      !this.$v.game.settings.startingMode.required &&
        errors.push('Starting mode is required.');
      return errors;
    }
  },
  methods: {
    async createGame() {
      this.$v.$touch()
      if (this.$v.$invalid) {
        this.submitStatus = 'ERROR'
      } else {
        try {
          await ImperialApi.CreateGame(this.game);
          this.$store.dispatch('updateState');
          this.$router.push({ name: "Home" });
        } catch (err) {
          console.log(err);
        }
        
      }
    },
    clear() {},
    remove(item) {
      this.game.players = this.game.players.filter(p => p.user_id != item.user_id);
    },
    updateTaxFlag(){
      this.game.settings.taxIncreaseBonus = this.game.settings.version == "Original";
    }
  },
  mounted(){
    const user = Object.assign({}, this.$auth.user, {user_id: this.$auth.user.sub, gravatar: util.gravatar(this.$auth.user.email)});
    this.items.push(user);
    this.game.players.push(user);

    ImperialApi.Players("")
    .then((res) => {
    this.items = res.map(i => Object.assign({}, i, {gravatar: util.gravatar(i.email)}));
    })
    .catch((err) => {
    console.log(err);
    })
    .finally(() => (this.isLoading = false));
  },
};
</script>
