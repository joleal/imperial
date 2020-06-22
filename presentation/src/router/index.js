import Vue from "vue";
import VueRouter from "vue-router";
import { authGuard } from "../auth/authGuard";
import Home from "../views/Home.vue";
import Game from "../views/Game.vue";
import Tournaments from "../views/Tournaments.vue";
import HallOfFame from "../views/HallOfFame.vue";
import Profile from "../views/Profile.vue";
import Create from '@/views/Create.vue'

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/game/:id",
    name: "Game",
    component: Game,
    beforeEnter: authGuard
  },
  {
    path: "/tournaments",
    name: "Tournaments",
    component: Tournaments,
    beforeEnter: authGuard
  },
  {
    path: "/hof",
    name: "Hall Of Fame",
    component: HallOfFame,
    beforeEnter: authGuard
  },
  {
    path: "/profile",
    name: "profile",
    component: Profile,
    beforeEnter: authGuard
  },
  {
    path: "/create",
    name: "create",
    component: Create,
    beforeEnter: authGuard
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
