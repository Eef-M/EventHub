import { fetchCreateTicket, fetchDeleteTicket, fetchUpdateTicket } from "@/services/ticketService";
import { createAsyncState } from "@/utils/asyncState";
import { defineStore } from "pinia";

export const useTicketStore = defineStore('ticket', {
  state: () => ({
    createState: createAsyncState(null),
    updateState: createAsyncState(null),
    deleteState: createAsyncState(null),
  }),

  actions: {
    async createTicket(payload: FormData) {
      this.createState.loading = true
      this.createState.error = null
      try {
        await fetchCreateTicket(payload)
      } catch (err: any) {
        this.createState.error = err?.response?.data?.error || 'Failed to create ticket'
        throw err
      } finally {
        this.createState.loading = false
      }
    },

    async updateTicket(id: string, payload: FormData) {
      this.updateState.loading = true
      this.updateState.error = null
      try {
        await fetchUpdateTicket(id, payload)
      } catch (err: any) {
        this.updateState.error = err?.response?.data?.error || 'Failed to update ticket'
        throw err
      } finally {
        this.updateState.loading = false
      }
    },

    async deleteTicket(id: string) {
      this.deleteState.loading = true
      this.deleteState.error = null
      try {
        await fetchDeleteTicket(id)
      } catch (err: any) {
        this.deleteState.error = err?.response?.data?.error || 'Failed to delete ticket'
        throw err
      } finally {
        this.deleteState.loading = false
      }
    }
  }
})