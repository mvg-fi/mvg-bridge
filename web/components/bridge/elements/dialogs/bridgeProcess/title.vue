<template>
  <!-- Title -->
  <v-row class="mt-3 align-center text-center" no-gutters>
    <v-col cols="1" class="pa-0"> </v-col>
    <v-col cols="10" class="pa-0">
      <span class="h6">
        {{
          $t("exchange_from_to", {
            from: bStore.fromAsset.symbol,
            to: bStore.toAsset.symbol,
          })
        }}</span
      >
    </v-col>
    <v-col cols="1" class="pa-0">
      <v-btn
        icon
        elevation="0"
        class="mb-8"
        style="background-color: var(--palette-background-1)"
        @click="bStore.mutateBridgeProcess(false)"
      >
        <v-icon size="20">
          <XMarkIcon />
        </v-icon>
      </v-btn>
    </v-col>
  </v-row>
</template>

<script setup>
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { useConnectStore } from "~/stores/connect/connect";
const bStore = useBridgeStore();
const cStore = useConnectStore();
</script>