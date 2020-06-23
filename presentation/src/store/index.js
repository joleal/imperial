import Vue from "vue";
import Vuex from "vuex";
import ImperialApi from "@/api";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    openGames: [],
    activeGames: [],
    currentGame: {}
  },
  mutations: {
    updateOpenGamesList (state, openGamesList) {
      state.openGames = openGamesList;
    },
    updateActiveGamesList (state, activeGamesList) {
      state.activeGames = activeGamesList;
    },
    loadGame (state, game) {
      state.currentGame = game;
    }
  },
  actions: {
    async updateOpenGamesList (context) {
      const data = await ImperialApi.OpenGames();
      context.commit('updateOpenGamesList', data);
    },
    async updateActiveGamesList (context){
      const data = await ImperialApi.ActiveGames();
      context.commit('updateActiveGamesList', data);
    },
    updateState ({dispatch}) {
      dispatch('updateOpenGamesList');
      dispatch('updateActiveGamesList');
    },
    updateGame({dispatch}) {
      dispatch('updateGame');
    }
  },
  modules: {}
});
