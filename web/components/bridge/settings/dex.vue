<template>
  <div class="d-flex flex-column">
    <span class="h7m" style="font-size: 16px">
      {{ $t("only_use_dex") }}
    </span>
    <span
      class="font-weight-light mt-2"
      style="font-size: 0.75rem; color: grey"
    >
      {{ $t("only_use_dex_info") }}
    </span>

    <v-switch
      v-model="onlyUseDex"
      hide-details
      inset
      color="blue"
      class="mt-1"
    ></v-switch>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { useSettingsStore } from '~/stores/settings';

const ss = useSettingsStore()
const { onlyUseDex } = storeToRefs(ss)
</script>

<style>
</style>