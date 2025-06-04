<template>
  <Card>
    <CardContent class="py-2 px-4">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Title</TableHead>
            <TableHead>Image</TableHead>
            <TableHead>Description</TableHead>
            <TableHead>Location</TableHead>
            <TableHead>Date</TableHead>
            <TableHead>Time</TableHead>
            <TableHead class="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-if="!events.length">
            <TableCell :colspan="7" class="text-center text-muted-foreground">
              No events available
            </TableCell>
          </TableRow>
          <TableRow v-for="event in events" :key="event.id">
            <TableCell>{{ event.title }}</TableCell>
            <TableCell>
              <img class="w-24 h-24 object-cover rounded-sm" :src="event?.banner_url" alt="Event Thumbnail" />
            </TableCell>
            <TableCell>
              {{ event.description ? event.description.slice(0, 30) + '...' : 'No description available.' }}
            </TableCell>
            <TableCell>{{ event.location }}</TableCell>
            <TableCell>{{ formatDate(event.date) }}</TableCell>
            <TableCell>{{ formatTime(event.time) }}</TableCell>
            <TableCell class="text-right space-x-2">
              <Button size="icon" class="bg-slate-600 hover:bg-slate-700 cursor-pointer" @click="$emit('edit', event)">
                <Pencil class="w-4 h-4" />
              </Button>
              <Button size="icon" class="bg-red-600 hover:bg-red-700 cursor-pointer" @click="$emit('delete', event)">
                <Trash class="w-4 h-4" />
              </Button>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </CardContent>
  </Card>
</template>

<script setup lang="ts">
import { Card, CardContent } from '@/components/ui/card'
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Pencil, Trash } from 'lucide-vue-next'
import type { EventInterface } from '@/types/event'
import { formatDate, formatTime } from '@/utils/format'

interface Props {
  events: EventInterface[]
}

defineProps<Props>()

defineEmits<{
  edit: [event: EventInterface]
  delete: [event: EventInterface]
}>()
</script>