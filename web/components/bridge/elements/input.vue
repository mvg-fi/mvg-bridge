<template>
  <div class="mx-2">
    <div
      :class="
        clsx(
          'd-flex',
          'flex-column',
          props.from ? 'input-top-bg' : 'input-down-bg'
        )
      "
    >
      <div class="ml-6 mt-4">
        <span class="h7m" style="opacity: 50%">
          {{ props.from ? $t("you_pay") : $t("you_receive") }}
        </span>
      </div>

      <div class="d-flex flex-row ml-6">
        <TextField :from="props.from" />
        <Asset :from="props.from" />
      </div>

      <div class="ml-6 h7m mb-4">
        <span style="color: var(--palette-black-50)">
          ${{ fromUSDPrice }}
        </span>
        <!-- <span style="opacity: 0" v-else> $ </span> -->
      </div>
    </div>
  </div>
</template>

<script setup>
import { asyncComputed } from "@vueuse/core";
import clsx from "clsx";
import Asset from "~/components/bridge/elements/input/asset.vue";
import TextField from "~/components/bridge/elements/input/textfield.vue";
import { BN } from "~/helpers/bignumber/bn";
import { fetchUSDPrice, useBridgeStore } from "~/stores/bridge/bridge";
const props = defineProps(["from"]);
const store = useBridgeStore();

let fromUSD = asyncComputed(async() => {
  return await fetchUSDPrice(store.fromAsset)
})
let fromUSDPrice = asyncComputed(async () => {
  console.log("fromUSDPrice");
  // if (!store.bridgeAmount) return 0;
  return BN(store.bridgeAmount).multipliedBy(BN(fromUSD)).toString();
});
</script>

<style lang="scss" scoped>
.input-top-bg {
  background-color: var(--palette-black-5);
  border-top-left-radius: 32px;
  border-top-right-radius: 32px;
}
.input-down-bg {
  background-color: var(--palette-black-5);
  border-bottom-left-radius: 32px;
  border-bottom-right-radius: 32px;
}
</style>