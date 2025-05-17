import { defineStore } from "pinia"
import type { DashboardStats } from "@/types/organizer"
import type { Event } from "@/types/event"
import { fetchDashboardStats, fetchMyEvents } from "@/services/organizerService"

export const useOrganizerStore = defineStore('organizer', {
  state: () => ({
    stats: null as DashboardStats | null,
    events: [] as Event[],
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async getDashboardStats() {
      this.loading = true
      this.error = null
      try {
        const data = await fetchDashboardStats()
        this.stats = data
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to get dashboard stats'
      } finally {
        this.loading = false
      }
    },

    async getMyEvents() {
      this.loading = true
      this.error = null
      try {
        const data = await fetchMyEvents()
        this.events = data
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to get Events'
      } finally {
        this.loading = false
      }
    }
  }
})