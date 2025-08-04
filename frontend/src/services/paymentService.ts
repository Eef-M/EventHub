import type { CreatePaymentPayload, PaymentIntentResponse } from "@/types/payment";
import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export async function fetchPayment(payload: CreatePaymentPayload): Promise<PaymentIntentResponse> {
  const response = await axios.post(`${API_BASE_URL}/payments/webhoook`, payload, { withCredentials: true })
  return response.data
}