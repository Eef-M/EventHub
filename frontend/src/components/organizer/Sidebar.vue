<template>
  <aside class="w-64 bg-white shadow h-full">
    <div class="p-4 font-bold text-lg border-b">
      <RouterLink to="/organizer/dashboard">
        <span class="text-purple-600">Event</span>Hub
      </RouterLink>
    </div>
    <nav class="mt-4">
      <ul>
        <li class="px-4 py-4 hover:shadow-md hover:text-purple-800">
          <RouterLink to="/organizer/dashboard">
            <v-icon name="md-dashboard" scale="1.3" />
            Dashboard
          </RouterLink>
        </li>
        <li class="px-4 py-4 hover:shadow-md hover:text-purple-800">
          <RouterLink to="/organizer/manage-events">
            <v-icon name="md-displaysettings-sharp" scale="1.3" />
            Manage Events
          </RouterLink>
        </li>
        <li class="px-4 py-4 hover:shadow-md hover:text-purple-800">
          <RouterLink to="/organizer/registrations">
            <v-icon name="fa-clipboard-list" scale="1.3" />
            Registrations
          </RouterLink>
        </li>
        <li class="px-4 py-4 hover:shadow-md hover:text-purple-800">
          <RouterLink to="/organizer/feedback">
            <v-icon name="la-comment" scale="1.3" />
            Feedback
          </RouterLink>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';

</script>