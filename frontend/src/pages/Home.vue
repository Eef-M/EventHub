<template>
  <DefaultLayout>
    <section class="bg-zinc-100 py-12 px-4 md:px-8">
      <div class="max-w-7xl mx-auto flex flex-col md:flex-row items-center justify-between">
        <div class="text-left mb-8 md:mb-0 md:w-1/2">
          <h1 class="text-3xl md:text-4xl lg:text-5xl font-bold text-purple-900 mb-4">Discover Exciting Events</h1>
          <p class="text-base text-gray-600 mb-6">
            Explore and manage events easily. Find events that match your interests and register in just a few clicks.
          </p>
          <button
            class="px-6 py-3 bg-purple-600 text-white rounded-md shadow hover:bg-purple-700 transition cursor-pointer">
            Browse Events
          </button>
        </div>

        <div class="md:w-1/2 relative">
          <div
            class="bg-purple-100 rounded-full w-64 h-64 md:w-80 md:h-80 mx-auto relative flex items-center justify-center">
            <img
              src="https://images.pexels.com/photos/2952834/pexels-photo-2952834.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
              alt="Person using smartphone" class="w-48 h-48 md:w-64 md:h-64 object-cover rounded-full" />
          </div>
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
          <RouterLink to="/events/id/detail">
            <EventCard />
          </RouterLink>
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
