<template>
  <div class="min-h-screen flex flex-col">
    <Navbar />
    <div class="flex flex-1 overflow-hidden">
      <Sidebar />
      <main class="flex-1 overflow-y-auto p-6 bg-muted/50">
        <slot />
      </main>
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Footer from '@/components/organizer/Footer.vue';
import Navbar from '@/components/organizer/Navbar.vue';
import Sidebar from '@/components/organizer/Sidebar.vue';

</script>