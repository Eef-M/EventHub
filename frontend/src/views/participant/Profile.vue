<template>
  <ParticipantLayout>
    <div class="min-h-screen bg-white">
      <div class="bg-slate-50 border-b border-slate-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="py-8">
            <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6">
              <div class="relative">
                <img :src="user?.avatar_url" alt="user avatar"
                  class="w-32 h-32 rounded-full object-cover border-4 border-white shadow-lg">
                <div class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 border-2 border-white rounded-full"></div>
              </div>

              <div class="flex-1 text-center sm:text-left">
                <h1 class="text-3xl font-bold text-gray-900">{{ user?.first_name }} {{ user?.last_name }}</h1>
                <p class="text-slate-600 mt-1 italic">@{{ user?.username }}</p>
                <p class="text-slate-600 mt-1 font-bold">{{ user?.role }}</p>

                <div class="flex flex-col sm:flex-row gap-3 mt-4">
                  <Button
                    class="bg-purple-600 text-white hover:bg-purple-700 px-6 py-2 rounded-full transition cursor-pointer">
                    Get in touch
                  </Button>
                  <Button variant="outline" class="px-6 py-2 rounded-full border-slate-300 cursor-pointer"
                    @click="goToEditProfile">
                    Edit Profile
                  </Button>
                  <Button variant="ghost" size="icon" class="w-10 h-10 rounded-full">
                    <MoreHorizontal class="w-5 h-5" />
                  </Button>
                </div>
              </div>
            </div>
          </div>

          <div class="border-b border-gray-200">
            <nav class="flex space-x-8">
              <button v-for="tab in tabs" :key="tab.id" @click="activeTab = tab.id" :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors cursor-pointer',
                activeTab === tab.id
                  ? 'border-purple-500 text-purple-600'
                  : 'border-transparent text-slate-500 hover:text-slate-700 hover:border-slate-300'
              ]">
                {{ tab.name }}
              </button>
            </nav>
          </div>
        </div>
      </div>

      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div v-if="activeTab === 'registrations'" class="space-y-6">
          <div class="flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">Event Registrations</h2>
            <Button class="bg-purple-600 hover:bg-purple-700 text-white px-4 py-2 rounded-lg transition cursor-pointer">
              <Plus class="w-4 h-4 mr-2" />
              Register New Event
            </Button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <Card v-for="myReg in myRegistrations" :key="myReg.registration_id"
              class="overflow-hidden hover:shadow-lg transition-shadow">
              <div class="aspect-video bg-gradient-to-br from-blue-400 to-purple-600 relative">
                <img v-if="myReg.banner_url" :src="myReg.banner_url" :alt="myReg.title"
                  class="w-full h-full object-cover">
                <div class="absolute top-3 right-3">
                  <Badge :class="myReg.status === 'registered' ? 'bg-green-600' : 'bg-red-600'">
                    {{ myReg.status }}
                  </Badge>
                </div>
              </div>
              <CardHeader>
                <CardTitle class="text-lg">{{ myReg.title }}</CardTitle>
                <p class="text-sm text-gray-600">{{ formatDate(myReg.date) }} â€¢ {{ myReg.location }}</p>
              </CardHeader>
              <CardContent>
                <p class="text-gray-700 text-sm">{{ myReg.description }}</p>
                <div class="flex justify-between items-center mt-4">
                  <span class="text-sm text-gray-500">{{ myReg.total_registrations }} attendees</span>
                  <Button variant="outline" size="sm">View Details</Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>

        <div v-if="activeTab === 'tickets'" class="space-y-6">
          <div class="flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">My Tickets</h2>
            <Button class="bg-slate-600 hover:bg-slate-700 text-white px-4 py-2 rounded-lg cursor-pointer">
              <Ticket class="w-4 h-4 mr-2" />
              Buy Tickets
            </Button>
          </div>

          <div class="space-y-4">
            <Card v-for="ticket in myTickets" :key="ticket.ticket_id" class="hover:shadow-md transition-shadow">
              <CardContent class="p-6">
                <div class="flex items-center justify-between">
                  <div class="flex items-center space-x-4">
                    <div
                      class="w-16 h-16 bg-gradient-to-br from-slate-400 to-purple-500 rounded-lg flex items-center justify-center">
                      <Ticket class="w-8 h-8 text-white" />
                    </div>
                    <div>
                      <h3 class="font-semibold text-lg">{{ ticket.title }}</h3>
                      <p class="text-gray-600">{{ formatDate(ticket.date) }} â€¢ {{ ticket.location }}</p>
                      <p class="text-sm text-gray-500">Ticket Code: {{ ticket.ticket_code }}</p>
                    </div>
                  </div>
                  <div class="text-right">
                    <p class="text-lg font-bold mt-2">${{ ticket.price }}</p>
                    <Button variant="outline" size="sm" class="mt-2">
                      <QrCode class="w-4 h-4 mr-1" />
                      Show QR
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>

        <div v-if="activeTab === 'events'" class="space-y-6">
          <div class="flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">My Events</h2>
            <Button class="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg transition cursor-pointer">
              <Plus class="w-4 h-4 mr-2" />
              Create New Event
            </Button>
          </div>

          <div v-if="organizerStore.eventsState.loading" class="flex justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
          </div>

          <div v-else-if="organizerStore.eventsState.error" class="text-center py-8">
            <p class="text-red-600">{{ organizerStore.eventsState.error }}</p>
          </div>

          <div v-else-if="myEvents.length === 0" class="text-center py-12">
            <div class="w-16 h-16 mx-auto mb-4 bg-gray-100 rounded-full flex items-center justify-center">
              <Calendar class="w-8 h-8 text-gray-400" />
            </div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">No events yet</h3>
            <p class="text-gray-500 mb-6">Create your first event to get started.</p>
            <Button class="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-3 rounded-lg transition cursor-pointer">
              <Plus class="w-5 h-5 mr-2" />
              Create Event
            </Button>
          </div>

          <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <Card v-for="event in myEvents" :key="event.id"
              class="overflow-hidden hover:shadow-xl transition-all duration-300 group border-0 shadow-lg">
              <div class="relative">
                <div
                  class="aspect-[16/9] bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 relative overflow-hidden">
                  <img v-if="event.banner_url" :src="event.banner_url" :alt="event.title"
                    class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500">
                  <div v-else class="w-full h-full flex items-center justify-center">
                    <Calendar class="w-16 h-16 text-white opacity-70" />
                  </div>

                  <div class="absolute top-4 right-4">
                    <Badge variant="secondary" class="bg-black/20 text-white border-0 backdrop-blur-sm font-medium">
                      {{ event.category }}
                    </Badge>
                  </div>

                  <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>

                  <div class="absolute bottom-4 left-4 flex items-center space-x-1 text-white">
                    <MapPin class="w-4 h-4" />
                    <span class="text-sm">{{ event.location }}</span>
                  </div>
                </div>

                <CardContent class="p-6">
                  <div class="space-y-4">
                    <div>
                      <h3
                        class="text-xl font-bold text-gray-900 group-hover:text-indigo-600 transition-colors line-clamp-2">
                        {{ event.title }}
                      </h3>
                      <div class="flex items-center space-x-2 mt-2 text-gray-500">
                        <Clock class="w-4 h-4" />
                        <span class="text-sm">{{ formatDate(event.date) }}</span>
                        <span class="text-xs">â€¢</span>
                        <span class="text-sm">{{ formatTime(event.time) || 'Time TBD' }}</span>
                      </div>
                    </div>

                    <p class="text-gray-600 text-sm line-clamp-3 leading-relaxed">
                      {{ event.description }}
                    </p>

                    <div class="flex space-x-2 pt-2">
                      <Button variant="outline" size="sm"
                        class="flex-1 hover:bg-indigo-50 hover:border-indigo-300 cursor-pointer">
                        <Eye class="w-4 h-4 mr-2" />
                        View
                      </Button>
                      <Button variant="outline" size="sm"
                        class="flex-1 hover:bg-purple-50 hover:border-purple-300 cursor-pointer">
                        <Edit class="w-4 h-4 mr-2" />
                        Edit
                      </Button>
                      <Button variant="outline" size="sm"
                        class="hover:bg-green-50 hover:border-green-300 cursor-pointer">
                        <BarChart3 class="w-4 h-4" />
                      </Button>
                    </div>
                  </div>
                </CardContent>
              </div>
            </Card>
          </div>
        </div>
      </div>
    </div>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { MoreHorizontal, Plus, Ticket, QrCode, Calendar, MapPin, Clock, Eye, Edit, BarChart3 } from 'lucide-vue-next'
