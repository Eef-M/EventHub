<template>
  <div class="flex justify-center items-center min-h-screen">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-red-600 mb-4">403 - Forbidden</h1>
      <p class="text-gray-700">You do not have permission to access this page.</p>
      <RouterLink to="/" class="text-blue-600 mt-4 inline-block">Go to Home</RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>