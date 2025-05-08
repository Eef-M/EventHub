import axios from "axios"
import { defineStore } from "pinia"


export interface Event {
  id: string
  organizer_id: string
  title: string
  description: string
  location: string
  category: string
  date: string
  time: string
  banner_url: string
}

export const useEventStore = defineStore('event', {
  state: () => ({
    events: [] as Event[],
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async fetchEvents() {
      this.loading = true
      this.error = null

      try {
        const response = await axios.get<{data: Event[]}>('http://localhost:8000/api/v1/events')
        this.events = response.data.data
      } catch (err: any) {
        this.error = err.message || 'Failed to fetch events'
      } finally {
        this.loading = false
      }
    }
  }
})