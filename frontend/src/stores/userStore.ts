import { defineStore } from "pinia";
import type { User } from "@/types/user";
import { fetchMyProfile } from "@/services/userService";
import { createAsyncState } from "@/utils/asyncState";

export const useUserStore = defineStore('user', {
  state: () => ({
    userState: createAsyncState<User | null>(null)
  }),

  actions: {
    async getMyProfile() {
      this.userState.loading = true
      this.userState.error = null
      try {
        const data = await fetchMyProfile()
        this.userState.data = data
      } catch (err: any) {
        this.userState.error = err?.response?.data?.error || 'Failed to get profile data'
      } finally {
        this.userState.loading = false
      }
    }
  }
})