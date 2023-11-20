<template>
  <div class="d-flex flex-column">
    <span class="h7m" style="font-size: 16px">
      {{ $t("use_custom_address") }}
    </span>
    <span
      class="font-weight-light mt-2"
      style="font-size: 0.75rem; color: grey"
    >
      {{ $t("use_custom_address_info") }}
    </span>

    <v-switch
      v-model="useCustomAddress"
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
const { useCustomAddress } = storeToRefs(ss);
</script>

<style lang="scss" scoped>
.v-selection-control {
  min-height: 48px !important;
}
</style>