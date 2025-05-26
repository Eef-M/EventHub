import type { TicketType } from "@/types/ticket"
import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchCreateTicket(payload: FormData): Promise<TicketType> {
  const response = await axios.post(`${API_BASE_URL}/tickets`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}

export async function fetchUpdateTicket(id: string, payload: FormData): Promise<TicketType> {
  const response = await axios.put(`${API_BASE_URL}/tickets/${id}`, payload, {
    headers: { 'Content-Type': 'multipart/form-data' },
    withCredentials: true,
  })
  return response.data.data
}