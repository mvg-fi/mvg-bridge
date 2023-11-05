<template>
  <!-- Title -->
  <div class="px-5 d-flex justify-space-between align-center text-center">
    <div class="pa-0" style="visibility: hidden;">
      <v-btn
        icon
        size="24"
        elevation="0"
        class="bg-background"
        @click="bStore.mutateBridgeProcess(false)"
      >
        <v-icon size="20">
          <XMarkIcon />
        </v-icon>
      </v-btn>
    </div>
    <span class="h7m" style="font-size: 16px;"> {{ $t("confirm_exchange") }}</span>
    <div>
      <v-btn
        icon
        size="24"
        elevation="0"
        class="bg-background"
        @click="bStore.mutateBridgeProcess(false)"
      >
        <v-icon size="20">
          <XMarkIcon />
        </v-icon>
      </v-btn>
    </div>
  </div>
</template>

<script setup>
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { useConnectStore } from "~/stores/connect/connect";
const bStore = useBridgeStore();
const cStore = useConnectStore();
</script>