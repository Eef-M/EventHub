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
            <AvatarImage src="/avatar.png" />
            <AvatarFallback>A</AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
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
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { useAuthStore } from '@/stores/authStroe'
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

onMounted(() => {
  userStore.getCurrentUser()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
  window.location.reload()
}
</script>
