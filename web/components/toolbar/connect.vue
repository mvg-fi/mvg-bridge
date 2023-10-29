<template>
  <v-btn
    size="large"
    @click="connect"
    :class="ConnectBtnClass"
    @mouseover="btnHover = true"
    @mouseleave="btnHover = false"
    style="background: linear-gradient(90deg, var(--palette-main-primary), var(--palette-main-secondary))"
  >
      <span style="color: var(--palette-background-1)">{{ $t("connect") }}</span>
  </v-btn>
</template>

<script setup>
import clsx from "clsx";
import { useConnectStore } from "~/stores/connect/connect";
const cStore = useConnectStore()

const btnHover = ref(false);
const ConnectBtnClass = computed(() =>
  clsx(
    "rounded-pill",
    "text-capitalize",
    "font-weight-bold",
    btnHover.value ? "btn-hover elevation-4" : "elevation-1"
  )
);

const connect = async () => {
  try {
    cStore.mutateDialog(true)
  } catch (e) {
    console.log(e);
  }
};
</script>

<style lang="scss" scoped>
.btn-hover {
  opacity: 0.85;
}
</style>