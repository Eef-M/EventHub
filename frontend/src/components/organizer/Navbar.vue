<template>
  <header class="bg-white shadow py-4 px-10 flex justify-between items-center">
    <h1 class="text-xl font-semibold">Dashboard</h1>
    <div class="relative" ref="dropdownRef">
      <span class="text-lg font-medium cursor-pointer hover:text-slate-500" @click="toggleDropdown">{{
        userStore.user?.username
      }}</span>
      <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-white border rounded shadow-md z-50">
        <ul>
          <RouterLink to="/profile">
            <li class="px-4 py-2 hover:bg-gray-100 cursor-pointer">Profile</li>
          </RouterLink>
          <li class="px-4 py-2 hover:bg-gray-100 cursor-pointer" @click="handleLogout">
            Logout
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script lang="ts" setup>
import { onClickOutside } from '@vueuse/core'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/userStore'
import { useAuthStore } from '../../stores/authStroe'

const dropdownOpen = ref(false)
const dropdownRef = ref(null)
const router = useRouter()
const userStore = useUserStore()
const authStore = useAuthStore()

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

onMounted(() => {
  userStore.getCurrentUser()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
  window.location.reload()
}

onClickOutside(dropdownRef, () => {
  dropdownOpen.value = false
})
</script>
