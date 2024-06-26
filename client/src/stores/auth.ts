import axios, { type AxiosResponse } from "axios";
import { defineStore } from "pinia";
import { computed, ref, type ComputedRef, type Ref } from "vue";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useAuthStore = defineStore("auth", () => {
  const token: Ref<string | null> = ref(
    localStorage.getItem("auth-token") || null,
  );
  const isLoading: Ref<boolean> = ref(false);
  const headers: ComputedRef<{ Authorization: string }> = computed(() => {
    return { Authorization: `Bearer ${token.value}` };
  });

  const setToken = (newToken: string | null) => {
    if (newToken) {
      localStorage.setItem("auth-token", newToken);
      token.value = newToken;
    } else {
      localStorage.removeItem("auth-token");
      token.value = null;
    }
  };

  const register = async (email: string, password: string) => {
    isLoading.value = true;
    try {
      await axios.post(`${BASE_URL}/register`, {
        email,
        password,
      });
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    } finally {
      isLoading.value = false;
    }
  };

  const login = async (email: string, password: string) => {
    isLoading.value = true;
    try {
      const response = await axios.post(`${BASE_URL}/login`, {
        email,
        password,
      });

      setToken(response.data.token);
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    } finally {
      isLoading.value = false;
    }
  };

  const getUserInfo = async (): Promise<AxiosResponse<any, any> | null> => {
    try {
      const response = await axios.get(`${BASE_URL}/user`, {
        headers: headers.value,
      });
      return response;
    } catch (error: any) {
      if (axios.isAxiosError(error) && error.response) {
        console.log(
          "getUserInfo error response status:",
          error.response.status,
        );
      } else {
        console.log("getUserInfo error:", error.message);
      }
      return null;
    }
  };

  const checkAuth = async (): Promise<boolean> => {
    if (!token.value) {
      console.log("No token found");
      return false;
    }

    const response = await getUserInfo();

    if (!response || response.status !== 200) {
      logout();
      return false;
    }

    return true;
  };

  const logout = () => {
    setToken(null);
  };

  return { token, isLoading, login, logout, register, getUserInfo, checkAuth };
});
