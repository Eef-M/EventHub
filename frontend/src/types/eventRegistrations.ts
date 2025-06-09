export interface EventRegistrationsInterface {
  id: string
  user_id: string
  event_id: string
  ticket_id: string
  status: string
}

export interface MyRegistrationsInterface {
  registration_id: string
  event_id: string
  title: string
  description: string
  location: string
  date: string
  time: string
  banner_url: string
  status: string
  total_registrations: number
}