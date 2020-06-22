import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

// Import the Auth0 configuration
import { domain, clientId, audience } from "../auth_config.json";

// Import the plugin here
import { Auth0Plugin, getInstance  } from "./auth";

import vuetify from "./plugins/vuetify";

// Install the authentication plugin here
Vue.use(Auth0Plugin, {
  domain,
  clientId,
  audience,
  onRedirectCallback: appState => {
    router.push(
      appState && appState.targetUrl
        ? appState.targetUrl
        : window.location.pathname
    );
  }
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify,
  created: function(){
    const instance = getInstance();
    instance.$watch("loading", async loading => {
      if (!loading && instance.isAuthenticated) { 
        this.$store.dispatch('updateState');
      }
    });
  },
  render: h => h(App)
}).$mount("#app");
