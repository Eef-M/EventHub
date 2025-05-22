<template>
  <OrganizerLayout>
    <div class="p-6">
      <h2 class="text-2xl font-semibold mb-4">Event Registrations</h2>
      <Card>
        <CardContent class="py-2 px-4">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Username</TableHead>
                <TableHead>Email</TableHead>
                <TableHead>Event</TableHead>
                <TableHead>Ticket</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Registered at</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-if="!organizerStore.registrationsState.data.length">
                <TableCell :colspan="6" class="text-center text-muted-foreground">
                  No events found
                </TableCell>
              </TableRow>
              <TableRow v-for="reg in organizerStore.registrationsState.data" :key="reg.id">
                <TableCell>{{ reg.username }}</TableCell>
                <TableCell>{{ reg.email }}</TableCell>
                <TableCell>{{ reg.event_title }}</TableCell>
                <TableCell>{{ reg.ticket_name }}</TableCell>
                <TableCell>
                  <Badge class="text-sm font-bold" :class="reg.status === 'registered' ? 'bg-green-600' : 'bg-red-600'">
                    {{ reg.status }}
                  </Badge>
                </TableCell>
                <TableCell>{{ reg.registered_at }}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  </OrganizerLayout>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import OrganizerLayout from '@/layouts/OrganizerLayout.vue';
import { useOrganizerStore } from '@/stores/organizerStore';
import { Card, CardContent } from '@/components/ui/card'
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table'
import { Badge } from '@/components/ui/badge';

const organizerStore = useOrganizerStore()

onMounted(() => {
  organizerStore.getEventRegistrations()
})
</script>
