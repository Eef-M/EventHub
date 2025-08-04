import { fetchPayment } from "@/services/paymentService";
import type { CreatePaymentPayload, PaymentIntentResponse } from "@/types/payment";
import { defineStore } from "pinia";
import { ref } from "vue";

export const usePaymentStore = defineStore('payment', () => {
  const clientSecret = ref<string | null>(null)
  const loading = ref(false)

  const createPayment = async (payload: CreatePaymentPayload) => {
    loading.value = true
    try {
      const data: PaymentIntentResponse = await fetchPayment(payload)
      clientSecret.value = data.client_secret
    } catch (error) {
      console.error('Failed to create payment intent:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    clientSecret,
    loading,
    createPayment
  }
})