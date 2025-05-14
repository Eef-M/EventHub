import { defineStore } from "pinia"
import type { Event } from "../types/event"
import { getEvents as getEventsAPI } from "../services/eventService"

export const useEventStore = defineStore('event', {
  state: () => ({
    events: [] as Event[],
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async getEvents() {
      this.loading = true
      this.error = null

      try {
        const data = await getEventsAPI()
        this.events = data
      } catch (err: any) {
        this.error = err.message || 'Failed to fetch events'
      } finally {
        this.loading = false
      }
    }
  }
})