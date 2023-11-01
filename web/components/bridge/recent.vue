<template>
  <v-sheet class="recent-sheet px-4 bg-background" elevation="3">
    <!-- Title -->
    <div class="d-flex justify-space-between align-center mt-6 mx-4">
      <span class="h7" style="font-size: 16px;">
        {{ $t("recent_exchanges") }}
      </span>
      <button
        class="bg-background px-0"
        @click="bStore.setRecentCardState(1)"
      >
        <v-icon size="20">
          <ChevronDownIcon />
        </v-icon>
      </button>
    </div>

    <!-- Items -->
    <div class="d-flex flex-column mt-4">
      <!-- Icon -->
      <div class="">

      </div>

      <!-- Amount -->
      <div>

      </div>

      <!-- State -->
    </div>
  </v-sheet>
</template>

<script setup>
import { ChevronDownIcon } from "@heroicons/vue/20/solid";
import { useBridgeStore } from "~/stores/bridge/bridge";
const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.recent-sheet {
  height: 344px;
  width: 308px;
  border-top-left-radius: 32px;
  border-top-right-radius: 32px;
}
</style>