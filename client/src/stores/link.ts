import type { GetUserLinksResponse } from "@/types";
import axios from "axios";
import { defineStore } from "pinia";
import { computed, ref, type ComputedRef, type Ref } from "vue";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useLinkStore = defineStore("link", () => {
  const token: Ref<string | null> = ref(
    localStorage.getItem("auth-token") || null,
  );
  const data: Ref<GetUserLinksResponse | null> = ref(null);
  const headers: ComputedRef<{ Authorization: string }> = computed(() => {
    return { Authorization: `Bearer ${token.value}` };
  });
  const isProfileIncomplete: Ref<boolean> = ref(false);
  const isLoading: Ref<boolean> = ref(false);

  const getLinkById = async (id: number) => {
    isLoading.value = true;
    try {
      const response = await axios.get(`${BASE_URL}/links/${id}`, {
        headers: headers.value,
      });

      data.value = response.data;

      if (
        !data.value?.user.pic_path.Valid ||
        !data.value.user.first_name.Valid ||
        !data.value.user.last_name.Valid ||
        !data.value.links
      ) {
        isProfileIncomplete.value = true;
      }
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    } finally {
      isLoading.value = false;
    }
  };

  const $reset = () => {
    data.value = null;
    isProfileIncomplete.value = false;
  };

  return { getLinkById, data, isProfileIncomplete, isLoading, $reset };
});
