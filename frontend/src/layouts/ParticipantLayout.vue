<template>
  <div>
    <Navbar />
    <main>
      <slot />
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Footer from '../components/participant/Footer.vue';
import Navbar from '../components/participant/Navbar.vue';
</script>
