<template>
  <footer class="bg-zinc-100 text-gray-700 mt-16">
    <div class="max-w-7xl mx-auto px-4 py-10 grid grid-cols-1 md:grid-cols-3 gap-8">
      <div>
        <h2 class="text-xl font-bold text-purple-700">EventHub</h2>
        <p class="mt-2 text-sm text-gray-600">
          Your go-to platform for discovering and managing events easily.
        </p>
      </div>

      <div>
        <h3 class="text-md font-semibold mb-2">Quick Links</h3>
        <ul class="space-y-1 text-sm">
          <li>
            <RouterLink to="/" class="hover:underline">Home</RouterLink>
          </li>
          <li>
            <RouterLink to="/events" class="hover:underline">Browse Events</RouterLink>
          </li>
          <li>
            <RouterLink to="/login" class="hover:underline">Login</RouterLink>
          </li>
          <li>
            <RouterLink to="/register" class="hover:underline">Register</RouterLink>
          </li>
        </ul>
      </div>

      <div>
        <h3 class="text-md font-semibold mb-2">Contact</h3>
        <ul class="text-sm space-y-1">
          <li>Email: support@eventhub.com</li>
          <li>Phone: +62 812 3456 7890</li>
          <li>Location: Jakarta, Indonesia</li>
        </ul>
      </div>
    </div>

    <div class="text-center text-xs text-gray-500 py-4 border-t border-gray-200">
      © 2025 EventHub. All rights reserved.
    </div>
  </footer>
</template>

<script setup>
</script>
