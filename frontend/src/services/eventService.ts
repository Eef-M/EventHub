import axios from "axios"
import type { EventInterface } from "@/types/event"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchEvents(): Promise<EventInterface[]> {
  const response = await axios.get<{ data: EventInterface[] }>(`${API_BASE_URL}/events`, { withCredentials: true })
  return response.data.data
}

export async function fetchEventById(id: string): Promise<EventInterface> {
  const response = await axios.get<{ data: EventInterface }>(`${API_BASE_URL}/events/${id}`, { withCredentials: true })
  return response.data.data
}

export async function fetchCreateEvent(payload: FormData): Promise<EventInterface> {
  const response = await axios.post(`${API_BASE_URL}/events`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchUpdateEvent(id: string, payload: FormData): Promise<EventInterface> {
  const response = await axios.put(`${API_BASE_URL}/events/${id}`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchDeleteEvent(id: string): Promise<void> {
  await axios.delete(`${API_BASE_URL}/events/${id}`, { withCredentials: true })
}

export async function fetchEventAvailability(id: string, payload: FormData): Promise<EventInterface> {
  const response = await axios.patch(`${API_BASE_URL}/event/${id}/availability`, payload, { withCredentials: true })
  return response.data.data
} 