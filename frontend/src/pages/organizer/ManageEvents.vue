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
      <DialogContent class="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>Add Event</DialogTitle>
        </DialogHeader>

        <div class="grid gap-4 py-4">
          <div class="grid gap-2">
            <Label for="title">Title</Label>
            <Input id="title" v-model="form.title" placeholder="Event title" />
          </div>

          <div class="grid gap-2">
            <Label for="description">Description</Label>
            <Textarea id="description" v-model="form.description" placeholder="Event description" />
          </div>

          <div class="grid gap-2">
            <Label for="location">Location</Label>
            <Input id="location" v-model="form.location" placeholder="Location" />
          </div>

          <div class="grid gap-2">
            <Label for="category">Category</Label>
            <Input id="category" v-model="form.category" placeholder="Category" />
          </div>

          <div class="grid gap-2 md:grid-cols-2 md:grid">
            <div class="grid gap-2">
              <Label for="date">Date</Label>
              <Input id="date" type="date" v-model="form.date" />
            </div>

            <div class="grid gap-2">
              <Label for="time">Time</Label>
              <Input id="time" type="time" v-model="form.time" />
            </div>
          </div>

          <div class="grid gap-2">
            <Label for="banner">Image</Label>
            <Input id="banner" type="file" accept="image/*"
              @change="(e: any) => form.banner_image = e.target.files[0]" />
            <div v-if="form.banner_image" class="mt-2">
              <p class="text-sm text-muted-foreground mb-1">New preview:</p>
              <img v-if="imagePreviewUrl" :src="imagePreviewUrl" alt="New Banner Preview"
                class="rounded-md border max-w-24 object-cover max-h-24" />
            </div>
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" @click="showAddModal = false">Cancel</Button>
          <Button :disabled="eventStore.createState.loading" @click="handleCreateEvent">Add Event</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <!-- Edit Event Modal -->
    <Dialog v-model:open="showEditModal">
      <DialogContent class="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>Edit Event</DialogTitle>
        </DialogHeader>

        <div class="grid gap-4 py-4">
          <div class="grid gap-2">
            <Label for="title">Title</Label>
            <Input id="title" v-model="form.title" placeholder="Event title" />
          </div>

          <div class="grid gap-2">
            <Label for="description">Description</Label>
            <Textarea id="description" v-model="form.description" placeholder="Event description" />
          </div>

          <div class="grid gap-2">
            <Label for="location">Location</Label>
            <Input id="location" v-model="form.location" placeholder="Location" />
          </div>

          <div class="grid gap-2">
            <Label for="category">Category</Label>
            <Input id="category" v-model="form.category" placeholder="Category" />
          </div>

          <div class="grid gap-2 md:grid-cols-2 md:grid">
            <div class="grid gap-2">
              <Label for="date">Date</Label>
              <Input id="date" type="date" v-model="form.date" />
            </div>

            <div class="grid gap-2">
              <Label for="time">Time</Label>
              <Input id="time" type="time" v-model="form.time" />
            </div>
          </div>

          <div class="grid gap-2">
            <Label for="banner">Image</Label>
            <Input id="banner" type="file" accept="image/*" @change="handleBannerChange" />
            <div class="flex gap-6 mt-2 items-start">
              <div v-if="selectedEvent?.banner_url" class="flex flex-col items-center gap-1">
                <p class="text-sm text-muted-foreground">Current image:</p>
                <img :src="selectedEvent.banner_url" alt="Current Banner"
                  class="rounded-md border w-24 h-24 object-cover" />
              </div>
              <div v-if="form.banner_image" class="flex flex-col items-center gap-1">
                <p class="text-sm text-muted-foreground">New preview:</p>
                <img v-if="imagePreviewUrl" :src="imagePreviewUrl" alt="New Banner Preview"
                  class="rounded-md border w-24 h-24 object-cover" />
              </div>
            </div>
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" @click="showEditModal = false">Cancel</Button>
          <Button :disabled="eventStore.updateState.loading" @click="handleUpdateEvent">Update</Button>
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
import { computed, onMounted, reactive, ref } from 'vue'
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
import { useEventStore } from '@/stores/eventStore'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { toast } from 'vue-sonner'

const organizerStore = useOrganizerStore()
const eventStore = useEventStore()

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
  const isoDate = new Date(event.date)
  const formattedDate = isoDate.toISOString().split('T')[0]

  form.title = event.title
  form.description = event.description
  form.location = event.location
  form.category = event.category
  form.date = formattedDate
  form.time = formatTime(event.time)
  form.banner_image = null
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

const form = reactive({
  title: '',
  description: '',
  location: '',
  category: '',
  date: '',
  time: '',
  banner_image: null as File | null,
})

function resetForm() {
  form.title = ''
  form.description = ''
  form.location = ''
  form.category = ''
  form.date = ''
  form.time = ''
  form.banner_image = null
}

async function handleCreateEvent() {
  try {
    const payload = new FormData()
    payload.append('title', form.title)
    payload.append('description', form.description)
    payload.append('location', form.location)
    payload.append('category', form.category)
    payload.append('date', form.date)
    payload.append('time', form.time)
    if (form.banner_image) {
      payload.append('banner_image', form.banner_image)
    }

    await eventStore.createEvent(payload)

    await organizerStore.getMyEvents()

    toast.success('Event created successfully', {
      description: new Date().toLocaleString(),
    })
    showAddModal.value = false
    resetForm()
  } catch (err) {
    toast.error('Failed to create event', {
      description: eventStore.createState.error || 'An unexpected error occurred.',
    })
    console.error('Create Event Failed:', eventStore.createState.error)
  }
}

function handleBannerChange(e: Event & { target: HTMLInputElement }) {
  form.banner_image = e.target.files?.[0] || null
}

const imagePreviewUrl = computed(() =>
  form.banner_image ? URL.createObjectURL(form.banner_image) : null
)


async function handleUpdateEvent() {
  if (!selectedEvent.value) return

  try {
    const payload = new FormData()
    payload.append('title', form.title)
    payload.append('description', form.description)
    payload.append('location', form.location)
    payload.append('category', form.category)
    payload.append('date', form.date)
    payload.append('time', form.time)
    if (form.banner_image) {
      payload.append('banner_image', form.banner_image)
    }

    await eventStore.updateEvent(selectedEvent.value.id, payload)

    await organizerStore.getMyEvents()

    toast.success('Event updated successfully', {
      description: new Date().toLocaleString(),
    })
    showEditModal.value = false
  } catch (err) {
    toast.error('Failed to create event', {
      description: eventStore.updateState.error || 'An unexpected error occurred.',
    })
    console.error('Update Event Failed:', eventStore.updateState.error)
  }
}
</script>
