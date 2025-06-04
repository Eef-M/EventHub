import type { MyRegistrationsInterface } from "@/types/eventRegistrations";
import axios from "axios";

const API_BASE_URL = 'http://localhost:8000/api/v1'

export async function fetchMyRegistrations(): Promise<MyRegistrationsInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/events/my-registrations`, { withCredentials: true })
  return response.data.data
}