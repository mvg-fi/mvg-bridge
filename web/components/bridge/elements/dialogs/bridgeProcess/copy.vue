<template>
  <div :class="clsx(props.class)">
    <v-btn
      icon
      @click="copy"
      elevation="0"
      style="width: 24px; height: 24px"
      class="rounded-xl bg-transparent align-center justify-center d-flex"
    >
      <v-icon v-if="copied" size="16">
        <CheckIcon />
        <v-tooltip
          activator="parent"
          location="top"
          transition="fade-transition"
        >
          {{ $t("copied") }}
        </v-tooltip>
      </v-icon>
      <v-icon v-else size="16">
        <Copy />
        <v-tooltip
          activator="parent"
          location="top"
          transition="fade-transition"
        >
          {{ $t("copy_address") }}
        </v-tooltip>
      </v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import clsx from "clsx";
import Copy from "./copyIcon.vue";
import { CheckIcon } from "@heroicons/vue/24/solid";

let copied = ref(false);
const props = defineProps(["copyText", "class"]);
const copy = () => {
  copied.value = true;
  navigator.clipboard.writeText(props.copyText);
  setTimeout(() => {
    copied.value = false;
  }, 2000);
};
</script>