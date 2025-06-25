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
            <TableHead class="text-center">Availability</TableHead>
            <TableHead class="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-if="!events.length">
            <TableCell :colspan="8" class="text-center text-muted-foreground">
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
            <TableCell class="text-center">
              <div class="flex flex-col space-y-2">
                <!-- Public Status Toggle -->
                <div class="flex items-center justify-center space-x-2">
                  <label class="text-xs font-medium text-slate-600">Public:</label>
                  <Button size="sm" :class="[
                    'px-3 py-1 text-xs cursor-pointer',
                    event.is_public
                      ? 'bg-green-600 hover:bg-green-700'
                      : 'bg-black hover:bg-slate-800'
                  ]" @click="$emit('togglePublic', event)" :disabled="isUpdatingAvailability">
                    {{ event.is_public ? 'Yes' : 'No' }}
                  </Button>
                </div>

                <!-- Open Status Toggle -->
                <div class="flex items-center justify-center space-x-2">
                  <label class="text-xs font-medium text-gray-600">Open:</label>
                  <Button size="sm" :class="[
                    'px-3 py-1 text-xs cursor-pointer',
                    event.is_open
                      ? 'bg-green-600 hover:bg-green-700'
                      : 'bg-black hover:bg-slate-800'
                  ]" @click="$emit('toggleOpen', event)" :disabled="isUpdatingAvailability">
                    {{ event.is_open ? 'Yes' : 'No' }}
                  </Button>
                </div>
              </div>
            </TableCell>
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
  isUpdatingAvailability?: boolean
}

withDefaults(defineProps<Props>(), {
  isUpdatingAvailability: false
})

defineEmits<{
  edit: [event: EventInterface]
  delete: [event: EventInterface]
  togglePublic: [event: EventInterface]
  toggleOpen: [event: EventInterface]
}>()
</script>