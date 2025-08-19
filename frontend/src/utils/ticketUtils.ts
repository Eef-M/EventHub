import type { Ticket, TicketDisplay } from '@/types/ticket'

export function formatPrice(price: number): string {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(price)
}

export function transformToTicketDisplay(ticket: Ticket): TicketDisplay {
  const isSoldOut = ticket.quota <= 0
  const displayPrice = formatPrice(ticket.price)

  let remainingText = ''
  if (isSoldOut) {
    remainingText = 'Sold out'
  } else if (ticket.quota <= 10) {
    remainingText = `Only ${ticket.quota} left`
  } else {
    remainingText = `${ticket.quota} available`
  }

  return {
    ...ticket,
    is_sold_out: isSoldOut,
    display_price: displayPrice,
    remaining_text: remainingText
  }
}

export function validateTicketPurchase(
  ticket: Ticket,
  requestedQuantity: number
): { valid: boolean; message?: string } {
  if (ticket.quota <= 0) {
    return {
      valid: false,
      message: 'This ticket is sold out'
    }
  }

  if (requestedQuantity <= 0) {
    return {
      valid: false,
      message: 'Quantity must be at least 1'
    }
  }

  if (requestedQuantity > ticket.quota) {
    return {
      valid: false,
      message: `Only ${ticket.quota} tickets available`
    }
  }

  if (requestedQuantity > 10) {
    return {
      valid: false,
      message: 'Maximum 10 tickets per purchase'
    }
  }

  return { valid: true }
}

export function calculateTotalPrice(
  ticketPrice: number,
  quantity: number,
  processingFeeRate: number = 0.03
): {
  subtotal: number;
  processingFee: number;
  total: number;
} {
  const subtotal = ticketPrice * quantity
  const processingFee = subtotal * processingFeeRate
  const total = subtotal + processingFee

  return {
    subtotal,
    processingFee,
    total
  }
}

export function generateTicketCodePreview(eventTitle?: string): string {
  const prefix = eventTitle
    ? eventTitle.substring(0, 3).toUpperCase().replace(/[^A-Z]/g, 'X')
    : 'TKT'

  const timestamp = Date.now().toString().slice(-6)
  const random = Math.random().toString(36).substring(2, 5).toUpperCase()

  return `${prefix}-${timestamp}-${random}`
}

export function isTicketPurchasable(ticket: Ticket, eventDate?: string): boolean {
  if (ticket.quota <= 0) return false
  if (ticket.price < 0) return false

  if (eventDate) {
    const eventDateObj = new Date(eventDate)
    const now = new Date()
    if (eventDateObj < now) return false
  }

  return true
}

export function sortTicketsByPrice(tickets: Ticket[], ascending: boolean = true): Ticket[] {
  return [...tickets].sort((a, b) => {
    return ascending ? a.price - b.price : b.price - a.price
  })
}

export function filterAvailableTickets(tickets: Ticket[]): Ticket[] {
  return tickets.filter(ticket => ticket.quota > 0)
}

export function getCheapestTicket(tickets: Ticket[]): Ticket | null {
  const availableTickets = filterAvailableTickets(tickets)
  if (availableTickets.length === 0) return null

  return availableTickets.reduce((cheapest, current) =>
    current.price < cheapest.price ? current : cheapest
  )
}

export function formatTicketSummary(
  ticket: Ticket,
  quantity: number
): string {
  return `${ticket.name} × ${quantity} = ${formatPrice(ticket.price * quantity)}`
}