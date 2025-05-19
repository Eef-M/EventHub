<script setup lang="ts">
import { Toaster } from '@/components/ui/sonner'
</script>

<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <Toaster richColors closeButton theme="dark" />
    <RouterView />
  </div>
</template>

<style scoped></style>
