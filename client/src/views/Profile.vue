<script setup lang="ts">
import Header from "@/components/Header.vue";
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
import Image from "@/components/icons/Image.vue";
import Phone from "@/components/Phone.vue";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";
import { useProfileStore } from "@/stores/profile";

const authStore = useAuthStore();
const router = useRouter();
const profileStore = useProfileStore();

const handleLogout = () => {
  authStore.logout();
  router.replace("/login");
};
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
        class="relative flex flex-col overflow-auto bg-neutral-100 rounded-3xl flex-1 p-6 pb-0 md:p-10 md:pb-0"
      >
        <h1 class="text-2xl font-bold text-neutral-500 md:text-[2rem]">
          Profile Details
        </h1>
        <p class="text-neutral-400 mt-2 mb-10">
          Add your details to create a personal touch to your profile.
        </p>

        <section
          class="bg-neutral-200 p-5 rounded-xl md:flex md:gap-3 md:items-center"
        >
          <h3 class="text-neutral-400 mb-4 md:w-60 md:shrink-0 md:m-0">
            Profile pictre
          </h3>

          <div>
            <label
              for="file-input"
              class="cursor-pointer bg-brand-100 rounded-xl w-48 aspect-square grid place-items-center"
            >
              <span class="text-brand-300 font-semibold">
                <Image class="mb-2 mx-auto fill-brand-300" />
                + Upload Image
              </span>
            </label>
            <input
              id="file-input"
              class="hidden"
              type="file"
              accept="image/png, image/jpeg"
            />
          </div>

          <p class="mt-6 text-xs text-neutral-400 md:m-0">
            Image must be below 1024x1024px. Use PNG or JPG format.
          </p>
        </section>

        <section class="bg-neutral-200 p-5 rounded-xl mt-6 grid gap-3 mb-6">
          <div class="md:flex md:gap-4 md:items-center">
            <label
              class="block text-xs mb-1 w-60 md:text-neutral-400 md:text-base md:m-0"
              for="first_name"
              >First name*</label
            >
            <Input class="md:flex-1" id="first_name" placeholder="Your name" />
          </div>
          <div class="md:flex md:gap-4 md:items-center">
            <label
              class="block text-xs mb-1 w-60 md:text-neutral-400 md:text-base md:m-0"
              for="first_name"
              >Last Name*</label
            >
            <Input
              class="md:flex-1"
              id="first_name"
              placeholder="Your last name"
            />
          </div>
          <div class="md:flex md:gap-4 md:items-center">
            <label
              class="block text-xs mb-1 w-60 md:text-neutral-400 md:text-base md:m-0"
              for="first_name"
              >Email</label
            >
            <Input
              class="md:flex-1"
              id="first_name"
              placeholder="e.g. email@example.com"
            />
          </div>
        </section>

        <Button class="self-start mb-6" secondary @click="handleLogout"
          >Logout</Button
        >

        <div
          class="sticky mt-auto bottom-0 left-0 w-full py-4 bg-neutral-100 border-t border-neutral-300 md:py-6 md:flex"
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
