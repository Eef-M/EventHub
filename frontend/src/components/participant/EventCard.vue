<template>
  <div class="bg-white rounded-2xl shadow p-5 hover:shadow-xl transition">
    <img class="w-full h-48 object-cover rounded-lg mb-4" :src="event.banner_url || defaultImage"
      alt="Event Thumbnail" />

    <h2 class="text-xl font-semibold text-gray-800 mb-2">
      {{ event.title }}
    </h2>

    <p class="text-sm text-gray-600 mb-1">
      📍 {{ event.location }}
    </p>

    <p class="text-sm text-gray-600 mb-3">
      📅 {{ formatDate(event.date) }}
    </p>

    <p class="text-gray-700 text-sm">
      {{ event.description ? event.description.slice(0, 40) + '...' : 'No description available.' }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { formatDate } from '@/utils/format';
import type { EventInterface } from '../../types/event';

defineProps<{
  event: EventInterface
}>()

const defaultImage = 'https://images.pexels.com/photos/50675/banquet-wedding-society-deco-50675.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
</script>
