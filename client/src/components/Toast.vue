<template>
  <div
    v-if="visible"
    class="appear | flex items-center gap-2 fixed bg-neutral-500 px-6 py-4 rounded-xl bottom-10 left-1/2 -translate-x-1/2 font-semibold text-base text-neutral-100"
    :class="{
      'bg-success': type === 'success',
      'bg-red': type === 'error',
    }"
  >
    <component v-if="type === 'save'" :is="FloppyDisc" />
    <component v-if="type === 'copy'" class="fill-neutral-400" :is="Link" />
    <component
      v-if="type === 'success'"
      class="fill-neutral-100 w-5 aspect-square"
      :is="Check"
    />
    <component
      v-if="type === 'error'"
      class="fill-neutral-100 w-5 aspect-square"
      :is="CircleExclamation"
    />
    {{ toastMessage }}
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { useToast } from "@/services/ToastService";
import { nextTick } from "vue";
import FloppyDisc from "@/components/icons/FloppyDisc.vue";
import Link from "@/components/icons/Link.vue";
import Check from "@/components/icons/Check.vue";
import CircleExclamation from "@/components/icons/CircleExclamation.vue";

const props = withDefaults(
  defineProps<{
    message: string;
    type?: string;
    duration?: number;
  }>(),
  {
    type: "save",
    duration: 3500,
  },
);

const { toastMessage, toastType } = useToast();

const visible = ref(false);
const timer = ref<number | null>(null);

const showToast = () => {
  if (timer.value) {
    clearTimeout(timer.value);
  }
  visible.value = false;
  nextTick(() => {
    visible.value = true;
  });

  timer.value = setTimeout(() => {
    visible.value = false;
  }, props.duration);
};

watch(
  () => toastMessage.value,
  (newMessage) => {
    if (newMessage) {
      showToast();
    }
  },
);

watch(
  () => toastType.value,
  (_) => {
    if (toastMessage.value) {
      showToast();
    }
  },
);
</script>

<style scoped>
.appear {
  animation: appear 200ms ease forwards;
}

@keyframes appear {
  from {
    transform: translate(-50%, 5rem);
    opacity: 0;
  }
  to {
    transform: translate(-50%, 0);
    opacity: 1;
  }
}
</style>
