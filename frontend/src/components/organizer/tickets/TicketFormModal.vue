<template>
  <Dialog :open="isOpen" @update:open="$emit('update:isOpen', $event)">
    <DialogContent class="sm:max-w-md">
      <DialogHeader>
        <DialogTitle>{{ isEditing ? "Edit Ticket" : "Add Ticket" }}</DialogTitle>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="grid gap-2">
          <Label for="name">Ticket Name</Label>
          <Input v-model="form.name" id="name" placeholder="e.g. VIP" />
        </div>
        <div class="grid gap-2">
          <Label for="description">Description</Label>
          <Input v-model="form.description" id="description" placeholder="desc..." />
        </div>
        <div class="grid gap-2">
          <Label for="price">Price</Label>
          <Input v-model.number="form.price" id="price" type="number" />
        </div>
        <div class="grid gap-2">
          <Label for="quota">Quota</Label>
          <Input v-model.number="form.quota" id="quota" type="number" />
        </div>
        <div class="grid gap-2">
          <Label for="event">Event</Label>
          <Select v-model="form.event_id">
            <SelectTrigger id="event" class="w-full">
              <SelectValue placeholder="Select an event" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem v-for="event in events" :key="event.id" :value="event.id">
                {{ event.title }}
              </SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>
      <DialogFooter>
        <Button variant="outline" @click="handleCancel" class="cursor-pointer">Cancel</Button>
        <Button :disabled="isLoading"
          :class="isEditing ? 'bg-slate-600 hover:bg-slate-700 cursor-pointer' : 'bg-purple-600 hover:bg-purple-700 cursor-pointer'"
          @click="handleSubmit">
          {{ isEditing ? "Update" : "Create" }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
<script setup lang="ts">
import { reactive, watch } from "vue";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import type { Ticket } from "@/types/ticket";
import type { EventInterface } from "@/types/event";

interface Props {
  isOpen: boolean
  isEditing: boolean
  ticketData?: Ticket | null
  events: EventInterface[]
  isLoading?: boolean
}

interface FormData {
  event_id: string
  name: string
  description?: string
  price: number
  quota: number
}

const props = withDefaults(defineProps<Props>(), {
  ticketData: null,
  isLoading: false
})

const emit = defineEmits<{
  'update:isOpen': [value: boolean]
  submit: [formData: FormData]
  cancel: []
}>()

const defaultForm: FormData = {
  event_id: "",
  name: "",
  description: "",
  price: 0,
  quota: 0,
};

const form = reactive({ ...defaultForm });

watch([() => props.isOpen, () => props.isEditing, () => props.ticketData], () => {
  if (props.isOpen) {
    if (props.isEditing && props.ticketData) {
      Object.assign(form, {
        event_id: props.ticketData.event_id,
        name: props.ticketData.name,
        description: props.ticketData.description,
        price: props.ticketData.price,
        quota: props.ticketData.quota,
      });
    } else {
      Object.assign(form, defaultForm)
    }
  }
})

function handleSubmit() {
  emit('submit', { ...form })
}

function handleCancel() {
  emit('cancel')
  emit('update:isOpen', false)
}
</script>