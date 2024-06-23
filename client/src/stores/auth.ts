import axios from "axios";
import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useAuthStore = defineStore("auth", () => {
  const token: Ref<string | null> = ref(
    localStorage.getItem("auth-token") || null,
  );
  const isLoading: Ref<boolean> = ref(false);

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

  const logout = () => {
    setToken(null);
  };

  return { token, isLoading, login, logout, register };
});
