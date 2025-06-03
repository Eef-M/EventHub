import { fetchMyRegistrations } from "@/services/eventRegsitrationsService";
import type { MyRegistrations } from "@/types/eventRegistrations";
import { createAsyncState } from "@/utils/asyncState";
import { defineStore } from "pinia";

export const useEventRegistrationsStore = defineStore('event_registrations', {
  state: () => ({
    myRegState: createAsyncState<MyRegistrations[]>([])
  }),

  actions: {
    async getMyRegistrations() {
      this.myRegState.loading = true
      this.myRegState.error = null
      try {
        const data = await fetchMyRegistrations()
        this.myRegState.data = data
      } catch (err: any) {
        this.myRegState.error = err?.response?.data?.error || 'Failed to get My Registrations'
        throw err
      } finally {
        this.myRegState.loading = false
      }
    }
  }
})