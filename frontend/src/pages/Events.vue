<template>
  <DefaultLayout>
    <section class="py-10">
      <div class="max-w-7xl mx-auto px-4">
        <h2 class="text-3xl font-semibold text-center mb-10">News Events</h2>
        <div class="grid md:grid-cols-3 gap-6">
          <RouterLink to="/events/id/detail">
            <EventCard />
          </RouterLink>
          <RouterLink to="/events/id/detail">
            <EventCard />
          </RouterLink>
          <RouterLink to="/events/id/detail">
            <EventCard />
          </RouterLink>
        </div>
      </div>
    </section>
  </DefaultLayout>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
import DefaultLayout from '../layouts/DefaultLayout.vue';
import EventCard from '../components/EventCard.vue';

</script>