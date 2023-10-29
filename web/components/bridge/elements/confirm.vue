<template>
  <div class="mx-2">
    <v-btn
      v-if="cStore.connected"
      elevation="0"
      class="rounded-xl w-100"
      style="background-color: var(--palette-black-5)"
      size="x-large"
      height="64px"
    >
      <span> {{ $t("bridge") }}</span>
    </v-btn>
    <v-btn
      v-else
      elevation="0"
      height="64px"
      size="x-large"
      class="rounded-xl w-100"
      @click="cStore.mutateDialog(true)"
      @keyup.esc="cStore.mutateDialog(false)"
      style="background: linear-gradient(90deg, var(--palette-main-primary), var(--palette-main-secondary))"
    >
      <span class="h6" style="color: var(--palette-background-1)">
        {{ $t("connect_wallet") }}</span
      >
    </v-btn>

    <ConnectWalletDialog />
  </div>
</template>

<script setup>
import { useConnectStore } from "~/stores/connect/connect";
import ConnectWalletDialog from "~/components/bridge/elements/dialogs/connectwallet.vue";

const cStore = useConnectStore();
</script>

<style>
</style>