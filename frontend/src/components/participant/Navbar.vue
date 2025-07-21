<template>
  <nav class="bg-white shadow-md sticky top-0 z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <div class="flex-shrink-0">
          <RouterLink to="/" class="text-2xl font-bold hover:opacity-80 transition-opacity">
            <span class="text-purple-600">Event</span><span class="text-slate-600">Hub</span>
          </RouterLink>
        </div>

        <div class="flex items-center space-x-6">
          <RouterLink to="/events"
            class="text-gray-700 hover:text-purple-600 font-medium transition-colors duration-200 px-3 py-2 rounded-md hover:bg-gray-50">
            Events
          </RouterLink>

          <div class="flex items-center">
            <div v-if="userStore.isLoggedIn" class="relative">
              <DropdownMenu>
                <DropdownMenuTrigger as-child>
                  <Button variant="ghost" size="sm"
                    class="relative h-10 w-10 rounded-full hover:bg-gray-100 transition-colors">
                    <Avatar class="h-8 w-8">
                      <AvatarImage :src="user?.avatar_url || ''" :alt="user?.username || 'User'" />
                      <AvatarFallback class="text-sm font-medium bg-purple-100 text-purple-700">
                        {{ getUserInitials(user) }}
                      </AvatarFallback>
                    </Avatar>
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end" class="w-56 mt-2">
                  <DropdownMenuLabel class="font-normal">
                    <div class="flex flex-col space-y-1">
                      <p class="text-sm font-medium leading-none">{{ user?.username || 'User' }}</p>
                      <p class="text-xs leading-none text-muted-foreground">{{ user?.email }}</p>
                    </div>
                  </DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <RouterLink to="/profile">
                    <DropdownMenuItem class="cursor-pointer">
                      <span>Profile</span>
                    </DropdownMenuItem>
                  </RouterLink>
                  <div v-if="user?.role === 'organizer'">
                    <RouterLink to="/organizer/dashboard">
                      <DropdownMenuItem class="cursor-pointer">
                        <span>Dashboard</span>
                      </DropdownMenuItem>
                    </RouterLink>
                  </div>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem @click="handleLogout" class="cursor-pointer text-red-600 focus:text-red-600">
                    <span>Logout</span>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>

            <div v-else>
              <RouterLink to="/login"
                class="bg-purple-600 text-white px-4 py-2 rounded-md font-medium hover:bg-purple-700 transition-colors duration-200 shadow-sm">
                Login
              </RouterLink>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { computed, onMounted } from 'vue';
import { useUserStore } from '@/stores/userStore';
import { useAuthStore } from '@/stores/authStore';
import router from '@/router';
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'

const userStore = useUserStore()
const authStore = useAuthStore()

const user = computed(() => userStore.user)

onMounted(async () => {
  if (!userStore.isLoggedIn) {
    await userStore.checkAuthStatus()
  }
})

const handleLogout = async () => {
  try {
    await authStore.logout()
    userStore.clearUserData()
    router.push('/')
  } catch (error) {
    console.error('Logout failed:', error)
  }
}

// Helper function to get user initials
const getUserInitials = (user: any) => {
  if (!user) return 'U'

  if (user.name) {
    const names = user.name.split(' ')
    if (names.length >= 2) {
      return `${names[0][0]}${names[1][0]}`.toUpperCase()
    }
    return names[0][0].toUpperCase()
  }

  if (user.email) {
    return user.email[0].toUpperCase()
  }

  return 'U'
}
</script>