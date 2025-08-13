import type { CreatePaymentRequest, CreatePaymentResponse, Payment, PaymentHistoryResponse } from "@/types/payment"
import axios from "axios"

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

const api = axios.create({
  baseURL: API_BASE_URL
})

export class PaymentService {
  static async createPaymentIntent(data: CreatePaymentRequest): Promise<CreatePaymentResponse> {
    try {
      const response = await api.post('/payments/create', data)
      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.error || 'Failed to create payment')
    }
  }

  static async getPaymentHistory(page = 1, limit = 10): Promise<PaymentHistoryResponse> {
    try {
      const response = await api.get('/payments/history', {
        params: { page, limit }
      });
      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.error || 'Failed to get payment history')
    }
  }

  static async getPayment(paymentId: string): Promise<Payment> {
    try {
      const response = await api.get(`/payments/${paymentId}`)
      return response.data.payment
    } catch (error: any) {
      throw new Error(error.response?.data?.error || 'Failed to get payment details');
    }
  }
}

export const loadStripe = async (): Promise<any> => {
  if (window.Stripe) {
    return window.Stripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY)
  }

  const script = document.createElement('script');
  script.src = 'https://js.stripe.com/v3/';
  document.head.appendChild(script);

  return new Promise((resolve, reject) => {
    script.onload = () => {
      if (window.Stripe) {
        resolve(window.Stripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY))
      } else {
        reject(new Error('Stripe failed to load'));
      }
    };
    script.onerror = reject;
  });
};

declare global {
  interface Window {
    Stripe: any;
  }
}