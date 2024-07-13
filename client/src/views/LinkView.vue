<script setup lang="ts">
import PlatformLink from "@/components/PlatformLink.vue";
import { useLinkStore } from "@/stores/link";
import { onBeforeUnmount, onMounted } from "vue";
import { RouterLink, useRoute, useRouter } from "vue-router";
import LogoSmall from "@/components/icons/LogoSmall.vue";
import LogoFull from "@/components/icons/LogoFull.vue";
import Spinner from "@/components/icons/Spinner.vue";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

const linkStore = useLinkStore();
const route = useRoute();
const router = useRouter();

onMounted(async () => {
  await linkStore.getLinkById(+route.params.id);

  if (linkStore.isProfileIncomplete) {
    router.replace("/notfound");
  }
});

onBeforeUnmount(() => {
  linkStore.$reset;
});
</script>

<template>
  <div class="hidden md:pt-6 md:block"></div>
  <div
    class="hidden md:block absolute top-0 left-0 w-full h-[357px] bg-brand-300 rounded-b-[2rem] -z-10"
  ></div>

  <header
    class="relative flex justify-center items-center bg-neutral-100 p-4 pl-6 md:m-6 md:mt-0 md:rounded-2xl max-md:gap-4 flex-wrap"
  >
    <RouterLink to="/">
      <LogoSmall class="block md:hidden" />
      <LogoFull class="hidden md:block max-h-8 max-w-36" />
    </RouterLink>
  </header>

  <main
    v-if="!linkStore.isLoading"
    class="mt-16 md:bg-neutral-100 md:px-14 md:py-12 md:rounded-[1.5rem] md:mx-auto md:max-w-[22rem] md:mt-28 md:shadow-dropdownShadow"
  >
    <img
      class="w-[6.5rem] aspect-square mx-auto rounded-full object-cover border-4 border-brand-300"
      :src="BASE_URL + linkStore.data?.user.pic_path.String"
      alt=""
    />
    <h2 class="text-3xl font-bold text-center mt-6">
      {{
        linkStore.data?.user.first_name.String +
        " " +
        linkStore.data?.user.last_name.String
      }}
    </h2>
    <p class="text-center text-neutral-400 mt-2">
      {{ linkStore.data?.user.display_email.String }}
    </p>

    <section class="grid gap-5 mt-14 w-full">
      <PlatformLink
        v-for="link in linkStore.data?.links"
        class="max-w-60 mx-auto"
        large
        :platform="link.platform"
        :link="link.url"
      />
    </section>
  </main>

  <Spinner v-else class="mx-auto w-24 fill-brand-300 animate-spin mt-96" />
</template>
