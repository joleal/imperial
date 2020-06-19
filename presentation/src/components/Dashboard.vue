<template>
  <v-container fluid>
    <v-row> 
      <v-col cols="12" class="mb-4">
        <h2 class="display-2 font-weight-bold mb-3">
          Active Games
        </h2>
        

            <v-layout child-flex>
        <v-data-table
            height="400px"
            width="100%"
            fixed-header
            :headers="headers"
            :items="games"
            :items-per-page="5"
            class="elevation-1"
        >
            <template v-slot:item.gameID="{ item }">
                <v-btn text small color="primary" to="/about">{{item.gameID}}</v-btn>
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
        <h2>test</h2>
        <v-btn @click="testAPI">Test asd asd as</v-btn>
        <div>{{ result }}</div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios"

export default {
  name: "Dashboard",
  methods: {
    async testAPI() {
      const token = await this.$auth.getTokenSilently();
      const options = {
          url: `${process.env.VUE_APP_BASE_URI}/user/register`, 
          method: 'POST',
          headers: {
                "Authorization": `Bearer ${token}`,
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8'
              },
          data: {
                nickname: this.$auth.user.nickname,
                email: this.$auth.user.email
              }
      };
      const response = await axios(options);
      this.$set(this, 'result', response);
    }
    
  },
  data: () => ({
    headers:[
        {text: "Game Id", align:"start", sortable:"false", value: "gameID"},
        {text: "Name", align:"start", sortable:"false", value: "name"},
        {text: "Players", align:"start", sortable:"false", value: "players"},
        {text: "Active Player", align:"start", sortable:"false", value: "activePlayer"},
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