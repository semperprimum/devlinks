<script setup lang="ts">
import Button from "@/components/Button.vue";
import PlatformLink from "@/components/PlatformLink.vue";
import { useToast } from "@/services/ToastService";
import { useProfileStore } from "@/stores/profile";
import { useClipboard } from "@vueuse/core";
import { computed } from "vue";
import { RouterLink } from "vue-router";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

const profileStore = useProfileStore();
const prefix = computed(() =>
  profileStore.profile?.pic_path.String.startsWith("data") ? "" : BASE_URL,
);

const { copy } = useClipboard();
const { showToast } = useToast();

const copyLink = () => {
  copy(`https://devlinks.thekim.ca/${profileStore.profile?.id}`);
  showToast("The link has been copied to your clipboard!", "copy");
};
</script>

<template>
  <div class="hidden md:pt-6 md:block"></div>
  <div
    class="hidden md:block absolute top-0 left-0 w-full h-[357px] bg-brand-300 rounded-b-[2rem] -z-10"
  ></div>
  <header
    class="relative flex justify-between items-center bg-neutral-100 p-4 pl-6 md:m-6 md:mt-0 md:rounded-2xl max-md:gap-4 flex-wrap"
  >
    <RouterLink to="/"
      ><Button class="max-md:flex-1 min-w-fit" secondary
        >Back to Editor</Button
      ></RouterLink
    >
    <Button @click="copyLink" class="max-md:flex-1 min-w-fit"
      >Share Link</Button
    >
  </header>

  <main
    class="mt-16 md:bg-neutral-100 md:px-14 md:py-12 md:rounded-[1.5rem] md:mx-auto md:max-w-[22rem] md:mt-28 md:shadow-dropdownShadow"
  >
    <img
      class="w-[6.5rem] aspect-square mx-auto rounded-full object-cover border-4 border-brand-300"
      :src="prefix + profileStore.profile?.pic_path.String"
      alt=""
    />
    <h2 v-if="profileStore.profile" class="text-3xl font-bold text-center mt-6">
      {{
        profileStore.profile?.first_name.String +
        " " +
        profileStore.profile?.last_name.String
      }}
    </h2>
    <p v-if="profileStore.profile" class="text-center text-neutral-400 mt-2">
      {{ profileStore.profile.display_email.String }}
    </p>

    <section class="grid gap-5 mt-14 w-full">
      <PlatformLink
        v-for="link in profileStore.links?.links"
        class="max-w-60 mx-auto"
        large
        :platform="link.platform"
        :link="link.url"
      />
    </section>
  </main>
</template>
