<template>
  <div class="text-center">
    <span class="h5">
      {{
        $t("exchange_from_to", {
          from: bStore.fromAsset.symbol,
          to: bStore.toAsset.symbol,
        })
      }}
    </span>
  </div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";

const bStore = useBridgeStore();
</script>

<style>
</style>