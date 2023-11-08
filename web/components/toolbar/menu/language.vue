<template>
  <!-- Language -->
  <div class="d-flex justify-space-between align-center px-5 py-0 setting-item">
    <span class="h7m">
      {{ $t("language") }}
    </span>
    <button class="px-0 justify-end align-center" @click="moveToLang()">
      {{ activeLang || "EN" }}
      <v-icon size="16">
        <ChevronRightIcon />
      </v-icon>
    </button>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { getNameByKey } from "~/plugins/02-i18n";
import { useBridgeStore } from "~/stores/bridge/bridge";
import {
  EllipsisHorizontalIcon,
  ChevronRightIcon,
} from "@heroicons/vue/20/solid";

const t = useI18n();
const bStore = useBridgeStore();
const moveToLang = () => {
  bStore.setMenuState(1);
};
const activeLang = computed(() => {
  return getNameByKey(t.locale.value);
});
</script>