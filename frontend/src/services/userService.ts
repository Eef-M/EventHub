import type { UserInterface } from "@/types/user"
import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchMyProfile(): Promise<UserInterface> {
  const response = await axios.get(`${API_BASE_URL}/user/me`, { withCredentials: true })
  return response.data.data
}