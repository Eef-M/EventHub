import { defineStore } from "pinia";
import { ref } from "vue";
import type { User } from "../types/user";
import * as userService from "../services/userService"

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const getCurrentUser = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await userService.getMyProfile()
      user.value = res.data.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch user'
    } finally {
      loading.value = false
    }
  }

  return { user, getCurrentUser, loading, error }
})