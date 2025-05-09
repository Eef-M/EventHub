<template>
  <OrganizerLayout>
    <div class="p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-semibold">Manage Events</h2>
        <button class="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700" @click="addEvent">
          + Add Event
        </button>
      </div>

      <table class="w-full table-auto border-collapse">
        <thead>
          <tr class="bg-gray-100 text-left">
            <th class="px-4 py-2">Title</th>
            <th class="px-4 py-2">Date</th>
            <th class="px-4 py-2">Location</th>
            <th class="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="event in events" :key="event.id" class="border-b hover:bg-gray-50">
            <td class="px-4 py-2">{{ event.title }}</td>
            <td class="px-4 py-2">{{ event.date }}</td>
            <td class="px-4 py-2">{{ event.location }}</td>
            <td class="px-4 py-2 space-x-2">
              <button class="text-blue-600 hover:underline" @click="editEvent(event.id)">
                Edit
              </button>
              <button class="text-red-600 hover:underline" @click="deleteEvent(event.id)">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </OrganizerLayout>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import OrganizerLayout from '../../layouts/OrganizerLayout.vue'

const events = ref([
  {
    id: 1,
    title: 'Tech Expo 2025',
    date: '2025-06-10',
    location: 'Jakarta Convention Center',
  },
  {
    id: 2,
    title: 'Startup Fest',
    date: '2025-07-15',
    location: 'Bandung Tech Park',
  },
])

function addEvent() {
  // redirect to event creation form (optional route)
  alert('Redirect to Add Event Form')
}

function editEvent(id: number) {
  // redirect to edit page or open modal
  alert(`Edit event with ID: ${id}`)
}

function deleteEvent(id: number) {
  if (confirm('Are you sure you want to delete this event?')) {
    events.value = events.value.filter(event => event.id !== id)
  }
}
</script>
