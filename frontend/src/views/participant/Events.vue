<template>
  <ParticipantLayout>
    <section class="py-12 bg-gradient-to-br from-slate-50 to-slate-100 min-h-screen">
      <div class="max-w-7xl mx-auto px-4">
        <div class="text-center mb-12">
          <h1 class="text-4xl font-bold text-slate-900 mb-4">
            Discover Amazing Events
          </h1>
          <p class="text-xl text-slate-600 max-w-2xl mx-auto">
            Find and join exciting events happening around you. From workshops to concerts, discover your next
            adventure.
          </p>
        </div>

        <div class="mb-8 bg-white/80 backdrop-blur-sm border-0 shadow-lg rounded-lg">
          <div class="p-6">
            <div class="flex flex-col lg:flex-row gap-4 items-center">
              <div class="relative flex-1 w-full">
                <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 text-slate-400 h-4 w-4" fill="none"
                  stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input v-model="search" type="text" placeholder="Search by title or location..."
                  class="pl-10 h-12 w-full bg-white border-slate-200 rounded-lg focus:border-purple-500 focus:ring-purple-500 focus:ring-1 outline-none" />
              </div>

              <div class="flex flex-col sm:flex-row gap-4 w-full lg:w-auto">
                <div class="relative">
                  <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z" />
                  </svg>
                  <select v-model="category"
                    class="pl-10 pr-4 py-3 w-full sm:w-48 h-12 bg-white border border-slate-200 rounded-lg focus:border-purple-500 focus:ring-purple-500 focus:ring-1 outline-none appearance-none">
                    <option value="">All Categories</option>
                    <option value="Webinar">Webinar</option>
                    <option value="Workshop">Workshop</option>
                    <option value="Concert">Concert</option>
                  </select>
                </div>

                <div class="relative">
                  <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 text-slate-400 h-4 w-4" fill="none"
                    stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M3 4h13M3 8h9m-9 4h9m5-4v12m0 0l-4-4m4 4l4-4" />
                  </svg>
                  <select v-model="sortOrder"
                    class="pl-10 pr-4 py-3 w-full sm:w-48 h-12 bg-white border border-slate-200 rounded-lg focus:border-purple-500 focus:ring-purple-500 focus:ring-1 outline-none appearance-none">
                    <option value="asc">Date: Oldest First</option>
                    <option value="desc">Date: Newest First</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="filteredEvents.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <RouterLink v-for="event in filteredEvents" :key="event.id" :to="`/events/${event.id}/detail`"
            class="group transform transition-all duration-300 hover:scale-105">
            <EventCard :event="event" />
          </RouterLink>
        </div>

        <div v-else class="text-center py-16">
          <div class="bg-white rounded-2xl shadow-lg p-12 max-w-md mx-auto">
            <svg class="h-16 w-16 text-slate-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            <h3 class="text-xl font-semibold text-slate-900 mb-2">No Events Found</h3>
            <p class="text-slate-600">Try adjusting your search criteria or check back later for new events.</p>
          </div>
        </div>
      </div>
    </section>
  </ParticipantLayout>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import ParticipantLayout from '@/layouts/ParticipantLayout.vue';
import { useEventStore } from '@/stores/eventStore';
import EventCard from '@/components/participant/EventCard.vue';
import { RouterLink } from 'vue-router';

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