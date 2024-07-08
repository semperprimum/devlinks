import axios from "axios";
import { defineStore } from "pinia";
import {
  computed,
  reactive,
  ref,
  type ComputedRef,
  type Reactive,
  type Ref,
} from "vue";
import { useAuthStore } from "@/stores/auth";
import type { Buffer, GetUserLinksResponse, UserInfo, UserLink } from "@/types";
import { useToast } from "@/services/ToastService";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const useProfileStore = defineStore("profile", () => {
  const authStore = useAuthStore();
  const headers: ComputedRef<{ Authorization: string }> = computed(() => {
    return { Authorization: `Bearer ${authStore.token}` };
  });
  const profile: Ref<UserInfo | null> = ref(null);
  const links: Ref<GetUserLinksResponse | null> = ref(null);
  const buffer: Reactive<Buffer> = reactive({
    added: [],
    deleted: [],
    updated: [],
  });
  const changesMade: Ref<boolean> = ref(false);
  const profileChangesMade: Ref<boolean> = ref(false);
  const { showToast } = useToast();

  const getUserInfo = async () => {
    try {
      const response = await axios.get(`${BASE_URL}/user`, {
        headers: headers.value,
      });

      profile.value = response.data;
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  const getUserLinks = async () => {
    try {
      const response = await axios.get(
        `${BASE_URL}/links/${profile.value?.id}`,
        { headers: headers.value },
      );

      links.value = response.data;
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  const updatePlatform = (id: number | string, platform: string) => {
    if (!links.value?.links) return;
    let linkIndex: number | null = null;

    if (typeof id === "number") {
      linkIndex = links.value.links.findIndex((link) => link.id === id);
      if (!buffer.updated.includes(id)) {
        buffer.updated.push(id);
      }
    } else {
      linkIndex = links.value.links.findIndex((link) => link.temp_id === id);
    }

    if (linkIndex !== -1) {
      links.value.links[linkIndex].platform = platform;
      changesMade.value = true;
    }
  };

  const addLink = () => {
    const newLink: UserLink = {
      temp_id: crypto.randomUUID(),
      platform: "github",
      url: "",
    };

    if (!newLink.temp_id) return;

    links.value?.links.push(newLink);
    buffer.added.push(newLink.temp_id);
    changesMade.value = true;
  };

  const updateUrl = (id: number | string, url: string) => {
    if (!links.value?.links) return;
    let linkIndex: number | null = null;

    if (typeof id === "number") {
      linkIndex = links.value.links.findIndex((link) => link.id === id);
      if (!buffer.updated.includes(id)) {
        buffer.updated.push(id);
      }
    } else {
      linkIndex = links.value.links.findIndex((link) => link.temp_id === id);
    }

    if (linkIndex !== -1) {
      links.value.links[linkIndex].url = url;
      changesMade.value = true;
    }
  };

  const removeLink = (id: number | string) => {
    if (!links.value?.links) return;

    if (typeof id === "number") {
      links.value.links = links.value.links.filter((link) => link.id !== id);
      buffer.updated = buffer.updated.filter((linkId) => linkId !== id);
      buffer.deleted.push(id);
      changesMade.value = true;
    } else {
      links.value.links = links.value.links.filter(
        (link) => link.temp_id !== id,
      );
      buffer.added = buffer.added.filter((temp_id) => temp_id !== id);
      changesMade.value = true;
    }
  };

  const updateFirstName = (value: string) => {
    if (!profile.value) return;

    profile.value.first_name.String = value;

    if (profile.value.first_name.String !== "") {
      profile.value.first_name.Valid = true;
    } else {
      profile.value.first_name.Valid = false;
    }
    profileChangesMade.value = true;
  };

  const updateLastName = (value: string) => {
    if (!profile.value) return;

    profile.value.last_name.String = value;

    if (profile.value.last_name.String !== "") {
      profile.value.last_name.Valid = true;
    } else {
      profile.value.last_name.Valid = false;
    }
    profileChangesMade.value = true;
  };

  const updateDispalyEmail = (value: string) => {
    if (!profile.value) return;

    profile.value.display_email.String = value;

    if (profile.value.display_email.String !== "") {
      profile.value.display_email.Valid = true;
    } else {
      profile.value.display_email.Valid = false;
    }
    profileChangesMade.value = true;
  };

  const savePicture = async (file: File) => {
    try {
      const formData = new FormData();
      formData.append("file", file);
      await axios.post(`${BASE_URL}/user/upload`, formData, {
        headers: { ...headers.value, "Content-Type": "multipart/form-data" },
      });
      if (profile.value?.pic_path) {
        profile.value.pic_path.Valid = true;
      }
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  const saveLinks = async () => {
    if (!buffer.added && !buffer.deleted && !buffer.updated) return;

    // Add New Links
    if (buffer.added.length !== 0) {
      for (const id of buffer.added) {
        const link = links.value?.links.find((l) => l.temp_id === id);
        if (!link) continue;

        try {
          const response = await axios.post(
            `${BASE_URL}/user/link`,
            {
              platform: link.platform,
              url: link.url,
              temp_id: link.temp_id,
            },
            { headers: headers.value },
          );

          link.temp_id = undefined;
          link.id = response.data.id;
        } catch (e: any) {
          console.error(e.response?.data?.message || e.message);
        }
      }
    }

    // Update Links
    if (buffer.updated.length !== 0) {
      for (const id of buffer.updated) {
        const link = links.value?.links.find((l) => l.id === id);
        if (!link) continue;

        try {
          await axios.put(
            `${BASE_URL}/user/link/${id}`,
            {
              url: link.url,
              platform: link.platform,
            },
            { headers: headers.value },
          );
        } catch (e: any) {
          console.error(e.response?.data?.message || e.message);
        }
      }
    }

    // Delete Links
    if (buffer.deleted.length !== 0) {
      for (const id of buffer.deleted) {
        try {
          await axios.delete(`${BASE_URL}/user/link/${id}`, {
            headers: headers.value,
          });
        } catch (e: any) {
          console.error(e.response?.data?.message || e.message);
        }
      }
    }

    // Update Links Order
    const link_ids = links.value?.links.map((link) => link.id);
    try {
      await axios.put(
        `${BASE_URL}/user/link`,
        {
          link_ids,
        },
        { headers: headers.value },
      );
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }

    buffer.updated = [];
    buffer.deleted = [];
    buffer.added = [];
    changesMade.value = false;
    showToast("Your changes have been successfully saved!", "save");
  };

  const saveProfile = async (
    firstName: string,
    lastName: string,
    email?: string,
  ) => {
    try {
      axios.put(
        `${BASE_URL}/user/update`,
        {
          first_name: firstName,
          last_name: lastName,
          display_email: email || undefined,
        },
        { headers: headers.value },
      );

      profileChangesMade.value = false;
      showToast("Your changes have been successfully saved!", "save");
    } catch (e: any) {
      console.error(e.response?.data?.message || e.message);
    }
  };

  const $reset = () => {
    profile.value = null;
    links.value = null;
    buffer.updated = [];
    buffer.deleted = [];
    buffer.added = [];
    changesMade.value = false;
  };

  return {
    getUserInfo,
    getUserLinks,
    updatePlatform,
    updateUrl,
    removeLink,
    addLink,
    updateFirstName,
    updateDispalyEmail,
    updateLastName,
    savePicture,
    saveLinks,
    saveProfile,
    buffer,
    changesMade,
    profileChangesMade,
    profile,
    links,
    $reset,
  };
});
