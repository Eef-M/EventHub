import axios from "axios"

const API_BASE_URL = 'http://localhost:8000/api/v1'

export const getDashboardStats = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/organizer/dashboard`, {
      withCredentials: true
    })
    return response
  } catch (error) {
    console.error('Error fetching dashboard stats:', error)
    throw error
  }
}