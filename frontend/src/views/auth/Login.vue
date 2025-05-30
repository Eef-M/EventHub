<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white shadow-md rounded-lg p-8 w-full max-w-lg">
      <h1 class="text-3xl font-bold text-slate-600 mb-2 text-center">
        <RouterLink to="/">
          <span class="text-purple-600">Event</span>Hub
        </RouterLink>
      </h1>
      <hr class="text-purple-600">
      <h2 class="text-lg text-purple-600 font-medium mb-6 mt-2 text-center">Login</h2>

      <form @submit.prevent="handleSubmit">
        <div class="mb-8">
          <input type="email" id="email"
            class="w-full px-4 py-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="E-Mail" v-model="form.email" />
        </div>

        <div class="mb-8">
          <input type="password" id="password"
            class="w-full px-4 py-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Password" v-model="form.password" />
        </div>

        <p v-if="authStore.loginState.error" class="text-red-500 text-sm mb-4">{{ authStore.loginState.error }}</p>

        <button type="submit" :disabled="authStore.loginState.loading"
          class="w-full bg-purple-600 text-white py-4 rounded-md hover:bg-purple-700 transition cursor-pointer">
          {{ authStore.loginState.loading ? 'Logging in...' : 'Login' }}
        </button>
      </form>

      <p class="mt-6 text-sm text-center text-gray-600">
        Don't have an account yet?
        <RouterLink to="/register" class="text-purple-600 hover:underline">Register</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../stores/authStroe'
import { reactive } from 'vue'

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
  email: '',
  password: ''
})

const handleSubmit = async () => {
  if (!form.email || !form.password) {
    authStore.loginState.error = 'field required'
    return
  }

  await authStore.login(form)
  if (!authStore.loginState.error) {
    router.push('/')
  }
}
</script>