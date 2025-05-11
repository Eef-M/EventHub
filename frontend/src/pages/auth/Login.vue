<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white shadow-md rounded-lg p-8 w-full max-w-lg">
      <h1 class="text-3xl font-bold text-slate-600 mb-2 text-center"><span class="text-purple-600">Event</span>Hub</h1>
      <hr class="text-purple-600">
      <h2 class="text-lg text-purple-600 font-medium mb-6 mt-2 text-center">Login</h2>

      <form>
        <div class="mb-8">
          <input type="email" id="email"
            class="w-full px-4 py-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="E-Mail" />
        </div>

        <div class="mb-8">
          <input type="password" id="password"
            class="w-full px-4 py-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Password" />
        </div>

        <button type="submit"
          class="w-full bg-purple-600 text-white py-4 rounded-md hover:bg-purple-700 transition cursor-pointer">
          Login
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
</script>