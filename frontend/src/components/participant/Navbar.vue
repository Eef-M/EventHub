<template>
  <nav class="bg-white shadow-md">
    <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
      <RouterLink to="/" class="text-2xl font-bold">
        <span class="text-purple-600">Event</span><span class="text-slate-600">Hub</span>
      </RouterLink>
      <div class="flex gap-6">
        <RouterLink to="/events" class="text-gray-700 hover:text-purple-600 transition">Events</RouterLink>
        <template v-if="userStore.user">
          <details class="relative">
            <summary class="cursor-pointer text-gray-700 hover:text-purple-600 font-medium">
              {{ userStore.user.username }}
            </summary>
            <ul class="absolute right-0 bg-white shadow-md rounded mt-2 w-40 z-50">
              <template v-if="userStore.user.role === 'organizer'">
                <li>
                  <RouterLink to="/organizer/dashboard" class="block px-4 py-2 hover:bg-gray-100 cursor-pointer">
                    Dashboard</RouterLink>
                </li>
              </template>
              <li>
                <RouterLink to="/profile" class="block px-4 py-2 hover:bg-gray-100 cursor-pointer">Profile</RouterLink>
              </li>
              <li>
                <button @click="handleLogout" class="w-full text-left px-4 py-2 hover:bg-gray-100 cursor-pointer">
                  Logout
                </button>
              </li>
            </ul>
          </details>
        </template>
        <template v-else>
          <RouterLink to="/login" class="text-gray-700 hover:text-purple-600 font-medium">Login</RouterLink>
        </template>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useUserStore } from '../../stores/userStore';
import { onMounted } from 'vue';
import { useAuthStore } from '../../stores/authStroe';
import router from '../../router';

const userStore = useUserStore()
const authStore = useAuthStore()

onMounted(() => {
  userStore.getCurrentUser()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
  window.location.reload()
}
</script>
