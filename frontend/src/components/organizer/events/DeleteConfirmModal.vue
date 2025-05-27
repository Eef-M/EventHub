<template>
  <Dialog :open="isOpen" @update:open="$emit('update:isOpen', $event)">
    <DialogContent class="sm:max-w-sm">
      <DialogHeader>
        <DialogTitle>Delete Event</DialogTitle>
        <DialogDescription>
          Are you sure you want to delete <strong>{{ eventTitle }}</strong>? This action cannot be undone.
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <Button class="cursor-pointer" variant="outline" @click="handleCancel">Cancel</Button>
        <Button class="bg-red-600 hover:bg-red-700 cursor-pointer" :disabled="isLoading" @click="handleConfirm">
          Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogDescription
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'

interface Props {
  isOpen: boolean
  eventTitle?: string
  isLoading?: boolean
}

withDefaults(defineProps<Props>(), {
  eventTitle: '',
  isLoading: false
})

const emit = defineEmits<{
  'update:isOpen': [value: boolean]
  confirm: []
  cancel: []
}>()

function handleConfirm() {
  emit('confirm')
}

function handleCancel() {
  emit('cancel')
  emit('update:isOpen', false)
}
</script>