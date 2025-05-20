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
import Profile from "../pages/participant/Profile.vue";
import { useAuthStore } from "../stores/authStroe";
import { useUserStore } from "../stores/userStore";
import NotFound from "../pages/error/NotFound.vue";
import Forbidden from "../pages/error/Forbidden.vue";

const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/login',
    component: Login,
    meta: {
      allowAnonymous: true
    }
  },
  {
    path: '/register',
    component: Register,
    meta: {
      allowAnonymous: true
    }
  },
  {
    path: '/events',
    component: Events
  },
  {
    path: '/events/:id/detail',
    component: EventDetail
  },
  {
    path: '/organizer/dashboard',
    component: Dashboard,
    meta: {
      requiresAuth: true,
      requiresOrganizer: true
    }
  },
  {
    path: '/organizer/manage-events',
    component: ManageEvents,
    meta: {
      requiresAuth: true,
      requiresOrganizer: true
    }
  },
  {
    path: '/organizer/registrations',
    component: Registrations,
    meta: {
      requiresAuth: true,
      requiresOrganizer: true
    }
  },
  {
    path: '/organizer/feedback',
    component: Feedback,
    meta: {
      requiresAuth: true,
      requiresOrganizer: true
    }
  },
  {
    path: '/profile',
    component: Profile,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/403',
    name: 'forbidden',
    component: Forbidden
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: NotFound
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore();
  const userStore = useUserStore();

  if (to.name === 'forbidden' || to.name === 'not-found') {
    next();
    return;
  }

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!authStore.isAuthenticated) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      });
      return;
    }

    if (to.matched.some(record => record.meta.requiresOrganizer)) {
      if (userStore.userState.data && userStore.userState.data.role !== 'organizer') {
        next({ path: '/403' });
        return;
      }
    }
  } else if (to.meta.allowAnonymous) {
    if (authStore.isAuthenticated) {
      next({ path: '/' });
      return;
    }
  }

  next();
});

export default router