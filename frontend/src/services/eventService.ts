import axios from "axios"
import type { Event } from "../types/event"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchEvents(): Promise<Event[]> {
  const response = await axios.get<{ data: Event[] }>(`${API_BASE_URL}/events`)
  return response.data.data
}

export async function fetchEventById(id: string): Promise<Event> {
  const response = await axios.get<{ data: Event }>(`${API_BASE_URL}/events/${id}`)
  return response.data.data
}

export async function fetchCreateEvent(payload: FormData): Promise<Event> {
  const response = await axios.post(`${API_BASE_URL}/events`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchUpdateEvent(id: string, payload: FormData): Promise<Event> {
  const response = await axios.put(`${API_BASE_URL}/events/${id}`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
  return response.data.data
}

export async function fetchDeleteEvent(id: string): Promise<void> {
  await axios.delete(`${API_BASE_URL}/events/${id}`)
}