<template>
  <header class="bg-white shadow p-4 flex justify-between items-center">
    <h1 class="text-xl font-semibold">Dashboard</h1>
    <div class="relative" ref="dropdownRef">
      <img src="https://i.pravatar.cc/40" alt="User Avatar" class="w-10 h-10 rounded-full cursor-pointer border"
        @click="toggleDropdown" />
      <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-white border rounded shadow-md z-50">
        <ul>
          <li class="px-4 py-2 hover:bg-gray-100 cursor-pointer">Profile</li>
          <li class="px-4 py-2 hover:bg-gray-100 cursor-pointer" @click="logout">
            Logout
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script lang="ts" setup>
import { onClickOutside } from '@vueuse/core'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const dropdownOpen = ref(false)
const dropdownRef = ref(null)
const router = useRouter()

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function logout() {
  // Clear auth logic here
  router.push('/login') // Temporary
}

onClickOutside(dropdownRef, () => {
  dropdownOpen.value = false
})
</script>
