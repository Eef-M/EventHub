<template>
  <aside class="w-64 bg-white shadow h-full">
    <div class="p-4 font-bold text-lg border-b">EventHub</div>
    <nav class="mt-4">
      <ul>
        <li class="px-4 py-2 hover:bg-gray-200">
          <RouterLink to="/admin">Dashboard</RouterLink>
        </li>
        <li class="px-4 py-2 hover:bg-gray-200">
          <RouterLink to="/admin/events">Manage Events</RouterLink>
        </li>
        <li class="px-4 py-2 hover:bg-gray-200">
          <RouterLink to="/admin/registrations">Registrations</RouterLink>
        </li>
        <li class="px-4 py-2 hover:bg-gray-200">
          <RouterLink to="/admin/feedbacks">Feedback</RouterLink>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';

</script>