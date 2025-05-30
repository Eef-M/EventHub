<template>
  <ParticipantLayout>
    <div class="min-h-screen bg-white">
      <div class="bg-slate-50 border-b border-slate-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="py-8">
            <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6">
              <div class="relative">
                <img src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=150&h=150&fit=crop&crop=face"
                  alt="user avatar" class="w-32 h-32 rounded-full object-cover border-4 border-white shadow-lg">
                <div class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 border-2 border-white rounded-full"></div>
              </div>

              <div class="flex-1 text-center sm:text-left">
                <h1 class="text-3xl font-bold text-gray-900">{{ user.first_name }} {{ user.last_name }}</h1>
                <p class="text-slate-600 mt-1 italic">@{{ user.username }}</p>
                <p class="text-slate-600 mt-1 font-bold">{{ user.role }}</p>

                <div class="flex flex-col sm:flex-row gap-3 mt-4">
                  <Button
                    class="bg-purple-600 text-white hover:bg-purple-700 px-6 py-2 rounded-full trasnsition cursor-pointer">
                    Get in touch
                  </Button>
                  <Button variant="outline" class="px-6 py-2 rounded-full border-slate-300 cursor-pointer">
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
            <Card v-for="event in eventRegistrations" :key="event.id"
              class="overflow-hidden hover:shadow-lg transition-shadow">
              <div class="aspect-video bg-gradient-to-br from-blue-400 to-purple-600 relative">
                <img v-if="event.image" :src="event.image" :alt="event.title" class="w-full h-full object-cover">
                <div class="absolute top-3 right-3">
                  <Badge :class="event.status === 'registered' ? 'bg-green-600' : 'bg-red-600'">
                    {{ event.status }}
                  </Badge>
                </div>
              </div>
              <CardHeader>
                <CardTitle class="text-lg">{{ event.title }}</CardTitle>
                <p class="text-sm text-gray-600">{{ event.date }} • {{ event.location }}</p>
              </CardHeader>
              <CardContent>
                <p class="text-gray-700 text-sm">{{ event.description }}</p>
                <div class="flex justify-between items-center mt-4">
                  <span class="text-sm text-gray-500">{{ event.attendees }} attendees</span>
                  <Button variant="outline" size="sm">View Details</Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>

        <div v-if="activeTab === 'tickets'" class="space-y-6">
          <div class="flex justify-between items-center">
            <h2 class="text-2xl font-bold text-gray-900">My Tickets</h2>
            <Button class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg">
              <Ticket class="w-4 h-4 mr-2" />
              Buy Tickets
            </Button>
          </div>

          <div class="space-y-4">
            <Card v-for="ticket in tickets" :key="ticket.id" class="hover:shadow-md transition-shadow">
              <CardContent class="p-6">
                <div class="flex items-center justify-between">
                  <div class="flex items-center space-x-4">
                    <div
                      class="w-16 h-16 bg-gradient-to-br from-green-400 to-blue-500 rounded-lg flex items-center justify-center">
                      <Ticket class="w-8 h-8 text-white" />
                    </div>
                    <div>
                      <h3 class="font-semibold text-lg">{{ ticket.eventName }}</h3>
                      <p class="text-gray-600">{{ ticket.date }} • {{ ticket.venue }}</p>
                      <p class="text-sm text-gray-500">Ticket ID: {{ ticket.ticketId }}</p>
                    </div>
                  </div>
                  <div class="text-right">
                    <Badge :class="ticket.status === 'valid' ? 'bg-blue-600' : 'bg-yellow-600'">
                      {{ ticket.status }}
                    </Badge>
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
      </div>
    </div>
  </ParticipantLayout>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { MoreHorizontal, Plus, Ticket, QrCode } from 'lucide-vue-next'
import ParticipantLayout from '@/layouts/ParticipantLayout.vue'
import { useUserStore } from '@/stores/userStore'

const userStore = useUserStore()

const user = computed(() => userStore.userState.data)

onMounted(() => {
  userStore.getMyProfile()
})

const activeTab = ref('registrations')

const tabs = [
  { id: 'registrations', name: 'Event Registrations' },
  { id: 'tickets', name: 'Tickets' }
]

// TEMPORARY DATA !
// have to fix the backend part first if you want to use the API!
const eventRegistrations = ref([
  {
    id: 1,
    title: 'Tech Conference 2024',
    date: 'June 15, 2024',
    location: 'Jakarta Convention Center',
    description: 'Annual technology conference featuring the latest innovations in software development.',
    attendees: 1250,
    status: 'registered',
    image: 'https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=400&h=225&fit=crop'
  },
  {
    id: 2,
    title: 'Startup Pitch Night',
    date: 'July 8, 2024',
    location: 'Bandung Creative Hub',
    description: 'Evening event where innovative startups present their ideas to investors.',
    attendees: 85,
    status: 'cancelled',
    image: 'https://images.unsplash.com/photo-1559136555-9303baea8ebd?w=400&h=225&fit=crop'
  },
  {
    id: 3,
    title: 'Design Workshop',
    date: 'August 20, 2024',
    location: 'Yogyakarta Design Center',
    description: 'Hands-on workshop covering modern UI/UX design principles and tools.',
    attendees: 45,
    status: 'registered',
    image: 'https://images.unsplash.com/photo-1531482615713-2afd69097998?w=400&h=225&fit=crop'
  }
])

const tickets = ref([
  {
    id: 1,
    eventName: 'Music Festival 2024',
    date: 'September 15, 2024',
    venue: 'GBK Stadium, Jakarta',
    ticketId: 'MF2024-001234',
    price: 150,
    status: 'valid'
  },
  {
    id: 2,
    eventName: 'Food & Culture Expo',
    date: 'October 3, 2024',
    venue: 'Yogyakarta Exhibition Hall',
    ticketId: 'FCE2024-005678',
    price: 75,
    status: 'valid'
  },
  {
    id: 3,
    eventName: 'Art Gallery Opening',
    date: 'May 10, 2024',
    venue: 'Museum Nasional',
    ticketId: 'AGO2024-009876',
    price: 25,
    status: 'used'
  }
])
</script>