import axios from "axios"
import type { LoginInterface, RegisterInterface } from "@/types/auth"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchRegister(data: RegisterInterface) {
  const response = await axios.post(`${API_BASE_URL}/auth/register`, data, { withCredentials: true })
  return response
}

export async function fetchLogin(data: LoginInterface) {
  const response = await axios.post(`${API_BASE_URL}/auth/login`, data, { withCredentials: true })
  return response
}

export async function fetchLogout() {
  const response = await axios.post(`${API_BASE_URL}/auth/logout`, {}, { withCredentials: true })
  return response
}