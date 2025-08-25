export interface Ticket {
  id: string;
  event_id: string;
  name: string;
  description?: string;
  price: number;
  quota: number;
  ticket_code: string;
  created_at: string;
  updated_at: string;
}

export interface MyTicket {
  ticket_id: string;
  event_id: string;
  title: string;
  location: string;
  date: string;
  ticket_code: string;
  price: number;
}

export interface CreateTicketRequest {
  event_id: string;
  name: string;
  description?: string;
  price: number;
  quota: number;
  ticket_code?: string;
}

export interface UpdateTicketRequest {
  name?: string;
  description?: string;
  price?: number;
  quota?: number;
  ticket_code?: string;
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
  ticket_code?: string;
  event_id?: string;
}

export type TicketPurchaseStatus = 'pending' | 'confirmed' | 'cancelled' | 'refunded';

export interface TicketResponse {
  tickets: Ticket[];
  total: number;
  page: number;
  limit: number;
}

export interface TicketAvailability {
  ticket_id: string;
  available: boolean;
  remaining_quota: number;
  message?: string;
}

export interface TicketDisplay extends Ticket {
  is_sold_out: boolean;
  display_price: string; // formatted price
  remaining_text: string; // e.g., "5 tickets left"
}