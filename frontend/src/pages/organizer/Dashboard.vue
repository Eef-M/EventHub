<template>
  <OrganizerLayout>
    <h2 class="text-2xl font-semibold mb-4">Dashboard</h2>
    <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-5">
      <Card v-for="stat in stats" :key="stat.label" class="shadow">
        <CardHeader>
          <CardTitle class="flex items-center justify-between">
            {{ stat.label }}
            <component :is="stat.icon" class="w-5 h-5 text-muted-foreground" />
          </CardTitle>
        </CardHeader>
        <CardContent>
          <p class="text-2xl font-bold">{{ stat.value }}</p>
        </CardContent>
      </Card>
    </div>

    <div class="mt-8">
      <h2 class="text-lg font-semibold mb-4">Recent Registrations</h2>
      <Card>
        <CardContent class="p-2">
          <Table class="text-md">
            <TableHeader>
              <TableRow>
                <TableHead>Username</TableHead>
                <TableHead>Event Title</TableHead>
                <TableHead>Ticket Name</TableHead>
                <TableHead>Status</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-if="!organizerStore.stats?.recent_registrations?.length">
                <TableCell :colspan="3" class="text-center text-muted-foreground pt-5">No recent registrations found
                </TableCell>
              </TableRow>
              <TableRow v-for="reg in organizerStore.stats?.recent_registrations" :key="reg.id">
                <TableCell>{{ reg.username }}</TableCell>
                <TableCell>{{ reg.event_title }}</TableCell>
                <TableCell>{{ reg.ticket_name }}</TableCell>
                <TableCell>
                  <Badge class="text-sm font-bold" :class="reg.status === 'registered' ? 'bg-green-600' : 'bg-red-600'">
                    {{ reg.status }}
                  </Badge>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  </OrganizerLayout>
</template>

<script setup>
import { onMounted } from 'vue'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import { Calendar, Users, MessageSquare, CalendarCheck, CalendarX } from 'lucide-vue-next'
import OrganizerLayout from '@/layouts/OrganizerLayout.vue'
import { useOrganizerStore } from '@/stores/organizerStore'

const organizerStore = useOrganizerStore()

onMounted(() => {
  organizerStore.getDashboardStats()
})

const stats = [
  { label: 'Total Events', value: organizerStore.stats?.total_events, icon: Calendar },
  { label: 'Registrations', value: organizerStore.stats?.total_registrations, icon: Users },
  { label: 'Registered', value: organizerStore.stats?.registered, icon: CalendarCheck },
  { label: 'Cancelled', value: organizerStore.stats?.cancelled_registrations, icon: CalendarX },
  { label: 'Feedback Received', value: organizerStore.stats?.feedback_received, icon: MessageSquare },
]
</script>
