<template>
  <div class="flex justify-center items-center min-h-screen">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-gray-800 mb-4">404 - Page Not Found</h1>
      <p class="text-gray-600">The page you are looking for does not exist.</p>
      <RouterLink to="/" class="text-blue-600 mt-4 inline-block">Go to Home</RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>