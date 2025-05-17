import { defineStore } from "pinia";
import type { User } from "../types/user";
import { fetchMyProfile } from "@/services/userService";

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as User | null,
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async getMyProfile() {
      this.loading = true
      this.error = null
      try {
        const data = await fetchMyProfile()
        this.user = data
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Failed to get profile data'
      } finally {
        this.loading = false
      }
    }
  }
})