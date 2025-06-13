import type { EventRegistrationsInterface, MyRegistrationsInterface } from "@/types/eventRegistrations";
import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchMyRegistrations(): Promise<MyRegistrationsInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/events/my-registrations`, { withCredentials: true })
  return response.data.data
}

export async function fetchRegisterEvent(id: string, payload: FormData): Promise<EventRegistrationsInterface> {
  const response = await axios.post(`${API_BASE_URL}/events/${id}/register`, payload, { withCredentials: true })
  return response.data.data
}