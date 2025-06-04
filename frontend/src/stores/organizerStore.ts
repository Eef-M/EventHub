import { defineStore } from "pinia"
import type { DashboardStatsInterface, OrganizerEventFeedbackInterface, OrganizerEventRegistrationInterface, OrganizerTicketInterface } from "@/types/organizer"
import type { EventInterface } from "@/types/event"
import { fetchDashboardStats, fetchMyEvents, fetchOrganizerEventFeedback, fetchOrganizerEventRegistrations, fetchOrganizerTickets } from "@/services/organizerService"
import { createAsyncState } from "@/utils/asyncState"

export const useOrganizerStore = defineStore('organizer', {
  state: () => ({
    statsState: createAsyncState<DashboardStatsInterface | null>(null),
    eventsState: createAsyncState<EventInterface[]>([]),
    ticketsState: createAsyncState<OrganizerTicketInterface[]>([]),
    registrationsState: createAsyncState<OrganizerEventRegistrationInterface[]>([]),
    feedbackState: createAsyncState<OrganizerEventFeedbackInterface[]>([])
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
        throw err
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
        throw err
      } finally {
        this.eventsState.loading = false
      }
    },

    async getOrganizerEventRegistrations() {
      this.registrationsState.loading = true
      this.registrationsState.error = null
      try {
        const data = await fetchOrganizerEventRegistrations()
        this.registrationsState.data = data
      } catch (err: any) {
        this.registrationsState.error = err?.response?.data?.error || 'Failed to get Event Registrations'
        throw err
      } finally {
        this.registrationsState.loading = false
      }
    },

    async getOrganizerEventFeedback() {
      this.feedbackState.loading = true
      this.feedbackState.error = null
      try {
        const data = await fetchOrganizerEventFeedback()
        this.feedbackState.data = data
      } catch (err: any) {
        this.feedbackState.error = err?.response?.data?.error || 'Failed to get Event Registrations'
        throw err
      } finally {
        this.feedbackState.loading = false
      }
    },

    async getOrganizerTickets() {
      this.ticketsState.loading = true
      this.ticketsState.error = null
      try {
        const data = await fetchOrganizerTickets()
        this.ticketsState.data = data
      } catch (err: any) {
        this.ticketsState.error = err?.response?.data?.error || 'Failed to get Tickets'
        throw err
      } finally {
        this.ticketsState.loading = false
      }
    }
  }
})