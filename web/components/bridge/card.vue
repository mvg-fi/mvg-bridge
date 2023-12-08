<template>
  <v-row no-gutters class="fill-height d-flex justify-center">
    <div>
      <v-card
        :class="
          clsx(
            !mobile && 'bridge-card mt-10',
            mobile && 'mobile-bridge-card mt-10',
            'rounded-xl pa-0 d-flex flex-column'
          )
        "
        elevation="0"
      >
        <Title class="mx-6 mt-4" />

        <div class="mt-4">
          <Input :from="true" />
          <Switch />
          <Input :from="false" />
        </div>

        <Receiver v-if="showReceiver" class="my-2" />
        <Fees class="my-2" />
        <Confirm class="my-2" />
      </v-card>
    </div>
  </v-row>
</template>

<script setup>
import Title from "~/components/bridge/elements/title.vue";
import Input from "~/components/bridge/elements/input.vue";
import Switch from "~/components/bridge/elements/switch.vue";
import Receiver from "~/components/bridge/elements/receiver.vue";
import Fees from "~/components/bridge/elements/fees.vue";
import Confirm from "~/components/bridge/elements/confirm.vue";
import { useDisplay } from "vuetify";
import clsx from "clsx";
import { useSettingsStore } from "~/stores/settings";
import { storeToRefs } from "pinia";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { isEvmName, useConnectStore } from "~/stores/connect/connect";
import { isEVMAsset } from "~/helpers/networks";

const bs = useBridgeStore();
const cs = useConnectStore();
const ss = useSettingsStore();
const { useCustomAddress } = storeToRefs(ss);
const { fromAsset } = storeToRefs(bs);
const { connectedWallets } = storeToRefs(cs);

const { mobile } = useDisplay();

const showReceiver = computed(() => {
  // show when enabled
  if (useCustomAddress) return true;

  // receive asset not evm; show
  if (!isEVMAsset(fromAsset)) return true;

  if (connectedWallets.length == 0) return true;
  let evmCount = 0;
  connectedWallets.value.forEach((w) => {
    if (isEvmName(w.name)) evmCount += 1;
  });
  // connected evm == 0; show
  if (evmCount == 0) return true;
  // connected evm == 1; don't show
  if (evmCount == 1) return false;
  // connected evm >= 2; show selector
  if (evmCount >= 2) return true;
});

watch(showReceiver, () => {
  console.log(showReceiver);
});
</script>

<style lang="scss">
@import "assets/css/custom";
.bridge-card {
  width: 540px;
  box-shadow: 0 1px 8px #0003 !important;
}
.mobile-bridge-card {
  width: calc(100vw - 16px);
  box-shadow: 0 1px 8px #0003 !important;
}
</style>