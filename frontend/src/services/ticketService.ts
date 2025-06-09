import type { MyTicketInterface, TicketInterface } from "@/types/ticket"
import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchTickets(): Promise<TicketInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/tickets`, { withCredentials: true })
  return response.data.data
}

export async function fetchMyTickets(): Promise<MyTicketInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/tickets/my-tickets`, { withCredentials: true })
  return response.data.data
}

export async function fetchCreateTicket(payload: FormData): Promise<TicketInterface> {
  const response = await axios.post(`${API_BASE_URL}/tickets`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchUpdateTicket(id: string, payload: FormData): Promise<TicketInterface> {
  const response = await axios.put(`${API_BASE_URL}/tickets/${id}`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchDeleteTicket(id: string): Promise<void> {
  await axios.delete(`${API_BASE_URL}/tickets/${id}`, { withCredentials: true })
}