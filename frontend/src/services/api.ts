import axios from "axios"
import type { Event } from "../types/event"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export const getEvents = async (): Promise<Event[]> => {
  const response = await axios.get<{ data: Event[] }>(`${API_BASE_URL}/events`)
  return response.data.data
}

export const getEventById = async (id: string): Promise<Event> => {
  const response = await axios.get<{ data: Event }>(`${API_BASE_URL}/events/${id}`)
  return response.data.data
}