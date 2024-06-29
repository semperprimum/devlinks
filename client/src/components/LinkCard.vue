<script setup lang="ts">
import Bars from "@/components/icons/Bars.vue";
import Link from "@/components/icons/Link.vue";
import Select from "@/components/Select.vue";
import Input from "@/components/Input.vue";
import { platformOptions } from "@/platformOptions";
import { useProfileStore } from "@/stores/profile";
import { computed, reactive, watch } from "vue";

const profileStore = useProfileStore();

const props = defineProps<{
  link: {
    id?: number | undefined;
    url: string;
    platform: string;
    order_index?: number | undefined;
    temp_id?: string | undefined;
  };
  index: number;
}>();

const linkCopy = computed(() => reactive({ ...props.link }));

watch(
  () => linkCopy.value.platform,
  (newPlatform) => {
    if (!props.link.id && !props.link.temp_id) return;
    profileStore.updatePlatform(
      props.link.id || props.link.temp_id!,
      newPlatform,
    );
  },
);

watch(
  () => linkCopy.value.url,
  (newUrl) => {
    if (!props.link.id && !props.link.temp_id) return;
    profileStore.updateUrl(props.link.id || props.link.temp_id!, newUrl);
  },
);

const handleRemove = () => {
  if (!props.link.id && !props.link.temp_id) return;
  profileStore.removeLink(props.link.id || props.link.temp_id!);
};
</script>

<template>
  <section ref="r" class="bg-neutral-200 p-5 rounded-xl grid gap-3">
    <header class="flex justify-between items-center">
      <div class="flex gap-2 items-center font-bold text-neutral-400">
        <Bars class="handle | hover:cursor-grab" />
        <h3>Link #{{ index + 1 }}</h3>
      </div>
      <button @click="handleRemove" class="text-neutral-400">Remove</button>
    </header>

    <div>
      <span class="block text-xs pb-1 text-neutral-500">Platform</span>
      <Select
        title="Select a platform"
        v-model="linkCopy.platform"
        :options="platformOptions"
      />
    </div>

    <div>
      <span class="block text-xs pb-1 text-neutral-500">Link</span>
      <Input
        v-model="linkCopy.url"
        :leading="Link"
        placeholder="e.g. https://www.github.com/johnappleseed"
      />
    </div>
  </section>
</template>
