<template>
  <OrganizerLayout>
    <div class="flex items-center justify-between mb-4">
      <h1 class="text-2xl font-semibold">Manage Events</h1>
      <Button class="bg-purple-600 hover:bg-purple-700 cursor-pointer" @click="openCreateModal">
        <Plus class="w-4 h-4 mr-2" /> Create Event
      </Button>
    </div>

    <!-- Events data table -->
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
            <TableRow v-if="!organizerStore.eventsState.data.length">
              <TableCell :colspan="7" class="text-center text-muted-foreground">
                No events available
              </TableCell>
            </TableRow>
            <TableRow v-for="event in organizerStore.eventsState.data" :key="event.id">
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
                <Button size="icon" class="bg-slate-600 hover:bg-slate-700 cursor-pointer"
                  @click="openEditModal(event)">
                  <Pencil class="w-4 h-4" />
                </Button>
                <Button size="icon" class="bg-red-600 hover:bg-red-700 cursor-pointer" @click="openDeleteModal(event)">
                  <Trash class="w-4 h-4" />
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>

    <!-- Modal: Create/Edit -->
    <Dialog :open="showFormModal" @update:open="showFormModal = $event">
      <DialogContent class="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>{{ isEditing ? "Edit Event" : "Create Event" }}</DialogTitle>
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
            <div v-if="isEditing">
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
            <div v-else>
              <Input id="banner" type="file" accept="image/*"
                @change="(e: any) => form.banner_image = e.target.files[0]" />
              <div v-if="form.banner_image" class="mt-2">
                <p class="text-sm text-muted-foreground mb-1">New preview:</p>
                <img v-if="imagePreviewUrl" :src="imagePreviewUrl" alt="New Banner Preview"
                  class="rounded-md border max-w-24 object-cover max-h-24" />
              </div>
            </div>
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" @click="showFormModal = false" class="cursor-pointer">Cancel</Button>
          <Button :disabled="isEditing ? eventStore.updateState.loading : eventStore.createState.loading"
            @click="handleSubmitEvent"
            :class="isEditing ? 'bg-slate-600 hover:bg-slate-700 cursor-pointer' : 'bg-purple-600 hover:bg-purple-700 cursor-pointer'">
            {{ isEditing ? "Update" : "Create" }}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <!-- Modal: Confirm Delete -->
    <Dialog :open="showDeleteModal" @update:open="showDeleteModal = $event">
      <DialogContent class="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Delete Ticket</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete <strong>{{ selectedEvent?.title }}</strong>? This action cannot be undone.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button class="cursor-pointer" variant="outline" @click="showDeleteModal = false">Cancel</Button>
          <Button class="bg-red-600 hover:bg-red-700 cursor-pointer" @click="confirmDeleteEvent">Delete</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
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
  DialogFooter,
  DialogDescription
} from '@/components/ui/dialog'
import type { Event } from '@/types/event'
import { useEventStore } from '@/stores/eventStore'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { toast } from 'vue-sonner'
import { Label } from '@/components/ui/label'

const organizerStore = useOrganizerStore()
const eventStore = useEventStore()

const showFormModal = ref(false)
const showDeleteModal = ref(false)
const isEditing = ref(false)
const selectedEvent = ref<Event | null>(null)

onMounted(() => {
  organizerStore.getMyEvents()
})

const defaultForm = {
  title: '',
  description: '',
  location: '',
  category: '',
  date: '',
  time: '',
  banner_image: null as File | null,
}

const form = reactive({ ...defaultForm })

function openCreateModal() {
  isEditing.value = false;
  Object.assign(form, defaultForm);
  showFormModal.value = true;
}

function openEditModal(event: any) {
  isEditing.value = true;
  const isoDate = new Date(event.date)
  const formattedDate = isoDate.toISOString().split('T')[0]
  Object.assign(form, {
    title: event.title,
    description: event.description,
    location: event.location,
    category: event.category,
    date: formattedDate,
    time: formatTime(event.time),
    banner_image: null
  });
  selectedEvent.value = event;
  showFormModal.value = true;
}

function openDeleteModal(event: any) {
  selectedEvent.value = event;
  showDeleteModal.value = true;
}

async function handleSubmitEvent() {
  try {
    const formData = new FormData();
    formData.append('title', form.title)
    formData.append('description', form.description)
    formData.append('location', form.location)
    formData.append('category', form.category)
    formData.append('date', form.date)
    formData.append('time', form.time)
    if (form.banner_image) {
      formData.append('banner_image', form.banner_image)
    }

    if (isEditing.value && selectedEvent.value) {
      await eventStore.updateEvent(selectedEvent.value.id, formData)
    } else {
      await eventStore.createEvent(formData)
    }

    await organizerStore.getMyEvents()

    toast.success(`Event ${isEditing.value ? 'updated' : 'created'} successfuly`, {
      description: new Date().toLocaleString(),
    })
    showFormModal.value = false;
    Object.assign(form, defaultForm);
  } catch (err) {
    toast.error(`Failed to ${isEditing.value ? 'update' : 'create'} event`, {
      description: `An unexpected error occurred. ${isEditing.value ? eventStore.updateState.error : eventStore.createState.error}`,
    })
    console.error(`Failed to ${isEditing.value ? 'update' : 'create'} event`, err);
  }
}

async function confirmDeleteEvent() {
  try {
    if (!selectedEvent.value) return;

    await eventStore.deleteEvent(selectedEvent.value.id);

    await organizerStore.getMyEvents();

    toast.success('Event deleted successfully', {
      description: new Date().toLocaleString(),
    })
    showDeleteModal.value = false;
    selectedEvent.value = null;
  } catch (err) {
    toast.error('Failed to delete event', {
      description: eventStore.deleteState.error || `An unexpected error occurred. ${err}`,
    })
    console.error('Delete Event Failed: ', eventStore.deleteState.error, err)
  }
}

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

function handleBannerChange(e: Event & { target: HTMLInputElement }) {
  form.banner_image = e.target.files?.[0] || null
}

const imagePreviewUrl = computed(() =>
  form.banner_image ? URL.createObjectURL(form.banner_image) : null
)
</script>
