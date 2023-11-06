<template>
  <div class="d-flex align-center justify-center flex-row" style="gap: 24px">
    <v-btn
      size="large"
      class="bg-background rounded-pill"
      elevation="0"
      style="border: 1px solid var(--palette-black-10)"
      @click="viewAddress"
    >
      <span class="h7 text-black-darken-5">
        {{ $t("view_address") }}
      </span>
    </v-btn>

    <v-btn class="confirm-btn rounded-pill" size="large" @click="payWithWallet">
      <span class="h7 text-background">
        {{ $t("pay_with_wallet") }}
      </span>
    </v-btn>
  </div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";

const bStore = useBridgeStore();
const viewAddress = () => {
  bStore.setBridgeProcessState(1);
};
const payWithWallet = () => {
  bStore.setBridgeProcessState(2);
};
</script>

<style lang="scss">
.confirm-btn {
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
  text-transform: capitalize;
}
</style>