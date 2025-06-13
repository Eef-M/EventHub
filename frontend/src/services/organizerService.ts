import type { EventInterface } from "@/types/event"
import type { DashboardStatsInterface, OrganizerEventFeedbackInterface, OrganizerEventRegistrationInterface, OrganizerTicketInterface } from "@/types/organizer"
import axios from "axios"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchDashboardStats(): Promise<DashboardStatsInterface> {
  const response = await axios.get(`${API_BASE_URL}/organizer/dashboard`, { withCredentials: true })
  return response.data
}

export async function fetchMyEvents(): Promise<EventInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/events`, { withCredentials: true })
  return response.data.data
}

export async function fetchOrganizerEventRegistrations(): Promise<OrganizerEventRegistrationInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/registrations`, { withCredentials: true })
  return response.data.data
}

export async function fetchOrganizerEventFeedback(): Promise<OrganizerEventFeedbackInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/feedbacks`, { withCredentials: true })
  return response.data.data
}

export async function fetchOrganizerTickets(): Promise<OrganizerTicketInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/organizer/tickets`, { withCredentials: true })
  return response.data.data
}