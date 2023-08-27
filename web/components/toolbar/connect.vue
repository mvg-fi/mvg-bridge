<template>
  <v-btn
    size="large"
    @click="connect"
    :class="ConnectBtnClass"
    @mouseover="btnHover = true"
    @mouseleave="btnHover = false"
  >
    {{ $t("connect") }}
  </v-btn>
</template>

<script setup>
import clsx from "clsx";

const btnHover = ref(false);
const ConnectBtnClass = computed(() =>
  clsx(
    "bg-primary",
    "rounded-pill",
    "text-capitalize",
    "font-weight-bold",
    btnHover.value ? "btn-hover elevation-4" : "elevation-1"
  )
);

const connect = async () => {
  try {
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