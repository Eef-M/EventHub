<template>
  <nav class="bg-white shadow-md">
    <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
      <RouterLink to="/" class="text-2xl font-bold text-indigo-600">Events</RouterLink>
      <div class="flex gap-6">
        <RouterLink to="/events" class="text-gray-700 hover:text-indigo-600 transition">Events</RouterLink>
        <RouterLink to="/login" class="text-gray-700 hover:text-indigo-600 transition">Login</RouterLink>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
