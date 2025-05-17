import { defineStore } from "pinia"
import * as authService from "../services/authService"
import { ref } from "vue"
import type { LoginPayload, RegisterPayload } from "../types/auth"
import { useUserStore } from "./userStore"

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const userStore = useUserStore()

  const login = async (payload: LoginPayload) => {
    loading.value = true
    error.value = null
    try {
      await authService.login(payload)
      isAuthenticated.value = true
      await userStore.getMyProfile()
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Login failed'
    } finally {
      loading.value = false
    }
  }

  const register = async (payload: RegisterPayload) => {
    loading.value = true
    error.value = null
    try {
      await authService.register(payload)
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Register failed'
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    try {
      await authService.logout()
      isAuthenticated.value = false
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Logout failed'
    }
  }

  return { login, register, logout, isAuthenticated, loading, error }
})