import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "@/stores/userStore";
import { useAuthStore } from "@/stores/authStore";
import Home from "@/views/participant/Home.vue";
import Login from "@/views/auth/Login.vue";
import Register from "@/views/auth/Register.vue";
import Events from "@/views/participant/Events.vue";
import EventDetail from "@/views/participant/EventDetail.vue";
import Dashboard from "@/views/organizer/Dashboard.vue";
import ManageEvents from "@/views/organizer/ManageEvents.vue";
import Ticket from "@/views/organizer/Ticket.vue";
import Registrations from "@/views/organizer/Registrations.vue";
import Feedback from "@/views/organizer/Feedback.vue";
import Profile from "@/views/participant/Profile.vue";
import EditProfile from "@/views/participant/EditProfile.vue";
import Forbidden from "@/views/error/Forbidden.vue";
import NotFound from "@/views/error/NotFound.vue";

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
    path: '/organizer/manage-tickets',
    component: Ticket,
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
    path: '/profile/edit',
    component: EditProfile,
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