export interface Ticket {
  id: string;
  name: string;
  description: string;
  price: number;
  quota: number;
  event_id: string;
  created_at: string;
  updated_at: string;
}

export interface CreateTicketRequest {
  name: string;
  description: string;
  price: number;
  quota: number;
  event_id: string;
}

export interface UpdateTicketRequest {
  name?: string;
  description?: string;
  price?: number;
  quota?: number;
}

export interface TicketPurchase {
  id: string;
  ticket_id: string;
  user_id: string;
  quantity: number;
  total_amount: number;
  status: TicketPurchaseStatus;
  created_at: string;
  updated_at: string;
  ticket?: Ticket;
}

export type TicketPurchaseStatus = 'pending' | 'confirmed' | 'cancelled' | 'refunded';

export interface TicketResponse {
  tickets: Ticket[];
  total: number;
  page: number;
  limit: number;
}