import type { Event } from "@/types/event"
import type { DashboardStats } from "@/types/organizer"
import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchDashboardStats(): Promise<DashboardStats> {
  const response = await axios.get(`${API_BASE_URL}/organizer/dashboard`, { withCredentials: true })
  return response.data
}

export async function fetchMyEvents(): Promise<Event[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/events`, { withCredentials: true })
  return response.data.data
}