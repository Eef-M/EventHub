import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/participant/Home.vue";
import Login from "../pages/auth/Login.vue";
import Register from "../pages/auth/Register.vue";
import Events from "../pages/participant/Events.vue";
import EventDetail from "../pages/participant/EventDetail.vue";
import Dashboard from "../pages/organizer/Dashboard.vue";
import ManageEvents from "../pages/organizer/ManageEvents.vue";
import Registrations from "../pages/organizer/Registrations.vue";
import Feedback from "../pages/organizer/Feedback.vue";

const routes = [
  // participant
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/events', component: Events },
  { path: '/events/:id/detail', component: EventDetail },

  // organizer
  { path: '/organizer/dashboard', component: Dashboard },
  { path: '/organizer/manage-events', component: ManageEvents },
  { path: '/organizer/registrations', component: Registrations },
  { path: '/organizer/feedback', component: Feedback },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router