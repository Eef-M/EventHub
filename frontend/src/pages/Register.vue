<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100 py-12 px-4">
    <div class="max-w-md w-full bg-white p-8 rounded-2xl shadow-lg">
      <h2 class="text-2xl font-bold text-indigo-800 mb-6 text-center">Create an Account</h2>

      <form>
        <div class="mb-3">
          <label class="block text-gray-700 mb-2" for="username">Username</label>
          <input type="text" id="username"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400"
            placeholder="john123" />
        </div>

        <div class="mb-3">
          <label class="block text-gray-700 mb-2" for="first_name">First Name</label>
          <input type="text" id="first_name"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400"
            placeholder="John" />
        </div>

        <div class="mb-3">
          <label class="block text-gray-700 mb-2" for="last_name">Last Name</label>
          <input type="text" id="last_name"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400"
            placeholder="Doe" />
        </div>

        <div class="mb-3">
          <label class="block text-gray-700 mb-2" for="email">Email</label>
          <input type="email" id="email"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400"
            placeholder="email@example.com" />
        </div>

        <div class="mb-4">
          <label class="block text-gray-700 mb-2" for="password">Password</label>
          <input type="password" id="password"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400"
            placeholder="********" />
        </div>

        <button type="submit" class="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-blue-700 transition">
          Register
        </button>
      </form>

      <p class="mt-6 text-sm text-center text-gray-600">
        Already have an account?
        <RouterLink to="/login" class="text-indigo-600 hover:underline">Login</RouterLink>
      </p>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>
