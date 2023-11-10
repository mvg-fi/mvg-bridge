<template>
  <div class="d-flex justify-space-between align-center px-5 py-0 setting-item">
    <span class="h7m flex-grow-1">
      {{ $t("dark_mode") }}
    </span>
    <div class="justify-end align-center">
      <v-switch inset v-model="dark" hide-details color="blue" flat></v-switch>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTheme } from "vuetify";

const theme = useTheme();
const dark = computed({
  get() {
    return theme.global.current.value.dark;
  },
  set(value) {
    if (value) theme.global.name.value = "dark";
    else theme.global.name.value = "light";
  },
});

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? "light" : "dark";
}
</script>