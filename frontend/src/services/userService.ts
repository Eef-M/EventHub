import type { UserInterface } from "@/types/user"
import axios from "axios"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true,
  timeout: 10000,
})

apiClient.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

apiClient.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    console.warn(error)
    return Promise.reject(error)
  }
)

export async function fetchMyProfile(): Promise<UserInterface> {
  try {
    const response = await apiClient.get('/user/me')
    return response.data.data
  } catch (error: any) {
    throw error
  }
}

export async function fetchUpdateProfile(payload: FormData): Promise<UserInterface> {
  try {
    const response = await apiClient.put('/user/update', payload, {
      headers: { 'Content-Type': 'multipart/form-data' },
      withCredentials: true
    })
    return response.data.data
  } catch (error: any) {
    throw error
  }
}