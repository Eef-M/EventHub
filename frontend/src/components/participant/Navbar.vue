<template>
  <nav class="bg-white shadow-md">
    <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
      <RouterLink to="/" class="text-2xl font-bold">
        <span class="text-purple-600">Event</span><span class="text-slate-600">Hub</span>
      </RouterLink>
      <div class="flex gap-6">
        <RouterLink to="/events" class="text-gray-700 hover:text-purple-600 transition">Events</RouterLink>
        <RouterLink to="/login" class="text-gray-700 hover:text-purple-600 transition">Login</RouterLink>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
