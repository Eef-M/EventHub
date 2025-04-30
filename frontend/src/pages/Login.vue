<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white shadow-md rounded-lg p-8 w-full max-w-md">
      <h2 class="text-3xl font-semibold text-center text-indigo-800 mb-6">Login</h2>

      <form>
        <div class="mb-4">
          <label class="block text-gray-700 mb-2" for="email">Email</label>
          <input type="email" id="email"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="email@example.com" />
        </div>

        <div class="mb-6">
          <label class="block text-gray-700 mb-2" for="password">Password</label>
          <input type="password" id="password"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="********" />
        </div>

        <button type="submit" class="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-blue-700 transition">
          Login
        </button>
      </form>

      <p class="mt-6 text-sm text-center text-gray-600">
        Don't have an account yet?
        <RouterLink to="/register" class="text-indigo-600 hover:underline">Register</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>