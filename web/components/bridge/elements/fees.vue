<template>
  <div
    v-if="rate > 0"
    class="d-flex justify-start align-center rounded-xl mx-2"
    style="height: 48px; border: 1px solid var(--palette-black-5)"
  >
    <span
      class="ml-6 rate-text flex-grow-1"
      v-if="!rateReserve"
      @click="rateReserve = !rateReserve"
    >
      1 {{ store.fromAsset.symbol }} = {{ rate }} {{ store.toAsset.symbol }}
    </span>
    <span
      class="ml-6 rate-text flex-grow-1"
      v-else
      @click="rateReserve = !rateReserve"
    >
      1 {{ store.toAsset.symbol }} = {{ rrate }} {{ store.fromAsset.symbol }}
    </span>

    <span class="mr-6 fee-usd-text"> ${{ usdFee }} </span>
  </div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();

let rateReserve = ref(false);
const rate = reactive(store.receiveAmount / store.bridgeAmount);
const rrate = reactive(store.bridgeAmount / store.receiveAmount);
const usdFee = 0;
</script>

<style scoped>
.rate-text {
  color: var(--palette-black-100);
  font-size: 16px;
  font-family: Satoshi-Medium;
}
.fee-usd-text {
  color: var(--palette-black-75);
  font-size: 14px;
  font-family: Satoshi-Medium;
}
</style>