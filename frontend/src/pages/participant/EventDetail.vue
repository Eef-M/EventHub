<template>
  <ParticipantLayout>
    <section class="py-8 max-w-7xl mx-auto">
      <div class="rounded-2xl overflow-hidden shadow-md mb-8">
        <img :src="eventStore?.event?.banner_url" alt="Event Banner" class="w-full h-64 object-cover" />
      </div>

      <div v-if="eventStore.event" class="space-y-4 mb-10">
        <h1 class="text-3xl font-bold text-gray-900">{{ eventStore?.event?.title }}</h1>
        <p class="text-gray-600">Category: <span class="font-medium">{{ eventStore?.event?.category }}</span></p>
        <p class="text-gray-600">Location: <span class="font-medium">{{ eventStore?.event?.location }}</span></p>
        <p class="text-gray-600">Date: <span class="font-medium">{{ formatDate(eventStore?.event?.date) }}</span></p>
        <p class="text-gray-600">Time: <span class="font-medium">{{ formatTime(eventStore?.event?.time) }}</span></p>
      </div>

      <div class="mb-10">
        <h2 class="text-xl font-semibold mb-2">Description</h2>
        <p class="text-gray-700">
          {{ eventStore?.event?.description || 'No description available.' }}
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
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ParticipantLayout from '../../layouts/ParticipantLayout.vue'
import { useEventStore } from '@/stores/eventStore'

const route = useRoute()
const eventId = route.params.id as string
const eventStore = useEventStore()

onMounted(async () => {
  eventStore.getEventById(eventId)
})

function formatDate(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

function formatTime(timeStr: string): string {
  const time = new Date(timeStr)
  if (time.getFullYear() === 1) return 'N/A'
  return time.toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  })
}
</script>
