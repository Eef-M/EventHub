<template>
  <DefaultLayout>
    <section class="py-20">
      <div class="max-w-7xl mx-auto px-4 text-center">
        <h1 class="text-5xl font-bold text-gray-900 mb-4">Discover Exciting Events</h1>
        <p class="text-lg text-gray-600 mb-8">
          Explore and manage event easily.
        </p>
        <div class="flex justify-center gap-4">
          <RouterLink to="/events"
            class="px-6 py-3 bg-blue-600 text-white rounded-xl shadow hover:bg-blue-700 transition">
            Browse Events
          </RouterLink>
        </div>
      </div>
    </section>

    <section class="bg-white py-16">
      <div class="max-w-7xl mx-auto px-4">
        <h2 class="text-3xl font-semibold text-center mb-10">News Events</h2>
        <div class="grid md:grid-cols-3 gap-6">
          <RouterLink to="/events/id/detail">
            <EventCard />
          </RouterLink>
        </div>
      </div>
    </section>
  </DefaultLayout>
</template>

<script setup lang="ts">
import EventCard from '../components/EventCard.vue';
import DefaultLayout from '../layouts/DefaultLayout.vue';

</script>
