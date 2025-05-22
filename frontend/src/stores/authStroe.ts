import { defineStore } from "pinia"
import type { LoginPayload, RegisterPayload } from "@/types/auth"
import { fetchLogin, fetchLogout, fetchRegister } from "@/services/authService"
import { useUserStore } from "@/stores/userStore"
import { createAsyncState } from "@/utils/asyncState"

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    registerState: createAsyncState(null),
    loginState: createAsyncState(null),
    logoutState: createAsyncState(null),
    userStore: useUserStore()
  }),

  actions: {
    async register(payload: RegisterPayload) {
      this.registerState.loading = true
      this.registerState.error = null
      try {
        await fetchRegister(payload)
      } catch (err: any) {
        this.registerState.error = err?.response?.data?.error || 'Register failed'
        throw err
      } finally {
        this.registerState.loading = false
      }
    },

    async login(payload: LoginPayload) {
      this.loginState.loading = true
      this.loginState.error = null
      try {
        await fetchLogin(payload)
        this.isAuthenticated = true
        await this.userStore.getMyProfile()
      } catch (err: any) {
        this.loginState.error = err?.response?.data?.error || 'Login failed'
        throw err
      } finally {
        this.loginState.loading = false
      }
    },

    async logout() {
      this.logoutState.loading = true
      this.logoutState.error = null
      try {
        await fetchLogout()
        this.isAuthenticated = false
      } catch (err: any) {
        this.logoutState.error = err?.response?.data?.error || 'Logout failed'
        throw err
      } finally {
        this.logoutState.loading = false
      }
    }
  }
})