export interface DashboardStats {
  total_events: number
  total_registrations: number
  pending_approvals: number
  feedback_received: number
  recent_registrations: {
    user_name: string
    event_title: string
  }[]
}