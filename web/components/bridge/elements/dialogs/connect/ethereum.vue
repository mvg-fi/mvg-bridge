<template>
  <div class="d-flex flex-column">
    <div class="h7m mb-2">
      <span style="color: var(--palette-black-50)">{{ $t("ethereum") }}</span>
    </div>

    <Single
      :w="w"
      :key="w.name"
      @click="connectEthereum(w)"
      v-for="w in cStore.ethereumWallets"
    />
  </div>
</template>

<script setup>
import { useWeb3ModalAccount } from "@web3modal/ethers5/vue";
import Single from "~/components/bridge/elements/dialogs/connect/single";
import { connectEthereum, useConnectStore } from "~/stores/connect/connect";
const cStore = useConnectStore();

// Check if wc connected
try {
  const { isConnected } = useWeb3ModalAccount();
  if (isConnected) {
    console.log(isConnected);
    cStore.setConnected(true);
    cStore.setConnectedWallets([]);
  }
} catch (e) {
  console.log(e);
}
</script>

<style lang="scss" scoped>
</style>