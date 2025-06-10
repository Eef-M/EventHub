import { fetchFeedbacks } from "@/services/feedbackService";
import type { FeedbackInterface } from "@/types/feedback";
import { createAsyncState } from "@/utils/asyncState";
import { defineStore } from "pinia";

export const useFeedbackStore = defineStore('feedback', {
  state: () => ({
    feedbacksState: createAsyncState<FeedbackInterface[]>([])
  }),

  actions: {
    async getFeedbacks(id: string) {
      this.feedbacksState.loading = true
      this.feedbacksState.error = null
      try {
        const data = await fetchFeedbacks(id)
        this.feedbacksState.data = data
      } catch (err: any) {
        this.feedbacksState.error = err?.response?.data?.error || 'Failed to get Feedbacks'
      } finally {
        this.feedbacksState.loading = false
      }
    }
  }
})