<template>
  <UserLayout>
    <section class="py-8 max-w-7xl mx-auto">
      <div class="rounded-2xl overflow-hidden shadow-md mb-8">
        <img :src="event?.bannerURL" alt="Event Banner" class="w-full h-64 object-cover" />
      </div>

      <div class="space-y-4 mb-10">
        <h1 class="text-3xl font-bold text-gray-900">{{ event?.title }}</h1>
        <p class="text-gray-600">Category: <span class="font-medium">{{ event?.category }}</span></p>
        <p class="text-gray-600">Location: <span class="font-medium">{{ event?.location }}</span></p>
        <p class="text-gray-600">Date: <span class="font-medium">{{ event?.date }}</span></p>
        <p class="text-gray-600">Time: <span class="font-medium">14:00 - 17:00</span></p>
      </div>

      <div class="mb-10">
        <h2 class="text-xl font-semibold mb-2">Description</h2>
        <p class="text-gray-700">
          {{ event?.description || 'No description available.' }}
        </p>
      </div>

      <div class="mb-10">
        <h2 class="text-xl font-semibold mb-4">Available Tickets</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="border p-4 rounded-xl shadow-sm">
            <h3 class="text-lg font-bold">General Admission</h3>
            <p class="text-gray-600">Price: Rp 150.000</p>
            <p class="text-sm text-gray-500">Available: 100</p>
            <button class="mt-3 w-full bg-purple-600 hover:bg-purple-700 text-white py-2 rounded-xl cursor-pointer">Buy
              Ticket</button>
          </div>
          <div class="border p-4 rounded-xl shadow-sm">
            <h3 class="text-lg font-bold">VIP</h3>
            <p class="text-gray-600">Price: Rp 350.000</p>
            <p class="text-sm text-gray-500">Available: 25</p>
            <button class="mt-3 w-full bg-purple-600 hover:bg-purple-700 text-white py-2 rounded-xl cursor-pointer">Buy
              Ticket</button>
          </div>
        </div>
      </div>

      <div>
        <h2 class="text-xl font-semibold mb-2">Feedback</h2>
        <p class="text-gray-500 italic">No feedback yet.</p>
      </div>
    </section>
  </UserLayout>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import UserLayout from '../../layouts/UserLayout.vue';
import { dummyEvents } from '../../data/events';

const route = useRoute();
const eventId = route.params.id as string;
const event = computed(() => dummyEvents.find(e => e.id === eventId));
</script>
