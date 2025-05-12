<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-100 py-12 px-4">
    <div class="max-w-lg w-full bg-white p-8 rounded-2xl shadow-lg">

      <h1 class="text-3xl font-bold text-slate-600 mb-2 text-center">
        <RouterLink to="/">
          <span class="text-purple-600">Event</span>Hub
        </RouterLink>
      </h1>
      <hr class="text-purple-600">
      <h2 class="text-lg text-purple-600 font-medium mb-6 mt-2 text-center">Create an Account</h2>

      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <input type="text" id="username"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Username" v-model="form.username" />
        </div>

        <div class="mb-4">
          <input type="text" id="first_name"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="First Name" v-model="form.first_name" />
        </div>

        <div class="mb-4">
          <input type="text" id="last_name"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Last Name" v-model="form.last_name" />
        </div>

        <div class="mb-4">
          <input type="email" id="email"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="E-Mail" v-model="form.email" />
        </div>

        <div class="mb-4">
          <select id="role" v-model="form.role"
            class="w-full border rounded-md px-4 py-3 focus:outline-none focus:ring-2 focus:ring-purple-500">
            <option value="">I'm applying as...</option>
            <option value="organizer">Organizer</option>
            <option value="participant">Participant</option>
          </select>
        </div>

        <div class="mb-4">
          <input type="password" id="password"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Password" v-model="form.password" />
        </div>

        <div class="mb-4">
          <input id="confirmPassword" type="password"
            class="w-full border rounded-md px-4 py-3 focus:outline-none focus:ring-2 focus:ring-purple-500"
            placeholder="Confirm Password" v-model="form.confirmPassword" />
        </div>

        <p v-if="authStore.error" class="text-red-500 text-sm mb-4">{{ authStore.error }}</p>

        <button type="submit" :disabled="authStore.loading"
          class="w-full bg-purple-600 text-white py-3 rounded-md hover:bg-purple-700 transition cursor-pointer">
          {{ authStore.loading ? 'Registering...' : 'Register' }}
        </button>
      </form>

      <p class="mt-6 text-sm text-center text-gray-600">
        Already have an account?
        <RouterLink to="/login" class="text-purple-600 hover:underline">Login</RouterLink>
      </p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { useAuthStore } from '../../stores/authStroe';
import { reactive } from 'vue';

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
  username: '',
  first_name: '',
  last_name: '',
  email: '',
  password: '',
  confirmPassword: '',
  role: ''
})

const handleSubmit = async () => {
  if (form.password !== form.confirmPassword) {
    authStore.error = 'Passwords do not match'
    return
  }

  await authStore.register(form)
  if (!authStore.error) {
    router.push('/login')
  }
}

</script>
