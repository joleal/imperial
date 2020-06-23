<template>
  <v-container fluid>
    <v-row> 
      <v-col cols="12" class="mb-4">
        <h2 class="display-2 font-weight-bold mb-3">
          Active Games
        </h2>
        

            <v-layout child-flex>
        <v-data-table dense
            height="400px"
            width="100%"
            fixed-header
            :headers="headers"
            :items="activeGames"
            :items-per-page="5"
            class="elevation-1"
        >
            <template v-slot:item.ID="{ item }">
                <v-btn text small color="primary" to="/about">{{item.ID}}</v-btn>
            </template>
            <template v-slot:item.players="{ item }">
                <div v-for="(player, i) in item.players" :key="i" >{{ player }}</div>
            </template>
        </v-data-table> 
        </v-layout>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <OpenGames></OpenGames>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';
import OpenGames from "@/components/OpenGames.vue";

export default {
  name: "Dashboard",
  components: {
    OpenGames
  },
  methods: {
    
  },
  computed: {
    ...mapState(['activeGames', 'openGames'])
  },
  data: () => ({
    headers:[
        {text: "Game Id", align:"start", sortable:"false", value: "ID"},
        {text: "Name", align:"start", sortable:"false", value: "name"},
        {text: "Players", align:"start", sortable:"false", value: "players"},
        {text: "Active Player", align:"start", sortable:"false", value: "currentPlayer"},
    ],
    games: [
        {gameID: 1, name: "game 1", players: ["joao","pedro","paulo","nuno"], activePlayer: "joao"},
        {gameID: 2, name: "game 2", players: ["joao","pedro","paulo"], activePlayer: "joao"},
        {gameID: 3, name: "game 3", players: ["joao","pedro","paulo","zé","nuno"], activePlayer: "joao"}
    ],
    result: "nothing"
  })
};
</script>