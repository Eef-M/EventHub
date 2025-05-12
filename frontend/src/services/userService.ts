import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export const getMyProfile = () => {
  return axios.get(`${API_BASE_URL}/user/me`, { withCredentials: true })
}