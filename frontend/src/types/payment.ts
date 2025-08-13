export interface CreatePaymentRequest {
  ticket_id: string;
  quantity: number;
}

export interface CreatePaymentResponse {
  client_secret: string;
  payment_intent_id: string;
  amount: number;
  currency: string;
}

export interface Payment {
  id: string;
  ticket_id: string;
  user_id: string;
  payment_intent_id: string;
  amount: number;
  currency: string;
  status: PaymentStatus;
  quantity: number;
  created_at: string;
  updated_at: string;
  ticket?: {
    id: string;
    name: string;
    description: string;
    price: number;
    event_id: string;
  };
}

export type PaymentStatus = 'pending' | 'succeeded' | 'failed' | 'canceled';

export interface PaymentHistoryResponse {
  payments: Payment[];
  total: number;
  page: number;
  limit: number;
}

export interface StripeElements {
  create(type: string, options?: any): StripeElement;
}

export interface StripeElement {
  mount(element: string | HTMLElement): void;
  unmount(): void;
  destroy(): void;
  on(event: string, handler: (event: any) => void): void;
  update(options: any): void;
}

export interface Stripe {
  elements(options?: any): StripeElements;
  confirmCardPayment(clientSecret: string, options?: any): Promise<{
    error?: any;
    paymentIntent?: any;
  }>;
}