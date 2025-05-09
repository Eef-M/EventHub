<template>
  <OrganizerLayout>
    <section class="p-6">
      <h1 class="text-2xl font-semibold mb-4">Dashboard</h1>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <div class="bg-white shadow rounded-xl p-4">
          <h2 class="text-lg font-medium text-gray-700">Total Events</h2>
          <p class="text-3xl font-bold text-purple-600">24</p>
        </div>
        <div class="bg-white shadow rounded-xl p-4">
          <h2 class="text-lg font-medium text-gray-700">Total Registrations</h2>
          <p class="text-3xl font-bold text-purple-600">152</p>
        </div>
        <div class="bg-white shadow rounded-xl p-4">
          <h2 class="text-lg font-medium text-gray-700">Pending Approvals</h2>
          <p class="text-3xl font-bold text-purple-600">5</p>
        </div>
        <div class="bg-white shadow rounded-xl p-4">
          <h2 class="text-lg font-medium text-gray-700">Feedback Received</h2>
          <p class="text-3xl font-bold text-purple-600">32</p>
        </div>
      </div>

      <div class="bg-white shadow rounded-xl p-6">
        <h2 class="text-xl font-semibold mb-4">Recent Registrations</h2>
        <ul class="divide-y divide-gray-100">
          <li class="py-3 flex justify-between">
            <span>John Doe</span>
            <span class="text-gray-500">Music Fest 2025</span>
          </li>
          <li class="py-3 flex justify-between">
            <span>Jane Smith</span>
            <span class="text-gray-500">Tech Expo 2025</span>
          </li>
          <li class="py-3 flex justify-between">
            <span>Albert Reyes</span>
            <span class="text-gray-500">Art Exhibition</span>
          </li>
        </ul>
      </div>
    </section>
  </OrganizerLayout>
</template>

<script setup lang="ts">
import OrganizerLayout from '../../layouts/OrganizerLayout.vue';


</script>
