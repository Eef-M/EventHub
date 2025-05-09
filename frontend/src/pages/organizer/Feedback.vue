<template>
  <OrganizerLayout>
    <div class="p-6">
      <h2 class="text-2xl font-semibold mb-4">Event Feedback</h2>

      <div v-for="fb in feedbacks" :key="fb.id" class="mb-4 p-4 bg-white rounded shadow">
        <div class="font-medium">{{ fb.name }} ({{ fb.event }})</div>
        <div class="text-sm text-gray-600 mb-2">{{ fb.date }}</div>
        <p>{{ fb.message }}</p>
      </div>
    </div>
  </OrganizerLayout>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import OrganizerLayout from '../../layouts/OrganizerLayout.vue';

const feedbacks = ref([
  {
    id: 1,
    name: 'Albert Reyes',
    event: 'Music Fest 2025',
    date: '2025-04-25',
    message: 'Great event! Looking forward to the next one.',
  },
  {
    id: 2,
    name: 'Dina Kurnia',
    event: 'Startup Fest',
    date: '2025-04-20',
    message: 'Well-organized and insightful sessions.',
  },
])
</script>
