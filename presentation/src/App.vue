<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" :clipped="$vuetify.breakpoint.lgAndUp" app v-if="$auth.isAuthenticated" 
          expand-on-hover
      permanent>
      <v-list dense>
        <v-list-item link to="/">
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>
              Dashboard
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        
          <v-list-group
            no-action
            >
            <template v-slot:prependIcon>
              <v-badge color="pink" left overlap dot offset-x="0" offset-y="0" v-if="activeGames.length > 0" >
                    <v-icon>mdi-puzzle</v-icon>
                    </v-badge>
              </template>

            <template v-slot:activator>
              
              <v-list-item-title>
                Games
              </v-list-item-title>       
            </template>
            <v-list-item 
              v-for="(game, i) in activeGames" :key="i" link :to="`/game/${game.ID}`">
              <v-list-item-title v-text="game.name"></v-list-item-title>
              <v-list-item-action>
                <v-icon color="pink lighten-1">mdi-play</v-icon>
              </v-list-item-action>
            </v-list-item>
            <v-list-item link to="/create">
              <v-list-item-title>Create game</v-list-item-title>
              <v-list-item-action>
                <v-icon color="blue lighten-1">mdi-plus</v-icon>
              </v-list-item-action>
            </v-list-item>
          </v-list-group>
        <v-list-item link to="/tournaments">
          <v-list-item-action>
            <v-icon>mdi-tournament</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>
              Tournaments
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link to="/hof">
          <v-list-item-action>
            <v-icon>mdi-trophy</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>
              Hall of Fame
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar :clipped-left="$vuetify.breakpoint.lgAndUp" app color="green darken-3" dark dense v-if="$auth.isAuthenticated">
      <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
        <span class="hidden-sm-and-down">Imperial</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-menu v-model="menu" :close-on-content-click="false" :nudge-width="200" offset-x offset-y> 
        <template v-slot:activator="{ on, attrs }">
          <v-badge color="pink" overlap dot offset-x="18" offset-y="20">
            <v-btn  v-bind="attrs" v-on="on" icon>
              <v-icon>mdi-bell</v-icon>
            </v-btn>
          </v-badge>
        </template>
        <v-card>
          <v-list dense>
            <v-list-item-group>
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>Game #abcde</v-list-item-title>
                  <v-list-item-subtitle>Founder of Vuetify.js</v-list-item-subtitle>
                </v-list-item-content>

                <v-list-item-action>
                  <v-list-item-action-text v-text="'15min'"></v-list-item-action-text>
                </v-list-item-action>
              </v-list-item>
              <v-divider></v-divider>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-menu>
      <v-menu v-model="menu" :close-on-content-click="false" :nudge-width="200" offset-x offset-y> 
        <template v-slot:activator="{ on, attrs2 }">
          <v-btn  v-bind="attrs2" v-on="on" icon large>
            <v-avatar size="32px" item>
              <v-img :src="$auth.user.picture" />
            </v-avatar>
          </v-btn>
        </template>
        <v-card>
          <v-list dense>
            <v-list-item-group>
              <v-list-item>
                Profile
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item @click="logout()">
                Logout
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-menu>
      
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>
<script>
import { mapState } from 'vuex';

export default {
  name: "App",
  props: {
    source: String
  },
  computed: {
    ...mapState(['activeGames'])
  },
  components: {},
  methods: {
    // Log the user out
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin
      });
    }
  },
  data: () => ({
    drawer: true,
    mini: true,
    menu: false
  })
};
</script>
