import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/user/Home.vue";
import Login from "../pages/auth/Login.vue";
import Register from "../pages/auth/Register.vue";
import Events from "../pages/user/Events.vue";
import EventDetail from "../pages/user/EventDetail.vue";
import Dashboard from "../pages/admin/Dashboard.vue";
import ManageEvents from "../pages/admin/ManageEvents.vue";
import Registrations from "../pages/admin/Registrations.vue";
import Feedback from "../pages/admin/Feedback.vue";

const routes = [
  // user
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/events', component: Events },
  { path: '/events/id/detail', component: EventDetail },

  // admin
  { path: '/admin/dashboard', component: Dashboard },
  { path: '/admin/manage-events', component: ManageEvents },
  { path: '/admin/registrations', component: Registrations },
  { path: '/admin/feedback', component: Feedback },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router