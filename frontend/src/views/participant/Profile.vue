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
                <p class="text-sm text-gray-600">{{ formatDate(myReg.date) }} • {{ myReg.location }}</p>
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
                      <p class="text-gray-600">{{ formatDate(ticket.date) }} • {{ ticket.location }}</p>
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
      </div>
    </div>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { MoreHorizontal, Plus, Ticket, QrCode } from 'lucide-vue-next'
import ParticipantLayout from '@/layouts/ParticipantLayout.vue'
import { useUserStore } from '@/stores/userStore'
import { useEventRegistrationsStore } from '@/stores/eventRegistrationsStore'
import { formatDate } from '@/utils/format'
import { useTicketStore } from '@/stores/ticketStore'

const userStore = useUserStore()
const eventRegStore = useEventRegistrationsStore()
const ticketStore = useTicketStore()

const user = computed(() => userStore.userState.data)
const myRegistrations = computed(() => eventRegStore.myRegState.data)
const myTickets = computed(() => ticketStore.myTicketState.data)

onMounted(() => {
  userStore.getMyProfile()
  eventRegStore.getMyRegistrations()
  ticketStore.getMyTickets()
})

const activeTab = ref('registrations')

const tabs = [
  { id: 'registrations', name: 'Event Registrations' },
  { id: 'tickets', name: 'Tickets' }
]
</script>