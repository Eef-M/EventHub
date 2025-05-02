import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Login from "../pages/Login.vue";
import Register from "../pages/Register.vue";
import EventDetail from "../pages/EventDetail.vue";
import Events from "../pages/Events.vue";

const routes = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/events', component: Events },
  { path: '/events/id/detail', component: EventDetail },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router