<template>
  <v-container>
    <v-row class="text-center">
      
      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Play Imperial
        </h1>

        <p class="subheading font-weight-regular">
          Page designed to play the games Imperial and Imperial 2030 on the web
        </p>

        <!-- Check that the SDK client is not currently loading before accessing is methods -->
        <v-btn x-large color="success" dark @click="login" v-if="!$auth.loading">Login</v-btn>
        
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Landing",
  methods: {
    // Log the user in
    login() {
      this.$auth.loginWithRedirect();
    },
  },
  data: () => ({
    
  })
};
</script>
