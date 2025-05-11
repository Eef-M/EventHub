<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-100 py-12 px-4">
    <div class="max-w-lg w-full bg-white p-8 rounded-2xl shadow-lg">
      <h1 class="text-3xl font-bold text-slate-600 mb-2 text-center"><span class="text-purple-600">Event</span>Hub</h1>
      <hr class="text-purple-600">
      <h2 class="text-lg text-purple-600 font-medium mb-6 mt-2 text-center">Create an Account</h2>

      <form>
        <div class="mb-4">
          <input type="text" id="username"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Username" />
        </div>

        <div class="mb-4">
          <input type="text" id="first_name"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="First Name" />
        </div>

        <div class="mb-4">
          <input type="text" id="last_name"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Last Name" />
        </div>

        <div class="mb-4">
          <input type="email" id="email"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="E-Mail" />
        </div>

        <div class="mb-4">
          <select id="role"
            class="w-full border rounded-md px-4 py-3 focus:outline-none focus:ring-2 focus:ring-purple-500">
            <option value="">I'm applying as...</option>
            <option value="organizer">Organizer</option>
            <option value="participant">Participant</option>
          </select>
        </div>

        <div class="mb-4">
          <input type="password" id="password"
            class="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="Password" />
        </div>

        <div class="mb-4">
          <input id="confirmPassword" type="password" required
            class="w-full border rounded-md px-4 py-3 focus:outline-none focus:ring-2 focus:ring-purple-500"
            placeholder="Confirm Password" />
        </div>

        <button type="submit"
          class="w-full bg-purple-600 text-white py-3 rounded-md hover:bg-purple-700 transition cursor-pointer">
          Register
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
</script>
