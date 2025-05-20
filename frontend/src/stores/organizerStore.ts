import { defineStore } from "pinia"
import type { DashboardStats } from "@/types/organizer"
import type { Event } from "@/types/event"
import { fetchDashboardStats, fetchMyEvents } from "@/services/organizerService"
import { createAsyncState } from "@/utils/asyncState"

export const useOrganizerStore = defineStore('organizer', {
  state: () => ({
    statsState: createAsyncState<DashboardStats | null>(null),
    eventsState: createAsyncState<Event[]>([]),
  }),

  actions: {
    async getDashboardStats() {
      this.statsState.loading = true
      this.statsState.error = null
      try {
        const data = await fetchDashboardStats()
        this.statsState.data = data
      } catch (err: any) {
        this.statsState.error = err?.response?.data?.error || 'Failed to get dashboard stats'
      } finally {
        this.statsState.loading = false
      }
    },

    async getMyEvents() {
      this.eventsState.loading = true
      this.eventsState.error = null
      try {
        const data = await fetchMyEvents()
        this.eventsState.data = data
      } catch (err: any) {
        this.eventsState.error = err?.response?.data?.error || 'Failed to get Events'
      } finally {
        this.eventsState.loading = false
      }
    }
  }
})