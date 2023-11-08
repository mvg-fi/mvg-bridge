<template>
  <div class="d-flex justify-space-between align-center px-5 py-0 setting-item">
    <span class="h7m flex-grow-1">
      {{ $t("show_recent_exchanges") }}
    </span>
    <div class="justify-end align-center">
      <v-switch
        inset
        v-model="showRecentExchanges"
        hide-details
        color="blue"
        flat
      ></v-switch>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useBridgeStore } from "~/stores/bridge/bridge";

const bStore = useBridgeStore();

const showRecentExchanges = computed({
  get() {
    return bStore.recentCardState != 0;
  },
  set(value) {
    if (value) bStore.setRecentCardState(2);
    else bStore.setRecentCardState(0);
  },
});
</script>