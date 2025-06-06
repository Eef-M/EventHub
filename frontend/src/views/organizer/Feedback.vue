<template>
  <OrganizerLayout>
    <div class="p-6 bg-gray-50 min-h-screen">
      <div class="max-w-8xl mx-auto">
        <div class="mb-8">
          <h2 class="text-3xl font-bold text-gray-800 mb-2">Event Feedback</h2>
          <p class="text-gray-600">Review feedback from your event attendees</p>
        </div>

        <div v-if="feedbacks && feedbacks.length" class="space-y-6">
          <div v-for="(feedback, index) in feedbacks" :key="index"
            class="bg-white rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 border border-slate-100 overflow-hidden">
            <div class="bg-gradient-to-r from-purple-50 to-slate-50 px-6 py-4 border-b border-slate-100">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <div
                    class="w-10 h-10 bg-gradient-to-br from-purple-500 to-slate-600 rounded-full flex items-center justify-center">
                    <span class="text-white font-semibold text-sm">
                      {{ feedback.username.charAt(0).toUpperCase() }}
                    </span>
                  </div>
                  <div>
                    <h3 class="font-semibold text-gray-800 text-lg">{{ feedback.username }}</h3>
                    <p class="text-sm text-slate-600">{{ feedback.event_title }}</p>
                  </div>
                </div>

                <div class="flex items-center space-x-2 bg-white px-3 py-2 rounded-full shadow-sm">
                  <div class="flex items-center space-x-1">
                    <Star v-for="star in 5" :key="star"
                      :class="star <= feedback.rating ? 'text-yellow-400 fill-current' : 'text-slate-300'"
                      class="w-4 h-4" />
                  </div>
                  <span class="font-semibold text-gray-700 text-sm">{{ feedback.rating }}/5</span>
                </div>
              </div>
            </div>

            <div class="px-6 py-5">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center space-x-2 text-sm text-slate-500">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <span>{{ formatDate(feedback.created_at) }}</span>
                </div>

                <div :class="getRatingBadgeClass(feedback.rating)" class="px-3 py-1 rounded-full text-xs font-medium">
                  {{ getRatingText(feedback.rating) }}
                </div>
              </div>

              <div class="bg-slate-50 rounded-xl p-4 border-l-4 border-purple-500">
                <div class="flex items-start space-x-3">
                  <svg class="w-5 h-5 text-purple-500 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                    <path
                      d="M4.583 17.321C3.553 16.227 3 15 3 13.011c0-3.5 2.457-6.637 6.03-8.188l.893 1.378c-3.335 1.804-3.987 4.145-4.247 5.621.537-.278 1.24-.375 1.929-.311 1.804.167 3.226 1.648 3.226 3.489a3.5 3.5 0 01-3.5 3.5c-1.073 0-2.099-.49-2.748-1.179zm10 0C13.553 16.227 13 15 13 13.011c0-3.5 2.457-6.637 6.03-8.188l.893 1.378c-3.335 1.804-3.987 4.145-4.247 5.621.537-.278 1.24-.375 1.929-.311 1.804.167 3.226 1.648 3.226 3.489a3.5 3.5 0 01-3.5 3.5c-1.073 0-2.099-.49-2.748-1.179z" />
                  </svg>
                  <p class="text-gray-700 leading-relaxed">
                    {{ feedback.comment }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-else-if="organizerStore.feedbackState.loading" class="text-center py-12">
          <div class="inline-flex items-center space-x-2 text-slate-600">
            <svg class="animate-spin w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Loading feedback...</span>
          </div>
        </div>

        <div v-else class="text-center py-16">
          <div class="w-24 h-24 bg-slate-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-12 h-12 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-slate-800 mb-2">No feedback yet</h3>
          <p class="text-slate-600">Event feedback will appear here once attendees start sharing their thoughts.</p>
        </div>
      </div>
    </div>
  </OrganizerLayout>
</template>

<script lang="ts" setup>
import { computed, onMounted } from 'vue';
import OrganizerLayout from '@/layouts/OrganizerLayout.vue';
import { useOrganizerStore } from '@/stores/organizerStore';
import { Star } from 'lucide-vue-next';
import { formatDate } from '@/utils/format';

const organizerStore = useOrganizerStore()
const feedbacks = computed(() => organizerStore.feedbackState.data)

onMounted(() => {
  organizerStore.getOrganizerEventFeedback()
})

const getRatingBadgeClass = (rating: number) => {
  if (rating >= 4.5) return 'bg-green-100 text-green-800';
  if (rating >= 3.5) return 'bg-purple-100 text-purple-800';
  if (rating >= 2.5) return 'bg-yellow-100 text-yellow-800';
  return 'bg-red-100 text-red-800';
};

const getRatingText = (rating: number) => {
  if (rating >= 4.5) return 'Excellent';
  if (rating >= 3.5) return 'Good';
  if (rating >= 2.5) return 'Average';
  return 'Poor';
};
</script>