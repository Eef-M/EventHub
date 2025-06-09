import { fetchMyRegistrations, fetchRegisterEvent } from "@/services/eventRegsitrationsService";
import type { MyRegistrationsInterface } from "@/types/eventRegistrations";
import { createAsyncState } from "@/utils/asyncState";
import { defineStore } from "pinia";

export const useEventRegistrationsStore = defineStore('event_registrations', {
  state: () => ({
    myRegState: createAsyncState<MyRegistrationsInterface[]>([]),
    eventRegState: createAsyncState(null)
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
    },

    async registerEvent(id: string, payload: FormData) {
      this.eventRegState.loading = true
      this.eventRegState.error = null
      try {
        await fetchRegisterEvent(id, payload)
      } catch (err: any) {
        this.eventRegState.error = err?.response?.data?.error || 'Failed to register event'
        throw err
      } finally {
        this.eventRegState.loading = false
      }
    }
  }
})