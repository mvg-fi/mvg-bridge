<template>
  <div class="d-flex flex-column">
    <span class="h7m" style="font-size: 16px">
      {{ $t("enhance_privacy") }}
    </span>
    <span
      class="font-weight-light mt-2"
      style="font-size: 0.75rem; color: grey"
    >
      {{ $t("enhance_privacy_info") }}
    </span>

    <v-switch
      v-model="enhancePrivacy"
      hide-details
      inset
      color="blue"
      class="mt-1"
    ></v-switch>
  </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { useSettingsStore } from "~/stores/settings";

const ss = useSettingsStore();
const { enhancePrivacy } = storeToRefs(ss);</script>

<style lang="scss" scoped>
.v-selection-control {
  min-height: 48px !important;
}
</style>