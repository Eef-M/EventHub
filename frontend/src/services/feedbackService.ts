import type { FeedbackInterface } from "@/types/feedback";
import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchFeedbacks(id: string): Promise<FeedbackInterface[]> {
  const response = await axios.get(`${API_BASE_URL}/events/${id}/feedbacks`, { withCredentials: true })
  return response.data.data
}