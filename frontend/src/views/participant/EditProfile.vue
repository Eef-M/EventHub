<template>
  <ParticipantLayout>
    <div class="min-h-screen bg-white">
      <div class="bg-slate-50 border-b border-slate-200">
        <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div class="flex items-center gap-4 mb-6">
            <Button variant="ghost" @click="$router.back()" class="p-2">
              <ArrowLeft class="w-5 h-5" />
            </Button>
            <h1 class="text-3xl font-bold text-gray-900">Edit Profile</h1>
          </div>
        </div>
      </div>

      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <form @submit.prevent="handleSubmit" class="space-y-8">
          <div class="bg-white rounded-lg border border-gray-200 p-6">
            <h2 class="text-xl font-semibold text-gray-900 mb-6">Profile Photo</h2>
            <div class="flex items-center gap-6">
              <div class="relative">
                <img :src="avatarPreview || form.avatar_url" :alt="form.username + ' avatar'"
                  class="w-24 h-24 rounded-full object-cover border-4 border-white shadow-lg">
                <Button type="button" variant="outline" size="icon" @click="triggerFileInput"
                  class="absolute -bottom-1 -right-1 w-8 h-8 rounded-full bg-white border-2 border-gray-300 hover:bg-gray-50">
                  <Camera class="w-4 h-4" />
                </Button>
              </div>
              <div class="flex-1">
                <p class="text-sm text-gray-600 mb-2">Upload a new profile photo</p>
                <div class="flex gap-3">
                  <Button type="button" variant="outline" @click="triggerFileInput">
                    <Upload class="w-4 h-4 mr-2" />
                    Change Photo
                  </Button>
                  <Button type="button" variant="ghost" @click="removeAvatar" v-if="avatarPreview">
                    Remove
                  </Button>
                </div>
                <input ref="fileInput" type="file" accept="image/*" @change="handleFileChange" class="hidden">
                <p class="text-xs text-gray-500 mt-2">Supported formats: JPG, PNG, GIF. Max size: 5MB</p>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-lg border border-gray-200 p-6">
            <h2 class="text-xl font-semibold text-gray-900 mb-6">Basic Information</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label for="first_name" class="block text-sm font-medium text-gray-700 mb-2">
                  First Name *
                </label>
                <input id="first_name" v-model="form.first_name" type="text" required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                  :class="{ 'border-red-300': errors.first_name }">
                <p v-if="errors.first_name" class="text-red-500 text-xs mt-1">{{ errors.first_name }}</p>
              </div>

              <div>
                <label for="last_name" class="block text-sm font-medium text-gray-700 mb-2">
                  Last Name *
                </label>
                <input id="last_name" v-model="form.last_name" type="text" required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                  :class="{ 'border-red-300': errors.last_name }">
                <p v-if="errors.last_name" class="text-red-500 text-xs mt-1">{{ errors.last_name }}</p>
              </div>

              <div>
                <label for="username" class="block text-sm font-medium text-gray-700 mb-2">
                  Username *
                </label>
                <input id="username" v-model="form.username" type="text" required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                  :class="{ 'border-red-300': errors.username }">
                <p v-if="errors.username" class="text-red-500 text-xs mt-1">{{ errors.username }}</p>
              </div>

              <div>
                <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
                  Email Address *
                </label>
                <input id="email" v-model="form.email" type="email" required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                  :class="{ 'border-red-300': errors.email }">
                <p v-if="errors.email" class="text-red-500 text-xs mt-1">{{ errors.email }}</p>
              </div>

              <div class="md:col-span-2">
                <label for="role" class="block text-sm font-medium text-gray-700 mb-2">
                  Role
                </label>
                <input id="role" :value="form.role" type="text" disabled
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 cursor-not-allowed">
                <p class="text-xs text-gray-500 mt-1">Role cannot be changed</p>
              </div>
            </div>
          </div>

          <!-- Error Display -->
          <div v-if="userStore.updateError" class="bg-red-50 border border-red-200 rounded-md p-4">
            <div class="flex">
              <AlertCircle class="w-5 h-5 text-red-400" />
              <div class="ml-3">
                <h3 class="text-sm font-medium text-red-800">Update Failed</h3>
                <p class="text-sm text-red-700 mt-1">{{ userStore.updateError }}</p>
              </div>
            </div>
          </div>

          <!-- Success Message -->
          <div v-if="showSuccess" class="bg-green-50 border border-green-200 rounded-md p-4">
            <div class="flex">
              <CheckCircle class="w-5 h-5 text-green-400" />
              <div class="ml-3">
                <h3 class="text-sm font-medium text-green-800">Profile Updated</h3>
                <p class="text-sm text-green-700 mt-1">Your profile has been successfully updated!</p>
              </div>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex justify-end gap-4 pt-6 border-t border-gray-200">
            <Button type="button" variant="outline" @click="$router.back()" :disabled="userStore.updateLoading">
              Cancel
            </Button>
            <Button type="submit" :disabled="userStore.updateLoading || !isFormChanged"
              class="bg-purple-600 hover:bg-purple-700 text-white px-6 py-2">
              <div v-if="userStore.updateLoading" class="flex items-center">
                <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                Updating...
              </div>
              <span v-else>Save Changes</span>
            </Button>
          </div>
        </form>
      </div>
    </div>
  </ParticipantLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, reactive, watch } from 'vue'
