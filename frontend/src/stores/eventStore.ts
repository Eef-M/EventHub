import { defineStore } from "pinia"
import type { Event } from "@/types/event"
import { fetchCreateEvent, fetchDeleteEvent, fetchEventById, fetchEvents, fetchUpdateEvent } from "@/services/eventService"
import { createAsyncState } from "@/utils/asyncState"

export const useEventStore = defineStore('event', {
  state: () => ({
    eventsState: createAsyncState<Event[]>([]),
    singleEventState: createAsyncState<Event | null>(null),
    createState: createAsyncState(null),
    updateState: createAsyncState(null),
    deleteState: createAsyncState(null),
  }),
  actions: {
    async getEvents() {
      this.eventsState.loading = true
      this.eventsState.error = null
      try {
        const data = await fetchEvents()
        this.eventsState.data = data
      } catch (err: any) {
        this.eventsState.error = err?.response?.data?.error || 'Failed to get events'
        throw err
      } finally {
        this.eventsState.loading = false
      }
    },

    async getEventById(id: string) {
      this.singleEventState.loading = true
      this.singleEventState.error = null
      try {
        const data = await fetchEventById(id)
        this.singleEventState.data = data
      } catch (err: any) {
        this.singleEventState.error = err?.response?.data?.error || 'Failed to get event'
        throw err
      } finally {
        this.singleEventState.loading = false
      }
    },

    async createEvent(payload: FormData) {
      this.createState.loading = true
      this.createState.error = null
      try {
        await fetchCreateEvent(payload)
      } catch (err: any) {
        this.createState.error = err?.response?.data?.error || 'Failed to create event'
        throw err
      } finally {
        this.createState.loading = false
      }
    },

    async updateEvent(id: string, payload: FormData) {
      this.updateState.loading = true
      this.updateState.error = null
      try {
        await fetchUpdateEvent(id, payload)
      } catch (err: any) {
        this.updateState.error = err?.response?.data?.error || 'Failed to update event'
        throw err
      } finally {
        this.updateState.loading = false
      }
    },

    async deleteEvent(id: string) {
      this.deleteState.loading = true
      this.deleteState.error = null
      try {
        await fetchDeleteEvent(id)
      } catch (err: any) {
        this.deleteState.error = err?.response?.data?.error || 'Failed to delete event'
        throw err
      } finally {
        this.deleteState.loading = false
      }
    }
  }
})