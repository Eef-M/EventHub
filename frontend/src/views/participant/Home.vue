<template>
  <ParticipantLayout>
    <section class="bg-zinc-100 py-12 px-4 md:px-8">
      <div class="max-w-7xl mx-auto flex flex-col md:flex-row items-center justify-between">
        <div class="text-left mb-8 md:mb-0 md:w-1/2">
          <h1 class="text-3xl md:text-4xl lg:text-5xl font-bold text-purple-900 mb-4">Discover Exciting Events</h1>
          <p class="text-base text-gray-600 mb-6">
            Explore and manage events easily. Find events that match your interests and register in just a few clicks.
          </p>
          <button
            class="px-6 py-3 bg-purple-600 text-white rounded-md shadow hover:bg-purple-700 transition cursor-pointer">
            Browse Events
          </button>
        </div>

        <div class="md:w-1/2 relative">
          <div
            class="bg-purple-100 rounded-full w-64 h-64 md:w-80 md:h-80 mx-auto relative flex items-center justify-center">
            <img
              src="https://images.pexels.com/photos/2952834/pexels-photo-2952834.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
              alt="Person using smartphone" class="w-48 h-48 md:w-64 md:h-64 object-cover rounded-full" />
          </div>
        </div>
      </div>
    </section>

    <section class="bg-white py-16">
      <div class="max-w-7xl mx-auto px-4">
        <h2 class="text-3xl font-semibold text-center mb-10">News Events</h2>

        <div class="grid md:grid-cols-3 gap-6 mb-10">
          <RouterLink v-for="event in topEvents" :key="event.id" :to="`/events/${event.id}/detail`">
            <EventCard :event="event" />
          </RouterLink>
        </div>

        <div class="text-center">
          <RouterLink to="/events">
            <button
              class="px-6 py-3 bg-purple-600 text-white font-medium rounded-lg hover:bg-purple-700 transition shadow cursor-pointer">
              See more
            </button>
          </RouterLink>
        </div>
      </div>
    </section>

    <section class="bg-gray-100 py-16">
      <div class="max-w-7xl mx-auto px-4 text-center">
        <h2 class="text-3xl font-semibold mb-8">Popular Categories</h2>
        <div class="grid md:grid-cols-4 gap-6">
          <div class="bg-white p-6 rounded shadow">🎓 Workshop</div>
          <div class="bg-white p-6 rounded shadow">💼 Seminar</div>
          <div class="bg-white p-6 rounded shadow">🎤 Conference</div>
          <div class="bg-white p-6 rounded shadow">🤝 Meetup</div>
        </div>
      </div>
    </section>

    <section class="bg-white py-16">
      <div class="max-w-5xl mx-auto px-4 text-center">
        <h2 class="text-3xl font-semibold mb-10">How It Works</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <div>
            <div class="text-5xl mb-4">🔍</div>
            <h3 class="font-bold text-xl mb-2">Find Events</h3>
            <p class="text-gray-600">Browse curated events that match your interests and location.</p>
          </div>
          <div>
            <div class="text-5xl mb-4">📝</div>
            <h3 class="font-bold text-xl mb-2">Register Easily</h3>
            <p class="text-gray-600">Register with just a few clicks and get reminders via email.</p>
          </div>
          <div>
            <div class="text-5xl mb-4">🎉</div>
            <h3 class="font-bold text-xl mb-2">Join & Enjoy</h3>
            <p class="text-gray-600">Attend events and make meaningful connections.</p>
          </div>
        </div>
      </div>
    </section>

    <section class="bg-purple-600 py-16 text-white text-center">
      <div class="max-w-4xl mx-auto px-4">
        <h2 class="text-3xl font-bold mb-4">Ready to Host Your Event?</h2>
        <p class="mb-6">Join EventHub and start managing your own events today!</p>
        <button class="bg-white text-purple-600 font-semibold px-6 py-3 rounded-md hover:bg-gray-100 transition">
          Create an Event
        </button>
      </div>
    </section>

    <section class="bg-white py-16">
      <div class="max-w-6xl mx-auto px-4 text-center">
        <h2 class="text-3xl font-semibold mb-12">What Our Users Say</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <div class="bg-gray-100 p-6 rounded-lg shadow">
            <p class="text-gray-700 italic">"EventHub made it so easy to find tech meetups near me. The registration was
              seamless!"</p>
            <h4 class="font-semibold mt-4">– Sarah, Developer</h4>
          </div>
          <div class="bg-gray-100 p-6 rounded-lg shadow">
            <p class="text-gray-700 italic">"I hosted a small workshop and found more attendees than expected thanks to
              EventHub!"</p>
            <h4 class="font-semibold mt-4">– James, Organizer</h4>
          </div>
          <div class="bg-gray-100 p-6 rounded-lg shadow">
            <p class="text-gray-700 italic">"Nice UI, good experience. Definitely using it again for future events."</p>
            <h4 class="font-semibold mt-4">– Lina, Freelancer</h4>
          </div>
        </div>
      </div>
    </section>

    <section class="bg-gray-50 py-16">
      <div class="max-w-6xl mx-auto px-4 text-center">
        <h2 class="text-3xl font-semibold mb-12">Our Impact So Far</h2>
        <div class="grid md:grid-cols-4 gap-6 text-purple-700 font-bold text-2xl">
          <div>
            <p>1,200+</p>
            <span class="block text-sm text-gray-500 font-normal mt-2">Registered Users</span>
          </div>
          <div>
            <p>450+</p>
            <span class="block text-sm text-gray-500 font-normal mt-2">Events Listed</span>
          </div>
          <div>
            <p>90+</p>
            <span class="block text-sm text-gray-500 font-normal mt-2">Cities Covered</span>
          </div>
          <div>
            <p>98%</p>
            <span class="block text-sm text-gray-500 font-normal mt-2">Positive Feedback</span>
          </div>
        </div>
      </div>
    </section>

  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue';
import EventCard from '../../components/participant/EventCard.vue';
import { useEventStore } from '../../stores/eventStore';
import ParticipantLayout from '../../layouts/ParticipantLayout.vue';

const eventStore = useEventStore()

onMounted(() => {
  eventStore.getEvents()
})

const topEvents = computed(() => eventStore.eventsState.data.slice(0, 3))
</script>
