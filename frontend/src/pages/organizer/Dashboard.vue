<template>
  <OrganizerLayout>
    <section class="p-6">
      <h1 class="text-2xl font-semibold mb-4">Dashboard</h1>

      <div v-if="dashboardStore.loading" class="text-center py-8">
        <p class="text-gray-600">Loading dashboard data...</p>
      </div>

      <div v-else-if="dashboardStore.error" class="bg-red-50 border border-red-200 rounded-xl p-4 mb-6">
        <p class="text-red-600">{{ dashboardStore.error }}</p>
        <button @click="refreshData" class="mt-2 px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700">
          Try Again
        </button>
      </div>

      <template v-else>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
          <div class="bg-white shadow rounded-xl p-4">
            <h2 class="text-lg font-medium text-gray-700">Total Events</h2>
            <p class="text-3xl font-bold text-purple-600">{{ dashboardStore.stats?.total_events ?? 0 }}</p>
          </div>
          <div class="bg-white shadow rounded-xl p-4">
            <h2 class="text-lg font-medium text-gray-700">Total Registrations</h2>
            <p class="text-3xl font-bold text-purple-600">{{ dashboardStore.stats?.total_registrations ?? 0 }}</p>
          </div>
          <div class="bg-white shadow rounded-xl p-4">
            <h2 class="text-lg font-medium text-gray-700">Pending Approvals</h2>
            <p class="text-3xl font-bold text-purple-600">{{ dashboardStore.stats?.pending_approvals ?? 0 }}</p>
          </div>
          <div class="bg-white shadow rounded-xl p-4">
            <h2 class="text-lg font-medium text-gray-700">Feedback Received</h2>
            <p class="text-3xl font-bold text-purple-600">{{ dashboardStore.stats?.feedback_received ?? 0 }}</p>
          </div>
        </div>

        <div class="bg-white shadow rounded-xl p-6">
          <h2 class="text-xl font-semibold mb-4">Recent Registrations</h2>
          <ul class="divide-y divide-gray-100">
            <li v-if="!dashboardStore.stats?.recent_registrations?.length" class="py-2 text-gray-500">
              No recent registrations found
            </li>
            <li class="py-2 text-slate-500" v-for="(r, index) in dashboardStore.stats?.recent_registrations"
              :key="index">
              {{ r.user_name }} - {{ r.event_title }}
            </li>
          </ul>
        </div>
      </template>
    </section>
  </OrganizerLayout>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import OrganizerLayout from '../../layouts/OrganizerLayout.vue';
import { useDashboardStore } from '../../stores/dashboardStore';

const dashboardStore = useDashboardStore()

const refreshData = () => {
  dashboardStore.getStatsDashboard()
}

onMounted(() => {
  refreshData()
})
</script>