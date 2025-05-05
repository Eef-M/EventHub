<template>
  <AdminLayout>
    <div class="p-6">
      <h2 class="text-2xl font-semibold mb-4">Event Registrations</h2>

      <table class="w-full table-auto border-collapse">
        <thead>
          <tr class="bg-gray-100 text-left">
            <th class="px-4 py-2">Name</th>
            <th class="px-4 py-2">Email</th>
            <th class="px-4 py-2">Event</th>
            <th class="px-4 py-2">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="reg in registrations" :key="reg.id" class="border-b hover:bg-gray-50">
            <td class="px-4 py-2">{{ reg.name }}</td>
            <td class="px-4 py-2">{{ reg.email }}</td>
            <td class="px-4 py-2">{{ reg.event }}</td>
            <td class="px-4 py-2">{{ reg.status }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </AdminLayout>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import AdminLayout from '../../layouts/AdminLayout.vue';

const registrations = ref([
  {
    id: 1,
    name: 'John Doe',
    email: 'john@example.com',
    event: 'Tech Expo 2025',
    status: 'Approved',
  },
  {
    id: 2,
    name: 'Jane Smith',
    email: 'jane@example.com',
    event: 'Startup Fest',
    status: 'Pending',
  },
])
</script>
