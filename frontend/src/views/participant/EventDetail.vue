<template>
  <ParticipantLayout>
    <section class="py-8 max-w-7xl mx-auto px-4">
      <div class="rounded-3xl overflow-hidden shadow-lg mb-8 bg-gradient-to-r from-purple-500 to-slate-500">
        <img v-if="event?.banner_url" :src="event.banner_url" alt="Event Banner"
          class="w-full h-64 md:h-80 object-cover" />
        <div v-else class="w-full h-64 md:h-80 flex items-center justify-center text-white">
          <div class="text-center">
            <CalendarIcon class="w-16 h-16 mx-auto mb-4 opacity-70" />
            <p class="text-lg font-medium">Event Banner</p>
          </div>
        </div>
      </div>

      <Card class="mb-8 shadow-lg border-0 bg-gradient-to-br from-white to-slate-50">
        <CardHeader class="pb-4">
          <CardTitle class="text-4xl font-bold text-slate-900 leading-tight">
            {{ event?.title || 'Loading...' }}
          </CardTitle>
        </CardHeader>
        <CardContent class="space-y-4">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="flex items-center space-x-3">
              <Badge variant="secondary" class="px-3 py-1">
                <TagIcon class="w-4 h-4 mr-2" />
                {{ event?.category || 'N/A' }}
              </Badge>
            </div>
            <div class="flex items-center space-x-3 text-slate-600">
              <MapPinIcon class="w-5 h-5 text-primary" />
              <span class="font-medium">{{ event?.location || 'TBA' }}</span>
            </div>
            <div class="flex items-center space-x-3 text-gray-600">
              <CalendarIcon class="w-5 h-5 text-primary" />
              <span class="font-medium">{{ event ? formatDate(event.date) : 'TBA' }}</span>
            </div>
            <div class="flex items-center space-x-3 text-slate-600">
              <ClockIcon class="w-5 h-5 text-primary" />
              <span class="font-medium">{{ event ? formatTime(event.time) : 'TBA' }}</span>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card class="mb-8 shadow-lg border-0">
        <CardHeader>
          <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
            <FileTextIcon class="w-6 h-6 mr-3 text-primary" />
            Description
          </CardTitle>
        </CardHeader>
        <CardContent>
          <p class="text-slate-700 leading-relaxed text-lg">
            {{ event?.description || 'No description available.' }}
          </p>
        </CardContent>
      </Card>

      <Card v-if="isTicketsLoaded" class="mb-8 shadow-lg border-0">
        <CardHeader>
          <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
            <TicketIcon class="w-6 h-6 mr-3 text-primary" />
            Available Tickets
            <Badge v-if="tickets.length > 0" class="ml-3 bg-yellow-600">
              {{ tickets.length }} {{ tickets.length === 1 ? 'Type' : 'Types' }}
            </Badge>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div v-if="tickets.length > 0" class="grid gap-4">
            <Card v-for="ticket in tickets" :key="ticket.id"
              class="border-2 hover:border-purple-400 transition-all duration-300 hover:shadow-md">
              <CardContent class="p-6">
                <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
                  <div class="flex-1">
                    <div class="flex items-center gap-3 mb-2">
                      <h4 class="text-xl font-bold text-slate-900">{{ ticket.name }}</h4>
                      <Badge variant="outline" class="text-xs">
                        {{ ticket.quota }} left
                      </Badge>
                    </div>
                    <p class="text-gray-600 mb-2">{{ ticket.description }}</p>
                    <div class="flex items-center text-sm text-slate-500">
                      <UsersIcon class="w-4 h-4 mr-1" />
                      Quota: {{ ticket.quota }}
                    </div>
                  </div>
                  <div class="flex flex-col items-end gap-3">
                    <div class="text-right">
                      <p class="text-3xl font-bold text-primary">
                        ${{ ticket.price }}
                      </p>
                      <p class="text-sm text-gray-500">per ticket</p>
                    </div>
                    <Button size="lg"
                      class="w-full md:w-auto px-8 py-3 font-semibold bg-purple-600 hover:bg-purple-700 cursor-pointer transition"
                      @click="handleBuyTicket(ticket.id)" :disabled="ticket.quota === 0">
                      <ShoppingCartIcon class="w-4 h-4 mr-2" />
                      {{ ticket.quota === 0 ? 'Sold Out' : 'Buy Ticket' }}
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
          <div v-else class="text-center py-8">
            <TicketIcon class="w-16 h-16 mx-auto text-slate-300 mb-4" />
            <p class="text-slate-500 text-lg">No tickets available for this event.</p>
          </div>
        </CardContent>
      </Card>

      <Card v-else class="mb-8 shadow-lg border-0">
        <CardHeader>
          <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
            <TicketIcon class="w-6 h-6 mr-3 text-primary" />
            Available Tickets
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div class="flex items-center justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
            <span class="ml-3 text-slate-600">Loading tickets...</span>
          </div>
        </CardContent>
      </Card>

      <Card class="shadow-lg border-0">
        <CardHeader>
          <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
            <MessageSquareIcon class="w-6 h-6 mr-3 text-primary" />
            Feedback
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div class="text-center py-8">
            <MessageSquareIcon class="w-16 h-16 mx-auto text-slate-300 mb-4" />
            <p class="text-slate-500 text-lg italic">No feedback yet.</p>
            <p class="text-slate-400 text-sm mt-2">Be the first to share your thoughts about this event!</p>
          </div>
        </CardContent>
      </Card>
    </section>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import ParticipantLayout from '../../layouts/ParticipantLayout.vue'
import { useEventStore } from '@/stores/eventStore'
import { formatDate, formatTime } from '@/utils/format'
import { useTicketStore } from '@/stores/ticketStore'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import {
  CalendarIcon,
  ClockIcon,
  MapPinIcon,
  TagIcon,
  TicketIcon,
  FileTextIcon,
  MessageSquareIcon,
  ShoppingCartIcon,
  UsersIcon
} from 'lucide-vue-next'

const route = useRoute()
const eventId = route.params.id as string
const eventStore = useEventStore()
const ticketStore = useTicketStore()
const isTicketsLoaded = ref(false)

const event = computed(() => eventStore?.singleEventState?.data)
const tickets = computed(() => {
  const allTickets = ticketStore.ticketsState?.data || []
  return allTickets.filter(ticket => ticket.event_id === eventId)
})

onMounted(async () => {
  try {
    await eventStore.getEventById(eventId)
    await ticketStore.getTickets()
    isTicketsLoaded.value = true
  } catch (error) {
    console.error('Error loading event details:', error)
    isTicketsLoaded.value = true
  }
})

async function handleBuyTicket(ticketId: string) {
  console.log(ticketId)
  // buy ticket logic here
}
</script>