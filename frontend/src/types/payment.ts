export interface PaymentIntentResponse {
  client_secret: string
}

export interface CreatePaymentPayload {
  event_id: string
  amount: number
  currency: string
}