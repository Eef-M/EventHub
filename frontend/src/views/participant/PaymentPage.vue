<template>
  <div class="max-w-2xl mx-auto mt-10 space-y-6">
    <!-- Event Info -->
    <div v-if="event" class="p-6 border rounded-lg shadow-sm bg-white">
      <h1 class="text-2xl font-bold mb-2">{{ event.title }}</h1>
      <p class="text-gray-600 mb-4">{{ event.description }}</p>
      <div class="flex justify-between text-sm text-gray-500">
        <span>Tanggal: {{ formatDate(event.date) }}</span>
        <!-- <span>Harga: {{ formatPrice(event.price) }}</span> Need update backend first -->
      </div>
    </div>
    <div v-else class="text-center text-gray-500">Memuat info event...</div>

    <!-- Stripe Checkout -->
    <StripeCheckout v-if="event" :eventId="event.id" />
  </div>
</template>

<script setup lang="ts">
import { useEventStore } from '@/stores/eventStore';
import { formatDate } from '@/utils/format';
import { computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const eventStore = useEventStore()

const route = useRoute()
const eventId = route.params.eventId as string

const event = computed(() => eventStore?.singleEventState?.data)

onMounted(async () => {
  eventStore.getEventById(eventId)
})
</script>