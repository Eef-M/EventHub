<template>
  <OrganizerLayout>
    <div class="flex items-center justify-between mb-4">
      <h1 class="text-2xl font-semibold">Manage Events</h1>
      <Button class="bg-purple-600 hover:bg-purple-700 cursor-pointer" @click="openCreateModal">
        <Plus class="w-4 h-4 mr-2" /> Create Event
      </Button>
    </div>

    <EventsTable :events="organizerStore.eventsState.data" :isUpdatingAvailability="eventStore.eventAvl.loading"
      @edit="openEditModal" @delete="openDeleteModal" @togglePublic="handleTogglePublic"
      @toggleOpen="handleToggleOpen" />

    <EventFormModal v-model:isOpen="showFormModal" :isEditing="isEditing" :eventData="selectedEvent"
      :isLoading="isEditing ? eventStore.updateState.loading : eventStore.createState.loading"
      @submit="handleSubmitEvent" @cancel="handleCancelForm" />

    <DeleteConfirmModal v-model:isOpen="showDeleteModal" :eventTitle="selectedEvent?.title"
      :isLoading="eventStore.deleteState.loading" @confirm="confirmDeleteEvent" @cancel="handleCancelDelete" />
  </OrganizerLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-vue-next'
import OrganizerLayout from '@/layouts/OrganizerLayout.vue'
import EventsTable from '@/components/organizer/events/EventsTable.vue'
import EventFormModal from '@/components/organizer/events/EventFormModal.vue'
import DeleteConfirmModal from '@/components/organizer/events/DeleteConfirmModal.vue'
import { useOrganizerStore } from '@/stores/organizerStore'
import { useEventStore } from '@/stores/eventStore'
import type { EventInterface } from '@/types/event'
import { toast } from 'vue-sonner'

const organizerStore = useOrganizerStore()
const eventStore = useEventStore()

const showFormModal = ref(false)
const showDeleteModal = ref(false)
const isEditing = ref(false)
const selectedEvent = ref<EventInterface | null>(null)

onMounted(() => {
  organizerStore.getMyEvents()
})

function openCreateModal() {
  isEditing.value = false
  selectedEvent.value = null
  showFormModal.value = true
}

function openEditModal(event: EventInterface) {
  isEditing.value = true
  selectedEvent.value = event
  showFormModal.value = true
}

function openDeleteModal(event: EventInterface) {
  selectedEvent.value = event
  showDeleteModal.value = true
}

async function handleSubmitEvent(formData: any) {
  try {
    const formDataToSend = new FormData()
    formDataToSend.append('title', formData.title)
    formDataToSend.append('description', formData.description)
    formDataToSend.append('location', formData.location)
    formDataToSend.append('category', formData.category)
    formDataToSend.append('date', formData.date)
    formDataToSend.append('time', formData.time)

    if (formData.banner_image) {
      formDataToSend.append('banner_image', formData.banner_image)
    }

    if (isEditing.value && selectedEvent.value) {
      await eventStore.updateEvent(selectedEvent.value.id, formDataToSend)
    } else {
      await eventStore.createEvent(formDataToSend)
    }

    await organizerStore.getMyEvents()

    toast.success(`Event ${isEditing.value ? 'updated' : 'created'} successfully`, {
      description: new Date().toLocaleString(),
    })

    showFormModal.value = false
    selectedEvent.value = null
  } catch (err) {
    toast.error(`Failed to ${isEditing.value ? 'update' : 'create'} event`, {
      description: `An unexpected error occurred. ${isEditing.value ? eventStore.updateState.error : eventStore.createState.error}`,
    })
    console.error(`Failed to ${isEditing.value ? 'update' : 'create'} event`, err)
  }
}

async function confirmDeleteEvent() {
  try {
    if (!selectedEvent.value) return

    await eventStore.deleteEvent(selectedEvent.value.id)
    await organizerStore.getMyEvents()

    toast.success('Event deleted successfully', {
      description: new Date().toLocaleString(),
    })

    showDeleteModal.value = false
    selectedEvent.value = null
  } catch (err) {
    toast.error('Failed to delete event', {
      description: eventStore.deleteState.error || `An unexpected error occurred. ${err}`,
    })
    console.error('Delete Event Failed: ', eventStore.deleteState.error, err)
  }
}

async function handleTogglePublic(event: EventInterface) {
  try {
    const payload = {
      is_public: !event.is_public,
      is_open: event.is_open
    }

    await eventStore.eventAvailability(event.id, createAvailabilityFormData(payload))
    await organizerStore.getMyEvents()

    toast.success('Event availiability updated', {
      description: `Event is now ${!event.is_public ? 'public' : 'private'}`,
    })
  } catch (err) {
    toast.error('Failed to update event public status', {
      description: eventStore.eventAvl.error || 'An unexpected error occurred',
    })
    console.error('Toggle Public Failed: ', err)
  }
}

async function handleToggleOpen(event: EventInterface) {
  try {
    const payload = {
      is_public: event.is_public,
      is_open: !event.is_open
    }

    await eventStore.eventAvailability(event.id, createAvailabilityFormData(payload))
    await organizerStore.getMyEvents()

    toast.success('Event availiability updated', {
      description: `Event is now ${!event.is_open ? 'open' : 'closed'}`,
    })
  } catch (err) {
    toast.error('Failed to update event open status', {
      description: eventStore.eventAvl.error || 'An unexpected error occurred',
    })
    console.error('Toggle Open Failed: ', err)
  }
}

function createAvailabilityFormData(payload: { is_public: boolean; is_open: boolean }) {
  const formData = new FormData()
  formData.append('is_public', payload.is_public.toString())
  formData.append('is_open', payload.is_open.toString())
  return formData
}

function handleCancelForm() {
  selectedEvent.value = null
}

function handleCancelDelete() {
  selectedEvent.value = null
}
</script>