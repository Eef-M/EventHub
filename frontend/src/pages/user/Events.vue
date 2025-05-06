<template>
  <UserLayout>
    <section class="py-10">
      <div class="max-w-7xl mx-auto px-4">
        <!-- <h2 class="text-3xl font-semibold text-center mb-10">Events</h2> -->
        <div class="flex flex-col md:flex-row justify-between items-center gap-4 mb-8">
          <input type="text" v-model="searchQuery" placeholder="Search by title or location..."
            class="w-full md:w-1/3 px-4 py-2 border rounded-lg" />

          <select v-model="selectedCategory" class="px-4 py-2 border rounded-lg">
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

        <div v-if="filteredEvents.length === 0" class="text-center mt-10 text-gray-500">
          No events found.
        </div>
      </div>
    </section>
  </UserLayout>
</template>


<script lang="ts" setup>
import { RouterLink } from 'vue-router';
import UserLayout from '../../layouts/UserLayout.vue';
import EventCard from '../../components/user/EventCard.vue';
import { computed, ref } from 'vue';
import { dummyEvents } from '../../data/events';

const searchQuery = ref('')
const selectedCategory = ref('')
const sortOrder = ref<'asc' | 'desc'>('asc')

const filteredEvents = computed(() => {
  let filtered = [...dummyEvents]

  const query = searchQuery.value.toLowerCase()
  if (query) {
    filtered = filtered.filter(
      e =>
        e.title.toLowerCase().includes(query) ||
        e.location.toLowerCase().includes(query)
    )
  }

  if (selectedCategory.value) {
    filtered = filtered.filter(e => e.category === selectedCategory.value)
  }

  filtered.sort((a, b) => {
    const dateA = new Date(a.date).getTime()
    const dateB = new Date(b.date).getTime()
    return sortOrder.value === 'asc' ? dateA - dateB : dateB - dateA
  })

  return filtered
})
</script>