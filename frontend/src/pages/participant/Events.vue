<template>
  <ParticipantLayout>
    <section class="py-10">
      <div class="max-w-7xl mx-auto px-4">
        <div class="flex flex-col md:flex-row justify-between items-center gap-4 mb-8">
          <input v-model="search" type="text" placeholder="Search by title or location..."
            class="w-full md:w-1/3 px-4 py-2 border rounded-lg" />

          <select v-model="category" class="px-4 py-2 border rounded-lg">
            <option value="">All Categories</option>
            <option value="Webinar">Webinar</option>
            <option value="Workshop">Workshop</option>
            <option value="Concert">Concert</option>
          </select>

          <select v-model="sortOrder" class="px-4 py-2 border rounded-lg">
            <option value="asc">Date: Oldest First</option>
            <option value="desc">Date: Newest First</option>
          </select>
        </div>

        <div class="grid md:grid-cols-3 gap-6">
          <RouterLink v-for="event in filteredEvents" :key="event.id" :to="`/events/${event.id}/detail`">
            <EventCard :event="event" />
          </RouterLink>
        </div>
      </div>
    </section>
  </ParticipantLayout>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useEventStore } from '../../stores/eventStore';
import { RouterLink } from 'vue-router';
import EventCard from '../../components/participant/EventCard.vue';
import ParticipantLayout from '../../layouts/ParticipantLayout.vue';

const eventStore = useEventStore()

const search = ref('');
const category = ref('');
const sortOrder = ref('asc');

onMounted(() => {
  eventStore.getEvents()
});

const filteredEvents = computed(() => {
  let events = eventStore.eventsState.data;

  if (search.value) {
    const keyword = search.value.toLowerCase();
    events = events.filter(
      e =>
        e.title.toLowerCase().includes(keyword) ||
        e.location.toLowerCase().includes(keyword)
    );
  }

  if (category.value) {
    events = events.filter(e => e.category === category.value);
  }

  events = events.slice().sort((a, b) => {
    if (sortOrder.value === 'asc') {
      return new Date(a.date).getTime() - new Date(b.date).getTime();
    } else {
      return new Date(b.date).getTime() - new Date(a.date).getTime();
    }
  });

  return events;
});
</script>
