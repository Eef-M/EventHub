import axios from "axios"
import type { LoginPayload, RegisterPayload } from "../types/auth"

const API = 'http://localhost:8000/api/v1'

export const register = (data: RegisterPayload) => {
  return axios.post(`${API}/auth/register`, data, { withCredentials: true })
}

export const login = (data: LoginPayload) => {
  return axios.post(`${API}/auth/login`, data, { withCredentials: true })
}

export const getMe = () => {
  return axios.get('http://localhost:8000/api/v1/me', { withCredentials: true })
} 