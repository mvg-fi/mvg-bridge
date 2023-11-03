<template>
  <div class="d-flex flex-column">
    <div class="h7m mb-2 px-3">
      <span style="color: var(--palette-black-50)">{{ $t("ethereum") }}</span>
    </div>

    <Single
      :w="w"
      :key="w.name"
      class="px-3 b rounded-xl"
      @click="walletConnect(w)"
      v-for="w in cStore.ethereumWallets"
    />
  </div>
</template>

<script lang="ts" setup>
import wc from "~/assets/images/wallets/walletconnect.png";
import Single from "~/components/bridge/elements/dialogs/connect/single.vue";
import {
  ETHUUID,
  EthereumChainName,
  WalletConnectName,
} from "~/helpers/constants";
import {
  useWeb3Modal,
  useWeb3ModalEvents,
  useWeb3ModalState,
} from "@web3modal/ethers5/vue";
import {
  appendConnected,
  connectEthereum,
  useConnectStore,
} from "~/stores/connect/connect";
import type { Wallet } from "~/types/wallet";
const cStore = useConnectStore();
const { open } = useWeb3Modal();

const walletConnect = async (w: Wallet) => {
  await open();

  // Cancel loading when modal is off
  const states = useWeb3ModalState();
  watchEffect(() => {
    if (!states.open) {
      w.loading = false;
    }
  });

  // Handle Web3Modal events
  const events = useWeb3ModalEvents();
  watchEffect(async () => {
    switch (events.data.event) {
      case "CONNECT_SUCCESS":
        cStore.afterConnect();
        appendConnected({
          name: WalletConnectName,
          icon: wc,
          chain: EthereumChainName,
          chain_id: ETHUUID,
          address: "", //await signer?.getAddress(),
        });
        break;
      case "DISCONNECT_SUCCESS":
        cStore.afterDisconnect();
        break;
      default:
        console.log(events.data.event);
    }
  });
};
</script>

<style lang="scss" scoped>
.b:hover {
  background-color: var(--palette-black-5);
}
</style>