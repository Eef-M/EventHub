<template>
  <div v-if="isOpen"
    class="fixed inset-0 bg-black/50 backdrop-blur-sm transition-opacity flex items-center justify-center z-50"
    @click="closeModal">
    <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4" @click.stop>
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl font-bold text-gray-900">Ticket QR Code</h2>
        <button @click="closeModal" class="text-gray-400 hover:text-gray-600 transition-colors">
          <X class="w-6 h-6" />
        </button>
      </div>

      <div class="text-center">
        <div class="mb-4">
          <h3 class="font-semibold text-lg text-gray-900">{{ ticket?.title }}</h3>
          <p class="text-gray-600 text-sm">
            {{ ticket?.date ? formatDate(ticket.date) : 'Date TBD' }} • {{ ticket?.location || 'Location TBD' }}
          </p>
        </div>

        <div class="bg-gray-50 p-6 rounded-lg mb-4">
          <!-- Placeholder for QR Code - temporarily display ticket code -->
          <div class="w-32 h-32 mx-auto bg-gray-200 rounded-lg flex items-center justify-center mb-4">
            <QrCode class="w-16 h-16 text-gray-400" />
          </div>
          <p class="text-lg font-mono font-bold text-gray-900">{{ ticket?.ticket_code || 'N/A' }}</p>
          <p class="text-sm text-gray-500 mt-1">Ticket Code</p>
        </div>

        <p class="text-xs text-gray-500">
          Present this QR code at the event entrance for verification
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { MyTicket } from '@/types/ticket';
import { formatDate } from '@/utils/format';
import { X, QrCode } from 'lucide-vue-next'

interface Props {
  isOpen: boolean
  ticket: MyTicket | null
}

interface Emits {
  (e: 'close'): void
}

defineProps<Props>()
const emit = defineEmits<Emits>()

const closeModal = () => {
  emit('close')
}
</script>