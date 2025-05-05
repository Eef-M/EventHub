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
import Footer from '../components/user/Footer.vue';
import Navbar from '../components/user/Navbar.vue';
</script>
