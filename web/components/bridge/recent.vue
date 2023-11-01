<template>
  <v-sheet class="recent-sheet">
    <!-- Title -->
    <v-btn
      icon
      elevation="0"
      class="mb-8"
      style="background-color: var(--palette-background-1)"
      @click="bStore.setRecentCardState(false)"
    >
      <v-icon size="20">
        <ChevronDownIcon />
      </v-icon>
    </v-btn>
  </v-sheet>
</template>

<script setup>
import { ChevronDownIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.recent-sheet {
  height: 344px;
  width: 308px;
}
</style>