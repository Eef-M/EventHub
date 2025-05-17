import { useAuthStore } from "../stores/authStroe"
import { useUserStore } from "../stores/userStore"

export const initAuth = async () => {
  const authStore = useAuthStore();
  const userStore = useUserStore();

  try {
    await userStore.getMyProfile()
    if (userStore.user) {
      authStore.isAuthenticated = true;
      console.log("User authenticated:", userStore.user.role);
    }
  } catch (error) {
    console.error("Auth initialization failed:", error);
    authStore.isAuthenticated = false;
  }

  return {
    isReady: true,
    isAuthenticated: authStore.isAuthenticated,
    user: userStore.user
  };
}