import axios from "axios"
import type { LoginPayload, RegisterPayload } from "../types/auth"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchRegister(data: RegisterPayload) {
  const response = await axios.post(`${API_BASE_URL}/auth/register`, data, { withCredentials: true })
  return response
}

export async function fetchLogin(data: LoginPayload) {
  const response = await axios.post(`${API_BASE_URL}/auth/login`, data, { withCredentials: true })
  return response
}

export async function fetchLogout() {
  const response = await axios.post(`${API_BASE_URL}/auth/logout`, {}, { withCredentials: true })
  return response
}