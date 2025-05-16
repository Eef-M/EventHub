import { defineStore } from "pinia"
import type { DashboardStats } from "../types/organizer"
import { ref } from "vue"
import { getDashboardStats } from "../services/organizerService"

export const useDashboardStore = defineStore('dashboard', () => {
  const stats = ref<DashboardStats | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const getStatsDashboard = async () => {
    loading.value = true
    error.value = null
    try {
      const response = await getDashboardStats()

      if (response?.data) {
        stats.value = response.data
      } else if (response?.data?.data) {
        stats.value = response.data.data
      } else {
        throw new Error('Invalid response format')
      }
    } catch (err: any) {
      console.error('Dashboard error:', err)
      error.value = err.response?.data?.message || 'Failed to fetch dashboard stats'
    } finally {
      loading.value = false
    }
  }

  return {
    stats,
    loading,
    error,
    getStatsDashboard
  }
})