import { Button } from '@/components/ui/button'
import { ArrowLeft, Camera, Upload, AlertCircle, CheckCircle } from 'lucide-vue-next'
import ParticipantLayout from '@/layouts/ParticipantLayout.vue'
import { useUserStore } from '@/stores/userStore'
import type { UserInterface } from '@/types/user'

const userStore = useUserStore()
const fileInput = ref<HTMLInputElement>()
const avatarPreview = ref<string | null>(null)
const selectedFile = ref<File | null>(null)
const showSuccess = ref(false)

const form = reactive<UserInterface>({
  id: '',
  username: '',
  first_name: '',
  last_name: '',
  email: '',
  role: '',
  avatar_url: ''
})

// Original form data for comparison
const originalForm = reactive<UserInterface>({
  id: '',
  username: '',
  first_name: '',
  last_name: '',
  email: '',
  role: '',
  avatar_url: ''
})

const errors = reactive({
  first_name: '',
  last_name: '',
  username: '',
  email: ''
})

const isFormChanged = computed(() => {
  return (
    form.first_name !== originalForm.first_name ||
    form.last_name !== originalForm.last_name ||
    form.username !== originalForm.username ||
    form.email !== originalForm.email ||
    selectedFile.value !== null
  )
})

onMounted(async () => {
  if (!userStore.user) {
    await userStore.getMyProfile()
  }

  if (userStore.user) {
    Object.assign(form, userStore.user)
    Object.assign(originalForm, userStore.user)
  }
})

watch(() => userStore.user, (newUser) => {
  if (newUser) {
    Object.assign(form, newUser)
    Object.assign(originalForm, newUser)
  }
}, { deep: true })

const triggerFileInput = () => {
  fileInput.value?.click()
}

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]

  if (file) {
    // Validate file size (5MB limit)
    if (file.size > 5 * 1024 * 1024) {
      alert('File size must be less than 5MB')
      return
    }

    // Validate file type
    if (!file.type.startsWith('image/')) {
      alert('Please select a valid image file')
      return
    }

    selectedFile.value = file

    // Create preview
    const reader = new FileReader()
    reader.onload = (e) => {
      avatarPreview.value = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

const removeAvatar = () => {
  selectedFile.value = null
  avatarPreview.value = null
  if (fileInput.value) {
    fileInput.value.value = ''
  }
}

// Form validation
const validateForm = (): boolean => {
  let isValid = true

  Object.keys(errors).forEach(key => {
    errors[key as keyof typeof errors] = ''
  })

  if (!form.first_name.trim()) {
    errors.first_name = 'First name is required'
    isValid = false
  }

  if (!form.last_name.trim()) {
    errors.last_name = 'Last name is required'
    isValid = false
  }

  if (!form.username.trim()) {
    errors.username = 'Username is required'
    isValid = false
  } else if (form.username.length < 3) {
    errors.username = 'Username must be at least 3 characters'
    isValid = false
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!form.email.trim()) {
    errors.email = 'Email is required'
    isValid = false
  } else if (!emailRegex.test(form.email)) {
    errors.email = 'Please enter a valid email address'
    isValid = false
  }

  return isValid
}

// Form submission
const handleSubmit = async () => {
  if (!validateForm()) {
    return
  }

  try {
    const formData = new FormData()
    formData.append('first_name', form.first_name.trim())
    formData.append('last_name', form.last_name.trim())
    formData.append('username', form.username.trim())
    formData.append('email', form.email.trim())

    if (selectedFile.value) {
      formData.append('avatar_url', selectedFile.value)
    }

    await userStore.updateMyProfile(formData)

    showSuccess.value = true
    setTimeout(() => {
      showSuccess.value = false
    }, 3000)

    selectedFile.value = null
    avatarPreview.value = null

    Object.assign(originalForm, form)

    await userStore.getMyProfile()

  } catch (error) {
    console.error('Failed to update profile:', error)
  }
}
</script>