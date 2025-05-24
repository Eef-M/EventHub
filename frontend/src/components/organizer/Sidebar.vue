<template>
  <aside class="hidden md:block w-64 bg-card border-r p-4">
    <nav class="flex flex-col gap-2">
      <RouterLink v-for="item in menu" :key="item.label" :to="item.href"
        class="flex items-center gap-2 text-lg rounded-md px-3 py-2 hover:bg-slate-200 transition"
        active-class="bg-slate-200 text-purple-600">
        <component :is="item.icon" class="w-4 h-4" />
        {{ item.label }}
      </RouterLink>
    </nav>
  </aside>
</template>

<script setup>
import { Home, Calendar, Users, MessageSquare, Ticket } from 'lucide-vue-next'

const menu = [
  { label: 'Dashboard', href: '/organizer/dashboard', icon: Home },
  { label: 'Manage Events', href: '/organizer/manage-events', icon: Calendar },
  { label: 'Manage Tickets', href: '/organizer/manage-tickets', icon: Ticket },
  { label: 'Registrations', href: '/organizer/registrations', icon: Users },
  { label: 'Feedback', href: '/organizer/feedback', icon: MessageSquare },
]
</script>
