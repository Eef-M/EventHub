<template>
  <div class="flex h-screen bg-gray-100">
    <Sidebar />
    <div class="flex-1 flex flex-col">
      <Navbar />
      <main class="flex-1 overflow-y-auto p-4">
        <slot />
      </main>
      <Footer />
    </div>
  </div>
</template>

<script lang="ts" setup>
import Footer from '../components/organizer/Footer.vue';
import Navbar from '../components/organizer/Navbar.vue';
import Sidebar from '../components/organizer/Sidebar.vue';

</script>