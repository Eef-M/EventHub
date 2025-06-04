export interface TicketInterface {
  id: string
  event_id: string
  name: string
  description?: string
  price: number
  quota: number
}

export interface MyTicketInterface {
  ticket_id: string
  event_id: string
  title: string
  location: string
  date: string
  ticket_code: string
  price: number
}