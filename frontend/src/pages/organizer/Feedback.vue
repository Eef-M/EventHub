<template>
  <OrganizerLayout>
    <div class="p-6">
      <h2 class="text-2xl font-semibold mb-4">Event Feedback</h2>

      <div v-if="feedbacks && feedbacks.length" class="space-y-4">
        <div v-for="(feedback, index) in feedbacks" :key="index" class="bg-white p-4 rounded-xl shadow-md">
          <div class="flex justify-between items-center mb-1">
            <div class="font-semibold">
              {{ feedback.username }} ({{ feedback.event_title }})
            </div>
            <div class="flex justify-center gap-2 font-medium">
              <Star class="text-yellow-500" />
              {{ feedback.rating }}
            </div>
          </div>
          <div class="text-sm text-gray-500 mb-2">
            {{ formatDate(feedback.created_at) }}
          </div>
          <p class="text-gray-700">
            {{ feedback.comment }}
          </p>
        </div>
      </div>
      <div v-else-if="organizerStore.feedbackState.loading" class="text-slate-700">Loading...</div>
      <div v-else class="text-slate-700">No feedback available.</div>
    </div>
  </OrganizerLayout>
</template>

<script lang="ts" setup>
import { computed, onMounted } from 'vue';
import OrganizerLayout from '@/layouts/OrganizerLayout.vue';
import { useOrganizerStore } from '@/stores/organizerStore';
import { Star } from 'lucide-vue-next';

const organizerStore = useOrganizerStore()
const feedbacks = computed(() => organizerStore.feedbackState.data)

onMounted(() => {
  organizerStore.getOrganizerEventFeedback()
})

function formatDate(dateStr: string): string {
  const date = new Date(dateStr);
  return date.toLocaleDateString(undefined, {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
}
</script>
