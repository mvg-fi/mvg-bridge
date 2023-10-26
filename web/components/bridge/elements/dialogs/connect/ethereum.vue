<template>
  <div class="d-flex flex-column">
    <div class="h7m mb-2">
      <span style="color: var(--palette-black-50)">{{ $t("ethereum") }}</span>
    </div>

    <button
      v-for="w in ethereumWallets"
      :key="w.name"
      class="d-flex flex-row align-center mb-3"
      style="height: 40px; gap: 8px"
      @click="connect(w.name)"
    >
      <v-img
        class="flex-grow-0 rounded-circle"
        :src="w.icon"
        style="width: 32px; height: 32px"
      />

      <span
        class="flex-grow-1 justify-start text-start"
        style="font-size: 16px; line-height: 32px; font-family: Satoshi-SemiBold"
      >
        {{ w.name }}
      </span>

      <div v-if="false">
        <!-- If connected -->
      </div>

      <div v-if="false">
        <!-- If loading -->
      </div>
    </button>
  </div>
</template>

<script setup>
import mm from "~/assets/images/wallets/metamask.png";
import rb from "~/assets/images/wallets/rabby.png";
import wc from "~/assets/images/wallets/walletconnect.png";
import cb from "~/assets/images/wallets/coinbase.png";
import { useWeb3Modal } from '@web3modal/ethers5/vue'
import { useConnectStore } from "~/stores/connect/connect";

const store = useConnectStore()
const ethereumWallets = [
  { name: "Metamask", icon: mm },
  { name: "Rabby", icon: rb },
  { name: "WalletConnect", icon: wc },
  { name: "Coinbase", icon: cb },
];

const connect = (w) => {
  switch (w) {
    case "Metamask":
      metamask()
      break;
    case "Rabby":
      break;
    case "WalletConnect":
      walletconnect();
      break;
    case "Coinbase":
      break;
  }
};

const metamask = () => {
  if (typeof window.ethereum !== "undefined") {
    console.log("MetaMask is installed!");
  }
};
const rabby = () => {};
const walletconnect = async () => {
  const modal = useWeb3Modal()
  const state = await modal.open()
  console.log(state)
};
const coinbase = () => {};
</script>

<style lang="scss" scoped>

</style>