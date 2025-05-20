import { defineStore } from "pinia"
import type { Event } from "@/types/event"
import { fetchCreateEvent, fetchEventById, fetchEvents } from "@/services/eventService"
import { fetchMyEvents } from "@/services/organizerService"

export const useEventStore = defineStore('event', {
  state: () => ({
    events: [] as Event[],
    event: null as Event | null,
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async getEvents() {
      this.loading = true
      this.error = null

      try {
        const data = await fetchEvents()
        this.events = data
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to get events'
        throw err
      } finally {
        this.loading = false
      }
    },

    async getEventById(id: string) {
      this.loading = true
      this.error = null
      try {
        const data = await fetchEventById(id)
        this.event = data
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to get event'
      } finally {
        this.loading = false
      }
    },

    async createEvent(payload: FormData) {
      this.loading = true
      this.error = null
      try {
        await fetchCreateEvent(payload)
        await fetchMyEvents()
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to create event'
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})