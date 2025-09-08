import { defineStore } from "pinia";
import type { UserInterface } from "@/types/user";
import { fetchMyProfile, fetchUpdateProfile } from "@/services/userService";
import { createAsyncState } from "@/utils/asyncState";

export const useUserStore = defineStore('user', {
  state: () => ({
    userState: createAsyncState<UserInterface | null>(null),
    userUpdateState: createAsyncState(null),
    isAuthenticated: false
  }),

  getters: {
    user: (state) => state.userState.data,
    isLoading: (state) => state.userState.loading,
    hasError: (state) => state.userState.error,
    isLoggedIn: (state) => state.isAuthenticated && !!state.userState.data,
    updateUser: (state) => state.userUpdateState.data,
    updateLoading: (state) => state.userUpdateState.loading,
    updateError: (state) => state.userUpdateState.error,
    updateLoggedIn: (state) => state.isAuthenticated && !!state.userUpdateState.data
  },

  actions: {
    async getMyProfile() {
      if (this.userState.loading) return

      this.userState.loading = true
      this.userState.error = null

      try {
        const data = await fetchMyProfile()
        this.userState.data = data
        this.isAuthenticated = true
      } catch (err: any) {
        this.userState.error = 'Failed to get profile'
        throw err
      } finally {
        this.userState.loading = false
      }
    },

    async updateMyProfile(payload: FormData) {
      if (this.userUpdateState.loading) return

      this.userUpdateState.loading = true
      this.userState.error = null

      try {
        await fetchUpdateProfile(payload)
        this.isAuthenticated = true
      } catch (err: any) {
        this.userUpdateState.error = 'Failed to update profile'
        throw err
      } finally {
        this.userUpdateState.loading = false
      }
    },

    // Method to silently check auth status
    async checkAuthStatus() {
      try {
        await this.getMyProfile()
        return true
      } catch (error) {
        // Silently fail for auth checks
        return false
      }
    },

    // Clear user data on logout
    clearUserData() {
      this.userState.data = null
      this.userState.error = null
      this.userState.loading = false
      this.isAuthenticated = false
    }
  }
})