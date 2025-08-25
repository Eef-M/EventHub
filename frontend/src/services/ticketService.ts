import type { MyTicket, Ticket } from "@/types/ticket"
import axios from "axios"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchTickets(): Promise<Ticket[]> {
  const response = await axios.get(`${API_BASE_URL}/tickets`, { withCredentials: true })
  return response.data.data
}

export async function fetchMyTickets(): Promise<MyTicket[]> {
  const response = await axios.get(`${API_BASE_URL}/tickets/my-tickets`, { withCredentials: true })
  return response.data.data
}

export async function fetchCreateTicket(payload: FormData): Promise<Ticket> {
  const response = await axios.post(`${API_BASE_URL}/tickets`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchUpdateTicket(id: string, payload: FormData): Promise<Ticket> {
  const response = await axios.put(`${API_BASE_URL}/tickets/${id}`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchDeleteTicket(id: string): Promise<void> {
  await axios.delete(`${API_BASE_URL}/tickets/${id}`, { withCredentials: true })
}