import { defineStore } from "pinia"
import * as authService from "../services/authService"
import { ref } from "vue"
import type { LoginPayload, RegisterPayload } from "../types/auth"
import type { User } from "../types/user"

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const user = ref<User | null>(null)

  const login = async (payload: LoginPayload) => {
    loading.value = true
    error.value = null
    try {
      await authService.login(payload)
      isAuthenticated.value = true
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

  const getUser = async () => {
    try {
      const res = await authService.getMe()
      user.value = res.data
    } catch (err) {
      console.error('Failed to fetch user')
    }
  }

  return { login, register, getUser, isAuthenticated, loading, error }
})