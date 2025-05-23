import type { Event } from "@/types/event"
import type { DashboardStats, OrganizerEventFeedback, OrganizerEventRegistration } from "@/types/organizer"
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

export async function fetchOrganizerEventRegistrations(): Promise<OrganizerEventRegistration[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/registrations`, { withCredentials: true })
  return response.data.data
}

export async function fetchOrganizerEventFeedback(): Promise<OrganizerEventFeedback[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/feedbacks`, { withCredentials: true })
  return response.data.data
}