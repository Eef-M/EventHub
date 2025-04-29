<template>
  <div>
    <Navbar />
    <main>
      <slot />
    </main>
    <footer class="bg-gray-100 text-center text-sm py-4 mt-10">
      &copy; 2025 Event Management System. All rights reserved.
    </footer>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue';
</script>
