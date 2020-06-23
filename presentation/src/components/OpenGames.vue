<template>
    <v-layout child-flex>
        
            <v-data-table dense
                height="400px"
                width="100%"
                fixed-header
                :headers="headers"
                :items="openGames"
                :items-per-page="5"
                class="elevation-1"
            >
                <template v-slot:top>
                    <v-toolbar flat color="white">
                        <v-toolbar-title>Open Games</v-toolbar-title>
                    </v-toolbar>
                </template>
                <template v-slot:item.ID="{ item }">
                    <div>{{ ('' + item.ID).slice(-5) }}</div>
                </template>

                <template v-slot:item.actions="{ item }">
                    <v-icon color="success" v-if="!inGame(item)" @click="console.log(item.ID)">mdi-account-multiple-plus</v-icon>
                    <v-icon color="error" v-if="inGame(item)" @click="console.log(item.ID)">mdi-account-multiple-minus</v-icon>
                </template>

                <template v-slot:item.players="{ item }" width=300>
                    <div class="d-flex justify-start">
                        <v-list-item v-for="(player, i) in item.players" :key="i" class="pa-0">
                            <v-list-item-avatar size="18">
                                <img :src="gravatar(player.email)">
                            </v-list-item-avatar>
                            <v-list-item-content>
                                <v-list-item-title v-html="player.nickname"></v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </div>
                </template>
            </v-data-table> 
        </v-layout>
</template>

<script>
import { mapState } from 'vuex';
import util from '@/util';

export default {
  name: "OpenGames",
  methods: {
    gravatar: util.gravatar,
    inGame: function(game){
        return game.players.filter(a => a.user_id == this.$auth.user.sub).length > 0
    }
  },
  computed: {
    ...mapState(['openGames'])
  },
  data: () => ({
    headers:[
        {text: "Game Id", align:"start", sortable:"false", value: "ID", width: "10px", fixed: true},
        {text: "Name", align:"start", sortable:"false", value: "name", width: "200px"},
        {text: "Players", align:"start", sortable:"false", value: "players", width: "600px", fixed: true},
        {text: "Actions", align:"end", sortable:"false", value:"actions"},
    ]
  })
};
</script>