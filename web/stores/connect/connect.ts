import { defineStore } from "pinia";
import type { Wallet } from "~/types/wallet";
import mm from "~/assets/images/wallets/metamask.png";
import rb from "~/assets/images/wallets/rabby.png";
import wc from "~/assets/images/wallets/walletconnect.png";
import cb from "~/assets/images/wallets/coinbase.png";
import us from "~/assets/images/wallets/unisat.png";
import mx from "~/assets/images/wallets/mixin.png";
import fe from "~/assets/images/wallets/fennec.png";
import { useWeb3Modal, useWeb3ModalAccount, useWeb3ModalEvents, useWeb3ModalState } from "@web3modal/ethers5/vue";

export const useConnectStore = defineStore('connect', {
  state: () => ({
    connected: false,
    connectDialog: false,

    ethereumWallets: [
      { name: "Metamask", icon: mm, loading: false, connected: false },
      { name: "Rabby", icon: rb, loading: false, connected: false },
      { name: "WalletConnect", icon: wc, loading: false, connected: false },
      { name: "Coinbase", icon: cb, loading: false, connected: false },
    ],
    bitcoinWallets: [
      { name: "Unisat", icon: us },
    ],
    mixinWallets: [
      { name: "MixinMessenger", icon: mx },
      { name: "Fennec", icon: fe },
    ],

    connectedWallets: [],
  }),
  getters: {
  },
  actions: {
    setConnected(c: boolean) {
      this.connected = c
    },
    setConnectedWallets(ws: []) {
      this.connectedWallets = ws
    },
    mutateDialog(open: boolean) {
      this.connectDialog = open
    },
    afterConnect() {
      this.connectDialog = false
    },
    afterDisconnect() {
      this.connectedWallets = []
    }
  }
})

// Ethereum
export const connectEthereum = (w:Wallet) => {
  w.loading = true;
  switch (w.name) {
    case "Metamask":
      break;
    case "Rabby":
      break;
    case "WalletConnect":
      walletConnect(w);
      break;
    case "Coinbase":
      break;
  }
}

const walletConnect = async (w: Wallet) => {
  const cStore = useConnectStore()
  const { open } = useWeb3Modal();
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
  watchEffect(() => {
    switch (events.data.event) {
      case "CONNECT_SUCCESS":
        cStore.setConnected(true);
        cStore.afterConnect();
      case "DISCONNECT_SUCCESS":
        cStore.setConnected(false);
        cStore.afterDisconnect();
      default:
        console.log(events.data.event);
    }
  });
};

// Bitcoin


// Mixin