import ParticipantLayout from '@/layouts/ParticipantLayout.vue'
import { useUserStore } from '@/stores/userStore'
import { useEventRegistrationsStore } from '@/stores/eventRegistrationsStore'
import { useOrganizerStore } from '@/stores/organizerStore'
import { formatDate, formatTime } from '@/utils/format'
import { useTicketStore } from '@/stores/ticketStore'

const router = useRouter()
const userStore = useUserStore()
const eventRegStore = useEventRegistrationsStore()
const ticketStore = useTicketStore()
const organizerStore = useOrganizerStore()

const user = computed(() => userStore.userState.data)
const myRegistrations = computed(() => eventRegStore.myRegState.data)
const myTickets = computed(() => ticketStore.myTicketState.data)
const myEvents = computed(() => organizerStore.eventsState.data)

onMounted(() => {
  userStore.getMyProfile()
  eventRegStore.getMyRegistrations()
  ticketStore.getMyTickets()

  // Load organizer events if user is organizer
  if (user.value?.role === 'organizer') {
    organizerStore.getMyEvents()
  }
})

const activeTab = ref('registrations')

const tabs = computed(() => {
  const baseTabs = [
    { id: 'registrations', name: 'Event Registrations' },
    { id: 'tickets', name: 'Tickets' }
  ]

  // Add My Events tab if user is organizer
  // My Events feature is not finished yet, need to fix the backend first
  if (user.value?.role === 'organizer') {
    baseTabs.push({ id: 'events', name: 'My Events' })
  }

  return baseTabs
})

const goToEditProfile = () => {
  router.push('/profile/edit')
}
</script>