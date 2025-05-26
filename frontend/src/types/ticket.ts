export interface TicketType {
  id: string
  event_id: string
  name: string
  description?: string
  price: number
  quota: number
}