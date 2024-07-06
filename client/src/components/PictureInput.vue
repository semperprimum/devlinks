<script setup lang="ts">
import Image from "@/components/icons/Image.vue";
import { useProfileStore } from "@/stores/profile";
import { computed, ref, type Ref } from "vue";

const BASE_URL = import.meta.env.VITE_API_BASE_URL;
const profileStore = useProfileStore();

const handleFileUpload = async (event: Event) => {
  const input = event.target as HTMLInputElement;
  const file: File | null = input.files ? input.files[0] : null;
  if (file) {
    const reader = new FileReader();
    reader.onload = () => {
      if (profileStore.profile?.pic_path && typeof reader.result === "string") {
        profileStore.profile.pic_path.String = reader.result;
      }
    };
    reader.readAsDataURL(file);

    await profileStore.savePicture(file);
  }
};

const backgroundStyle = computed(() => {
  const prefix = profileStore.profile?.pic_path.String.startsWith("data")
    ? ""
    : BASE_URL;

  return {
    backgroundImage: profileStore.profile?.pic_path.Valid
      ? `linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url(${prefix + profileStore.profile?.pic_path.String})`
      : undefined,
  };
});
</script>

<template>
  <div>
    <label
      for="file-input"
      class="cursor-pointer rounded-xl w-48 aspect-square grid place-items-center"
      :class="{
        'bg-cover bg-center': profileStore.profile?.pic_path.Valid,
        'bg-brand-100': !profileStore.profile?.pic_path.Valid,
      }"
      :style="backgroundStyle"
    >
      <span
        class="font-semibold"
        :class="{
          'text-neutral-100': profileStore.profile?.pic_path.Valid,
          'text-brand-300': !profileStore.profile?.pic_path.Valid,
        }"
      >
        <Image
          class="mb-2 mx-auto"
          :class="{
            'fill-neutral-100': profileStore.profile?.pic_path.Valid,
            'fill-brand-300': !profileStore.profile?.pic_path.Valid,
          }"
        />
        {{
          profileStore.profile?.pic_path.Valid
            ? "Change Image"
            : "+ Upload Image"
        }}
      </span>
    </label>
    <input
      @change="handleFileUpload"
      id="file-input"
      class="hidden"
      type="file"
      accept="image/png, image/jpeg"
    />
  </div>
</template>
