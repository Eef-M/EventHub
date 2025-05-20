import { useAuthStore } from "../stores/authStroe"
import { useUserStore } from "../stores/userStore"

export const initAuth = async () => {
  const authStore = useAuthStore();
  const userStore = useUserStore();

  try {
    await userStore.getMyProfile()
    if (userStore.userState.data) {
      authStore.isAuthenticated = true;
      console.log("User authenticated:", userStore.userState.data.role);
    }
  } catch (error) {
    console.error("Auth initialization failed:", error);
    authStore.isAuthenticated = false;
  }

  return {
    isReady: true,
    isAuthenticated: authStore.isAuthenticated,
    user: userStore.userState.data
  };
}