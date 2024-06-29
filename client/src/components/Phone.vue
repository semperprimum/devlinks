<script setup lang="ts">
import { useProfileStore } from "@/stores/profile";
import PlatformLink from "@/components/PlatformLink.vue";
import { computed } from "vue";
const BASE_URL = import.meta.env.VITE_API_BASE_URL;

const profileStore = useProfileStore();

const links = computed(() => profileStore?.links?.links.slice(0, 5));

const filler = computed(() => {
  const f: [][] = [];

  if (links.value && links.value.length < 5) {
    for (let i = 0; i < 5 - links.value.length; i++) {
      f.push([]);
    }
  }

  return f;
});
</script>

<template>
  <div
    class="flex flex-col bg-phone bg-no-repeat bg-contain w-[307px] h-[631px] px-9 pt-16 pb-14"
  >
    <div class="flex-shrink-0">
      <img
        v-if="profileStore.profile?.pic_path.Valid"
        class="w-24 aspect-square mx-auto rounded-full object-cover border-4 border-brand-300"
        :src="BASE_URL + profileStore.profile?.pic_path.String"
        alt=""
      />
      <div
        v-else
        class="w-24 aspect-square mx-auto bg-neutral-250 rounded-full"
      ></div>

      <span
        class="block text-center mt-6 text-lg font-semibold"
        v-if="
          profileStore.profile?.first_name.Valid &&
          profileStore.profile.last_name.Valid
        "
        >{{
          profileStore.profile?.first_name.String +
          " " +
          profileStore.profile?.last_name.String
        }}
      </span>
      <div
        v-else
        class="w-40 h-4 rounded-full bg-neutral-250 mx-auto mt-6"
      ></div>

      <span
        class="block text-center text-neutral-400 mt-2 text-sm"
        v-if="profileStore.profile?.display_email.Valid"
        >{{ profileStore.profile.display_email.String }}</span
      >
      <div
        v-else
        class="w-16 h-2 rounded-full bg-neutral-250 mx-auto mt-3"
      ></div>
    </div>

    <div class="grid gap-5 mt-auto">
      <PlatformLink
        v-for="link in links"
        :platform="link.platform"
        :link="link.url"
      />
      <div
        v-for="fill in filler"
        class="w-full h-11 bg-neutral-250 rounded-lg"
      ></div>
    </div>
  </div>
</template>
