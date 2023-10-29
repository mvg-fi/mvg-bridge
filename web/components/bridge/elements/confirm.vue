<template>
  <div class="mx-2">
    <v-btn
      v-if="cStore.connected && !enabled"
      elevation="0"
      class="rounded-xl w-100"
      style="background-color: var(--palette-black-10)"
      size="x-large"
      height="64px"
    >
      <!-- Bridge disabled -->
      <span style="color: var(--palette-black-50)" class="h6">
        {{ $t("bridge") }}
      </span>
    </v-btn>

    <v-btn
      v-else-if="cStore.connected && enabled"
      elevation="0"
      :class="
        clsx(
          'rounded-xl w-100',
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
      size="x-large"
      height="64px"
    >
      <!-- Bridge enabled -->
      <span style="color: var(--palette-background-1)" class="h6">
        {{ $t("bridge") }}</span
      >
    </v-btn>

    <v-btn
      v-else
      elevation="0"
      height="64px"
      size="x-large"
      class="rounded-xl w-100"
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
      <span class="h6" style="color: var(--palette-background-1)">
        {{ $t("connect_wallet") }}</span
      >
    </v-btn>

    <ConnectWalletDialog />
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useConnectStore } from "~/stores/connect/connect";
import ConnectWalletDialog from "~/components/bridge/elements/dialogs/connectwallet.vue";
import { useBridgeStore } from "~/stores/bridge/bridge";

const btnHover = ref(false);
const bStore = useBridgeStore();
const cStore = useConnectStore();
const enabled = computed(() => {
  return (
    bStore.bridgeAmount > 0 && bStore.receiveAmount > 0 && bStore.priceLoaded
  );
});
</script>

<style>
</style>