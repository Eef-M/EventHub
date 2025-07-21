<template>
  <nav class="border-b bg-background shadow-sm px-6 py-3 flex items-center justify-between">
    <RouterLink to="/">
      <span class="text-purple-600 font-bold text-xl">Event</span><span
        class="text-slate-600 font-bold text-xl">Hub</span>
    </RouterLink>
    <DropdownMenu>
      <DropdownMenuTrigger as-child>
        <Button variant="ghost" size="icon">
          <Avatar>
            <AvatarImage :src="user?.avatar_url || ''" :alt="user?.username || 'User'" />
            <AvatarFallback class="text-sm font-medium bg-purple-100 text-purple-700">
              {{ getUserInitials(user) }}
            </AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel class="font-normal">
          <div class="flex flex-col space-y-1">
            <p class="text-sm font-medium leading-none">{{ user?.username || 'User' }}</p>
            <p class="text-xs leading-none text-muted-foreground">{{ user?.email }}</p>
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <RouterLink to="/profile">
          <DropdownMenuItem>Profile</DropdownMenuItem>
        </RouterLink>
        <DropdownMenuSeparator />
        <DropdownMenuItem @click="handleLogout">Logout</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  </nav>
</template>

<script lang="ts" setup>
import { computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { useAuthStore } from '@/stores/authStore'
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

const router = useRouter()
const userStore = useUserStore()
const authStore = useAuthStore()

const user = computed(() => userStore.userState.data)

onMounted(() => {
  userStore.getMyProfile()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
  window.location.reload()
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
