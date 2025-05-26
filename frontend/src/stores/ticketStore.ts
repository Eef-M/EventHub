import { fetchCreateTicket } from "@/services/ticketService";
import { createAsyncState } from "@/utils/asyncState";
import { defineStore } from "pinia";

export const useTicketStore = defineStore('ticket', {
  state: () => ({
    createState: createAsyncState(null),
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
    }
  }
})