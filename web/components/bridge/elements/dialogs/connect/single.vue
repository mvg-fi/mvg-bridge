<template>
  <button
    :key="props.w.name"
    class="d-flex flex-row align-center mb-3"
    style="height: 40px; gap: 8px"
  >
    <v-img
      class="flex-grow-0 rounded-circle"
      :src="props.w.icon"
      style="width: 32px; height: 32px"
    />

    <span
      class="flex-grow-1 justify-start text-start"
      style="font-size: 16px; line-height: 32px; font-family: Satoshi-SemiBold"
    >
      {{ props.w.name }}
    </span>

    <div v-if="connected">
      <v-icon size="16" color="primary">
        <CheckIcon />
      </v-icon>
    </div>

    <div v-if="loading">
      <v-progress-circular
        indeterminate
        color="primary"
        size="16"
        width="2"
      ></v-progress-circular>
    </div>
  </button>
</template>

<script lang="ts" setup>
import { CheckIcon } from "@heroicons/vue/24/outline";
import { useConnectStore } from "~/stores/connect/connect";
import type { Wallet } from "~/types/wallet";
const props = defineProps(["w"]);
const cStore = useConnectStore();
const connected = computed(() => {
  return cStore.connectedWallets.find((w: Wallet) => {
    w.name === props.w.name;
  });
});
const loading = computed(() => {
  return cStore.connecting.name === props.w.name;
});
</script>


<style lang="scss" scoped>
</style>