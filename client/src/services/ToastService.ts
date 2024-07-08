import { nextTick, ref } from "vue";

const toastMessage = ref("");
const toastType = ref("save");

export function useToast() {
  const showToast = (message: string, type: string = "save") => {
    toastMessage.value = "";
    toastType.value = "";

    nextTick(() => {
      toastMessage.value = message;
      toastType.value = type;
    });
  };

  return {
    toastMessage,
    toastType,
    showToast,
  };
}
