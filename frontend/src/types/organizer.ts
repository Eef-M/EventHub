export interface DashboardStatsInterface {
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

export interface OrganizerEventRegistrationInterface {
  id: string;
  username: string;
  email: string;
  event_title: string;
  ticket_name: string;
  status: string;
  registered_at: string;
}

export interface OrganizerEventFeedbackInterface {
  username: string;
  event_title: string;
  comment: string;
  rating: number;
  created_at: string;
}

export interface OrganizerTicketInterface {
  id: string;
  name: string;
  description: string;
  price: number;
  quota: number;
  event_id: string;
  event_title: string;
}