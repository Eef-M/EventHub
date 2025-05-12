import axios from "axios"
import type { LoginPayload, RegisterPayload } from "../types/auth"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export const register = (data: RegisterPayload) => {
  return axios.post(`${API_BASE_URL}/auth/register`, data, { withCredentials: true })
}

export const login = (data: LoginPayload) => {
  return axios.post(`${API_BASE_URL}/auth/login`, data, { withCredentials: true })
}

export const logout = () => {
  return axios.post(`${API_BASE_URL}/auth/logout`, {}, { withCredentials: true })
}