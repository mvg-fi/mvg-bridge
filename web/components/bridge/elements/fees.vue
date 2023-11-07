<template>
  <div
    v-if="rate > 0"
    class="d-flex justify-start align-center rounded-xl mx-2"
    style="height: 48px; border: 1px solid var(--palette-black-5)"
  >
    <div :class="clsx(mobile ? 'ml-4' : 'ml-6', 'flex-grow-1')">
      <span
        :class="clsx(mobile ? 'rate-text-mobile' : 'rate-text')"
        v-if="!rateReserve"
        @click="rateReserve = !rateReserve"
      >
        1 {{ store.fromAsset.symbol }} = {{ rate }} {{ store.toAsset.symbol }}
      </span>
      <span
        :class="clsx(mobile ? 'rate-text-mobile' : 'rate-text')"
        v-else
        @click="rateReserve = !rateReserve"
      >
        1 {{ store.toAsset.symbol }} = {{ rrate }} {{ store.fromAsset.symbol }}
      </span>
      <span class="ml-2 h7m" style="opacity: 40%">
        ({{ usdProfit > 0 ? "+" : "" }}{{ usdProfit }}%)
      </span>
    </div>

    <!-- <span class="mr-6 fee-usd-text"> ${{ usdFee }} </span> -->
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { BN, BN2 } from "~/helpers/bignumber/bn";
import { BridgeFee } from "~/helpers/constants";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();
const { mobile } = useDisplay();

let rateReserve = ref(false);
const rate = computed(() =>
  BN(store.receiveAmount).dividedBy(BN(store.bridgeAmount)).toString()
);
const rrate = computed(() =>
  BN(store.bridgeAmount).dividedBy(BN(store.receiveAmount)).toString()
);
const usdProfit = computed(() =>
  BN2(store.toTotalPrice)
    .minus(BN2(store.fromTotalPrice))
    .dividedBy(BN2(store.fromTotalPrice))
    .multipliedBy(BN2(100))
    .toString()
);
// const usdFee = computed(() =>
//   BN2(store.fromTotalPrice).multipliedBy(BN(BridgeFee)).toString()
// );
</script>

<style scoped>
.rate-text {
  color: var(--palette-black-100);
  font-size: 16px;
  font-family: Satoshi-Medium;
}
.rate-text-mobile {
  color: var(--palette-black-100);
  font-size: 14px;
  font-family: Satoshi-Medium;
}
.fee-usd-text {
  color: var(--palette-black-75);
  font-size: 14px;
  font-family: Satoshi-Medium;
}
</style>