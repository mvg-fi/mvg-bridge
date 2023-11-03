<template>
  <NuxtLink to="/bridge">
    <v-btn
      size="large"
      :class="EnterAppClass"
      @mouseover="btnHover = true"
      @mouseleave="btnHover = false"
      style="
        height: 40px;
        width: 132px;
        background: linear-gradient(
          90deg,
          var(--palette-main-primary),
          var(--palette-main-secondary)
        );
      "
    >
      <span class="text-background">{{ $t("enter_app") }}</span>
    </v-btn>
  </NuxtLink>
</template>

<script setup>
import clsx from "clsx";

const btnHover = ref(false);
const EnterAppClass = computed(() =>
  clsx(
    "rounded-pill",
    "text-background",
    "text-capitalize",
    "font-weight-bold",
    btnHover.value ? "btn-hover elevation-4" : "elevation-1"
  )
);
</script>