<template>
  <div
    class="d-flex flex-column h7m mx-6"
    style="gap: 8px; color: var(--palette-black-100)"
  >
    <div class="d-flex justify-space-between align-center">
      <span> {{ $t("recipient") }} </span>
      <span style="width: 280px; overflow-wrap: break-word" class="text-end">
        {{ bStore.receiver || $t("default") }}
      </span>
    </div>

    <div class="d-flex justify-space-between align-center">
      <span>
        {{ $t("estimated_time") }}
        <v-icon size="14">
          <v-tooltip
            activator="parent"
            location="bottom"
            transition="fade-transition"
          >
            {{ $t("estimated_time_intro") }}
          </v-tooltip>
          <CircleWarningIcon />
        </v-icon>
      </span>
      <span> {{ estimated }} </span>
    </div>

    <div class="d-flex justify-space-between align-center">
      <span>
        {{ $t("cost") }}
        <v-icon size="14">
          <v-tooltip
            activator="parent"
            location="bottom"
            transition="fade-transition"
          >
            {{ $t("cost_intro") }}
          </v-tooltip>
          <CircleWarningIcon />
        </v-icon>
      </span>
      <span> {{ bStore.costUSD }} </span>
    </div>
  </div>
</template>

<script setup>
import CircleWarningIcon from "./circleWarningIcon.vue";
import { useBridgeStore } from "~/stores/bridge/bridge";
const bStore = useBridgeStore();
const estimated = "1 Minute";
const getDefaultReceiver = () => {};
const getEstimatedBridgeTime = () => {};
</script>