export interface DashboardStats {
  total_events: number
  total_registrations: number
  registered: number
  cancelled_registrations: number
  feedback_received: number
  recent_registrations: {
    username: string
    event_title: string
    ticket_name: string
    status: string
  }[]
}