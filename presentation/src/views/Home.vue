<template>
  <div class="home">
    <Dashboard v-if="$auth.isAuthenticated"/>
    <Landing v-if="!$auth.isAuthenticated" />
  </div>
</template>

<script>
// @ is an alias to /src
import Dashboard from "@/components/Dashboard.vue";
import Landing from "@/components/Landing.vue";

export default {
  name: "home",
  components: {
    Dashboard,
    Landing
  },
  methods: {
    // Log the user in
    login() {
      this.$auth.loginWithRedirect();
    },
    // Log the user out
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin
      });
    }
  }
};
</script>
