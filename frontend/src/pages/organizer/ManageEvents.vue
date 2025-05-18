<template>
  <OrganizerLayout>
    <div class="flex items-center justify-between mb-4">
      <h1 class="text-2xl font-semibold">Manage Events</h1>
      <Button variant="default" @click="showAddModal = true">
        <Plus class="w-4 h-4 mr-2" /> Add Event
      </Button>
    </div>

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
            <TableRow v-if="!organizerStore.events.length">
              <TableCell :colspan="4" class="text-center text-muted-foreground">
                No events found
              </TableCell>
            </TableRow>
            <TableRow v-for="event in organizerStore.events" :key="event.id">
              <TableCell>{{ event.title }}</TableCell>
              <TableCell>
                <img class="w-24 h-24 object-cover rounded-sm" :src="event?.banner_url" alt="Event Thumbnail" />
              </TableCell>
              <TableCell>{{ event.description ? event.description.slice(0, 30) + '...' : 'No description available.' }}
              </TableCell>
              <TableCell>{{ event.location }}</TableCell>
              <TableCell>{{ formatDate(event.date) }}</TableCell>
              <TableCell>{{ formatTime(event.time) }}</TableCell>
              <TableCell class="text-right space-x-2">
                <Button size="icon" variant="outline" @click="openEditModal(event)">
                  <Pencil class="w-4 h-4" />
                </Button>
                <Button size="icon" variant="destructive" @click="openDeleteModal(event.id)">
                  <Trash class="w-4 h-4" />
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>

    <!-- Add Event Modal -->
    <Dialog v-model:open="showAddModal">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add Event</DialogTitle>
        </DialogHeader>
        <p>Form Add Event here</p>
        <DialogFooter>
          <Button variant="outline" @click="showAddModal = false">Cancel</Button>
          <Button>Add</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <!-- Edit Event Modal -->
    <Dialog v-model:open="showEditModal">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Event</DialogTitle>
        </DialogHeader>
        <p>Form Edit Event: {{ selectedEvent?.title }}</p>
        <DialogFooter>
          <Button variant="outline" @click="showEditModal = false">Cancel</Button>
          <Button>Update</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <!-- Delete Confirm Modal -->
    <AlertDialog v-model:open="showDeleteModal">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Confirm Delete</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete this event? This action cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel @click="showDeleteModal = false">Cancel</AlertDialogCancel>
          <AlertDialogAction @click="confirmDeleteEvent">Delete</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  </OrganizerLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Card, CardContent } from '@/components/ui/card'
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Pencil, Trash, Plus } from 'lucide-vue-next'
import OrganizerLayout from '@/layouts/OrganizerLayout.vue'
import { useOrganizerStore } from '@/stores/organizerStore'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from '@/components/ui/dialog'
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogCancel,
  AlertDialogAction
} from '@/components/ui/alert-dialog'
import type { Event } from '@/types/event'

const organizerStore = useOrganizerStore()

const showAddModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const selectedEvent = ref<Event | null>(null)
const eventToDelete = ref<string | null>(null)

onMounted(() => {
  organizerStore.getMyEvents()
})

function formatDate(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

function formatTime(timeStr: string): string {
  const time = new Date(timeStr)
  if (time.getFullYear() === 1) return 'N/A'
  return time.toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  })
}

function openEditModal(event: Event) {
  selectedEvent.value = event
  showEditModal.value = true
}

function openDeleteModal(id: string) {
  eventToDelete.value = id
  showDeleteModal.value = true
}

function confirmDeleteEvent(): void {
  console.log('Delete confirmed for event ID:', eventToDelete.value)
  showDeleteModal.value = false
}
</script>
