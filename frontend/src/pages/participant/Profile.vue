<template>
  <ParticipantLayout>
    <div class="max-w-3xl mx-auto mt-10 p-6 bg-white rounded shadow">
      <h1 class="text-2xl font-bold mb-6">My Profile</h1>
      <div v-if="userStore.userState.data">
        <div class="mb-4">
          <label class="block text-gray-700 font-semibold">Full Name</label>
          <p class="text-gray-900">{{ userStore.userState.data.first_name }} {{ userStore.userState.data.last_name }}
          </p>
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 font-semibold">Email</label>
          <p class="text-gray-900">{{ userStore.userState.data.email }}</p>
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 font-semibold">Role</label>
          <p class="text-gray-900 capitalize">{{ userStore.userState.data.role }}</p>
        </div>
      </div>
      <div v-else>
        <p>Loading profile...</p>
      </div>
    </div>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useUserStore } from '../../stores/userStore';
import ParticipantLayout from '../../layouts/ParticipantLayout.vue';

const userStore = useUserStore()

onMounted(() => {
  userStore.getMyProfile()
})
</script>
