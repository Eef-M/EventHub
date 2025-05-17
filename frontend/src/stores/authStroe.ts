import { defineStore } from "pinia"
import type { LoginPayload, RegisterPayload } from "../types/auth"
import { fetchLogin, fetchLogout, fetchRegister } from "@/services/authService"
import { useUserStore } from "@/stores/userStore"

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    loading: false,
    error: null as string | null,
    userStore: useUserStore()
  }),

  actions: {
    async register(payload: RegisterPayload) {
      this.loading = true
      this.error = null
      try {
        await fetchRegister(payload)
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Register failed'
      } finally {
        this.loading = false
      }
    },

    async login(payload: LoginPayload) {
      this.loading = true
      this.error = null
      try {
        await fetchLogin(payload)
        this.isAuthenticated = true
        await this.userStore.getMyProfile()
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Login failed'
      } finally {
        this.loading = false
      }
    },

    async logout() {
      this.loading = true
      this.error = null
      try {
        await fetchLogout()
        this.isAuthenticated = false
      } catch (err: any) {
        this.error = err?.response?.data?.error || 'Logout failed'
      } finally {
        this.loading = false
      }
    }
  }
})