<template>
  <DefaultLayout>
    <section class="py-8 max-w-7xl mx-auto">
      <!-- Event Banner -->
      <div class="rounded-2xl overflow-hidden shadow-md mb-8">
        <img
          src="https://images.pexels.com/photos/3861969/pexels-photo-3861969.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
          alt="Event Banner" class="w-full h-64 object-cover" />
      </div>

      <!-- Event Info -->
      <div class="space-y-4 mb-10">
        <h1 class="text-3xl font-bold text-gray-900">Event Title Placeholder</h1>
        <p class="text-gray-600">Category: <span class="font-medium">Conference</span></p>
        <p class="text-gray-600">Location: <span class="font-medium">Jakarta, Indonesia</span></p>
        <p class="text-gray-600">Date: <span class="font-medium">2025-08-30</span></p>
        <p class="text-gray-600">Time: <span class="font-medium">14:00 - 17:00</span></p>
      </div>

      <!-- Description -->
      <div class="mb-10">
        <h2 class="text-xl font-semibold mb-2">Description</h2>
        <p class="text-gray-700">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc vehicula, justo at malesuada tempor,
          velit nisl facilisis odio, sed sodales elit mauris id leo.
        </p>
      </div>

      <!-- Ticket Options -->
      <div class="mb-10">
        <h2 class="text-xl font-semibold mb-4">Available Tickets</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="border p-4 rounded-xl shadow-sm">
            <h3 class="text-lg font-bold">General Admission</h3>
            <p class="text-gray-600">Price: Rp 150.000</p>
            <p class="text-sm text-gray-500">Available: 100</p>
            <button class="mt-3 w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-xl">Buy Ticket</button>
          </div>
          <div class="border p-4 rounded-xl shadow-sm">
            <h3 class="text-lg font-bold">VIP</h3>
            <p class="text-gray-600">Price: Rp 350.000</p>
            <p class="text-sm text-gray-500">Available: 25</p>
            <button class="mt-3 w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-xl">Buy Ticket</button>
          </div>
        </div>
      </div>

      <!-- Feedback Section (if applicable) -->
      <div>
        <h2 class="text-xl font-semibold mb-2">Feedback</h2>
        <p class="text-gray-500 italic">No feedback yet.</p>
      </div>
    </section>
  </DefaultLayout>
</template>

<script lang="ts" setup>
import DefaultLayout from '../layouts/DefaultLayout.vue';

</script>
