<template>
  <Dialog :open="isOpen" @update:open="$emit('update:isOpen', $event)">
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
              <div v-if="eventData?.banner_url" class="flex flex-col items-center gap-1">
                <p class="text-sm text-muted-foreground">Current image:</p>
                <img :src="eventData.banner_url" alt="Current Banner"
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
        <Button variant="outline" @click="handleCancel" class="cursor-pointer">Cancel</Button>
        <Button :disabled="isLoading" @click="handleSubmit"
          :class="isEditing ? 'bg-slate-600 hover:bg-slate-700 cursor-pointer' : 'bg-purple-600 hover:bg-purple-700 cursor-pointer'">
          {{ isEditing ? "Update" : "Create" }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { computed, reactive, watch } from 'vue'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { Label } from '@/components/ui/label'
import type { EventInterface } from '@/types/event'
import { formatTime } from '@/utils/format'

interface Props {
  isOpen: boolean
  isEditing: boolean
  eventData?: EventInterface | null
  isLoading?: boolean
}

interface FormData {
  title: string
  description: string
  location: string
  category: string
  date: string
  time: string
  banner_image: File | null
}

const props = withDefaults(defineProps<Props>(), {
  eventData: null,
  isLoading: false
})

const emit = defineEmits<{
  'update:isOpen': [value: boolean]
  submit: [formData: FormData]
  cancel: []
}>()

const defaultForm: FormData = {
  title: '',
  description: '',
  location: '',
  category: '',
  date: '',
  time: '',
  banner_image: null,
}

const form = reactive({ ...defaultForm })

watch([() => props.isOpen, () => props.isEditing, () => props.eventData], () => {
  if (props.isOpen) {
    if (props.isEditing && props.eventData) {
      const isoDate = new Date(props.eventData.date)
      const formattedDate = isoDate.toISOString().split('T')[0]
      Object.assign(form, {
        title: props.eventData.title,
        description: props.eventData.description,
        location: props.eventData.location,
        category: props.eventData.category,
        date: formattedDate,
        time: formatTime(props.eventData.time),
        banner_image: null
      })
    } else {
      Object.assign(form, defaultForm)
    }
  }
})

function handleBannerChange(e: EventInterface & { target: HTMLInputElement }) {
  form.banner_image = e.target.files?.[0] || null
}

function handleSubmit() {
  emit('submit', { ...form })
}

function handleCancel() {
  emit('cancel')
  emit('update:isOpen', false)
}

const imagePreviewUrl = computed(() =>
  form.banner_image ? URL.createObjectURL(form.banner_image) : null
)
</script>