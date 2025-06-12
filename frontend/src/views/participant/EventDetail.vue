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

      <!-- Tikcets Section -->
      <Card v-if="isTicketsLoaded" class="mb-8 shadow-lg border-0">
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

      <Card v-else class="mb-8 shadow-lg border-0">
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
      <!-- Tikcets Section End -->

      <!-- Feedback Section -->
      <Card v-if="isFeedbacksLoaded" class="shadow-lg border-0 bg-gradient-to-br from-slate-50 to-white">
        <CardHeader class="pb-6">
          <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
            <MessageSquareIcon class="w-6 h-6 mr-3 text-primary" />
            Community Feedback
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div class="flex items-center justify-center py-12">
            <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-primary"></div>
            <span class="ml-4 text-slate-600 text-lg">Loading feedback...</span>
          </div>
        </CardContent>
      </Card>

      <Card v-else class="shadow-lg border-0 bg-gradient-to-br from-slate-50 to-white">
        <CardHeader class="pb-6">
          <div class="flex items-center justify-between">
            <CardTitle class="text-2xl font-semibold text-slate-900 flex items-center">
              <MessageSquareIcon class="w-6 h-6 mr-3 text-primary" />
              Feedback
            </CardTitle>
            <Badge v-if="feedbacks && feedbacks.length > 0" variant="secondary" class="px-3 py-1">
              {{ feedbacks.length }} {{ feedbacks.length === 1 ? 'Review' : 'Reviews' }}
            </Badge>
          </div>
        </CardHeader>
        <CardContent>
          <div v-if="feedbacks && feedbacks.length > 0" class="space-y-6">
            <div v-for="feedback in feedbacks" :key="feedback.id"
              class="bg-white p-6 rounded-xl shadow-sm border border-slate-100 hover:shadow-md transition-all duration-300">
              <div class="flex items-start space-x-4">
                <div class="flex-shrink-0">
                  <img class="w-12 h-12 rounded-full object-cover ring-2 ring-slate-100" :src="feedback.avatar_url"
                    :alt="feedback.username" />
                </div>
                <div class="flex-1 min-w-0">
                  <div class="flex items-center justify-between mb-2">
                    <div class="flex items-center space-x-2">
                      <h4 class="text-sm font-semibold text-slate-900">{{ feedback.username }}</h4>
                      <span class="text-xs text-slate-400">•</span>
                      <span class="text-xs text-slate-500">{{ formatFeedbackDate(feedback.created_at) }}</span>
                    </div>
                    <div class="flex items-center space-x-1">
                      <div class="flex">
                        <StarIcon v-for="star in 5" :key="star" :class="[
                          'w-4 h-4 transition-colors',
                          star <= feedback.rating
                            ? 'text-yellow-400 fill-current'
                            : 'text-gray-300'
                        ]" />
                      </div>
                      <span class="text-sm font-medium text-slate-700 ml-1">{{ feedback.rating }}/5</span>
                    </div>
                  </div>
                  <p class="text-slate-700 leading-relaxed text-sm">{{ feedback.comment }}</p>
                  <div class="flex items-center space-x-4 mt-3 pt-3 border-t border-slate-100">
                    <button
                      class="flex items-center space-x-1 text-xs text-slate-500 hover:text-slate-700 transition-colors">
                      <ThumbsUpIcon class="w-3 h-3" />
                      <span>Helpful</span>
                    </button>
                    <button
                      class="flex items-center space-x-1 text-xs text-slate-500 hover:text-slate-700 transition-colors">
                      <MessageCircleIcon class="w-3 h-3" />
                      <span>Reply</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div v-if="feedbacks.length >= 5" class="text-center pt-4">
              <Button variant="outline" class="px-6 py-2">
                Load More Reviews
              </Button>
            </div>
          </div>

          <div v-else class="text-center py-8">
            <MessageSquareIcon class="w-16 h-16 mx-auto text-slate-300 mb-4" />
            <p class="text-slate-500 text-lg italic">No feedback yet.</p>
            <p class="text-slate-400 text-sm mt-2">Be the first to share your thoughts about this event!</p>
            <div class="mt-6">
              <Button class="px-6 py-2 bg-purple-600 hover:bg-purple-700 cursor-pointer">
                <MessageSquareIcon class="w-4 h-4 mr-2" />
                Write First Feedback
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
      <!-- Feedback Section End -->
    </section>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ParticipantLayout from '../../layouts/ParticipantLayout.vue'
import { useEventStore } from '@/stores/eventStore'
import { formatDate, formatFeedbackDate, formatTime } from '@/utils/format'
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
  UsersIcon,
  StarIcon,
  ThumbsUpIcon,
  MessageCircleIcon
} from 'lucide-vue-next'
import { useFeedbackStore } from '@/stores/feedbackStore'
import { useUserStore } from '@/stores/userStore'

const eventStore = useEventStore()
const ticketStore = useTicketStore()
const feedbackStore = useFeedbackStore()
const userStore = useUserStore()

const route = useRoute()
const eventId = route.params.id as string
const userId = userStore.userState.data?.id

const isTicketsLoaded = computed(() => ticketStore.ticketsState.loading)
const isFeedbacksLoaded = computed(() => feedbackStore.feedbacksState.loading)

const event = computed(() => eventStore?.singleEventState?.data)
const tickets = computed(() => {
  const allTickets = ticketStore.ticketsState?.data || []
  return allTickets.filter(ticket => ticket.event_id === eventId)
})
const feedbacks = computed(() => feedbackStore.feedbacksState?.data)

onMounted(() => {
  eventStore.getEventById(eventId)
  ticketStore.getTickets()
  feedbackStore.getFeedbacks(eventId)
})

async function handleBuyTicket(ticketId: string) {
  console.log(`Tikcet ID: ${ticketId}`)
  console.log(`User ID: ${userId}`)
  // buy ticket logic here
}
</script>