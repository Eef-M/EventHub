export interface RegisterPayload {
  username: string
  first_name: string
  last_name: string
  email: string
  password: string
  role: string
}

export interface LoginPayload {
  email: string
  password: string
}