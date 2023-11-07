<template>
  <div class="mx-2">
    <v-btn
      v-if="cStore.connected && !bStore.confirmEnabled"
      elevation="0"
      class="rounded-xl w-100"
      style="background-color: var(--palette-black-10)"
      size="x-large"
      :height="mobile ? '56px': '64px'"
    >
      <!-- Bridge disabled -->
      <span style="color: var(--palette-black-50)" :class="clsx(mobile? 'h7':'h6')">
        {{ $t("bridge") }}
      </span>
    </v-btn>

    <v-btn
      v-else-if="cStore.connected && bStore.confirmEnabled"
      :key="bStore.confirmEnabled"
      elevation="0"
      :class="
        clsx(
          'rounded-xl w-100 text-background',
          btnHover.value ? 'btn-hover elevation-4' : 'elevation-1'
        )
      "
      @mouseover="btnHover = true"
      @mouseleave="btnHover = false"
      style="
        background: linear-gradient(
          90deg,
          var(--palette-main-primary),
          var(--palette-main-secondary)
        );
      "
      @click="bStore.mutateBridgeProcess(true)"
      size="x-large"
      :height="mobile ? '56px': '64px'"
    >
      <!-- Bridge enabled -->
      <span :class="clsx(mobile? 'h7':'h6')">
        {{ $t("bridge") }}</span
      >
    </v-btn>

    <v-btn
      v-else
      elevation="0"
      size="x-large"
      :height="mobile ? '56px': '64px'"
      class="rounded-xl w-100 text-background"
      @click="cStore.mutateDialog(true)"
      @keyup.esc="cStore.mutateDialog(false)"
      style="
        background: linear-gradient(
          90deg,
          var(--palette-main-primary),
          var(--palette-main-secondary)
        );
      "
    >
      <span :class="clsx(mobile? 'h7':'h6')">
        {{ $t("connect_wallet") }}</span
      >
    </v-btn>

    <ConnectWalletDialog />
    <BridgeProcessDialog />
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useConnectStore } from "~/stores/connect/connect";
import { useBridgeStore } from "~/stores/bridge/bridge";
import ConnectWalletDialog from "~/components/bridge/elements/dialogs/connectwallet.vue";
import BridgeProcessDialog from "~/components/bridge/elements/dialogs/bridgeProcess.vue";

const btnHover = ref(false);
const { mobile } = useDisplay();
const bStore = useBridgeStore();
const cStore = useConnectStore();
// const enabled = computed(() => {
//   return bStore.bridgeAmount > 0 && bStore.receiveAmount > 0; // && bStore.priceLoaded
// });
</script>

<style>
</style>