<script setup lang="ts">
import Header from "@/components/Header.vue";
import Phone from "@/components/Phone.vue";
import Button from "@/components/Button.vue";
import NoLinksGetStarted from "@/components/NoLinksGetStarted.vue";
import LinkCard from "@/components/LinkCard.vue";
import { useAuthStore } from "@/stores/auth";
import { useProfileStore } from "@/stores/profile";
import { onMounted, ref, watch, type Ref } from "vue";
import { useSortable } from "@vueuse/integrations/useSortable";

const authStore = useAuthStore();
const profileStore = useProfileStore();
const list: Ref<HTMLElement | null> = ref(null);

const initializeSortable = () => {
  if (list.value && profileStore.links?.links) {
    useSortable(list, profileStore.links.links, {
      animation: 300,
      handle: ".handle",
    });
  }
};

onMounted(async () => {
  if (authStore.token) {
    await profileStore.getUserInfo();
    await profileStore.getUserLinks();
  }

  initializeSortable();
});

watch(
  () => profileStore.links?.links,
  () => {
    initializeSortable();
  },
  { deep: true },
);
</script>

<template>
  <div class="flex flex-col min-h-screen max-h-screen">
    <Header />

    <main
      class="flex-grow p-4 md:p-8 flex items-stretch gap-6 md:pt-0 min-h-full"
    >
      <div
        class="hidden xl:grid py-4 flex-shrink-0 place-items-center w-[35rem] bg-neutral-100 rounded-3xl"
      >
        <Phone />
      </div>
      <div
        class="relative overflow-auto bg-neutral-100 rounded-3xl flex-1 p-6 pb-0 md:p-10 md:pb-0"
      >
        <h1 class="text-2xl font-bold text-neutral-500 md:text-[2rem]">
          Customize your links
        </h1>
        <p class="text-neutral-400 mt-2 mb-10">
          Add/edit/remove links below and then share all your profiles with the
          world!
        </p>
        <Button @click="profileStore.addLink" secondary class="w-full">
          + Add new link
        </Button>

        <div class="my-6">
          <NoLinksGetStarted v-if="profileStore.links?.links.length === 0" />

          <section ref="list" class="grid gap-6" v-else>
            <LinkCard
              :key="link.id"
              :link="link"
              :index="index"
              v-for="(link, index) in profileStore.links?.links"
            />
          </section>
        </div>

        <div
          class="sticky bottom-0 left-0 w-full py-4 bg-neutral-100 border-t border-neutral-300 md:py-6 md:flex"
        >
          <Button
            :disabled="!profileStore.changesMade"
            class="w-full md:w-auto md:ml-auto"
            >Save</Button
          >
        </div>
      </div>
    </main>
  </div>
</template>
