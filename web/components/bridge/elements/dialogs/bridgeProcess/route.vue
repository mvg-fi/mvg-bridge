<template>
  <div class="d-flex align-center justify-center">
    <img :src="bStore.fromAsset.icon" class="img-class mr-2" />
    <span class="h6m">
      {{ bStore.bridgeAmount }} {{ bStore.fromAsset.name }}
    </span>
    <v-icon size="16" style="color: var(--palette-black-75);" class="mx-4">
      <ArrowRightIcon />
    </v-icon>
    <img :src="bStore.toAsset.icon" class="img-class mr-2" />
    <span class="h6m">
      {{ bStore.receiveAmount }} {{ bStore.toAsset.name }}
    </span>
  </div>
</template>

<script setup>
import { ArrowRightIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";

const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.img-class {
  width: 24px;
  height: 24px;
}
</style>