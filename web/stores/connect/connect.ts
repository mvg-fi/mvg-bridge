import { defineStore } from "pinia";
import type { Wallet } from "~/types/wallet";
import mm from "~/assets/images/wallets/metamask.png";
import rb from "~/assets/images/wallets/rabby.png";
import wc from "~/assets/images/wallets/walletconnect.png";
import cb from "~/assets/images/wallets/coinbase.png";
import us from "~/assets/images/wallets/unisat.png";
import mx from "~/assets/images/wallets/mixin.png";
import fe from "~/assets/images/wallets/fennec.png";
import { useWeb3Modal, useWeb3ModalEvents, useWeb3ModalState } from "@web3modal/ethers5/vue";

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
      { name: "Unisat", icon: us, loading: false, connected: false },
    ],
    mixinWallets: [
      { name: "MixinMessenger", icon: mx, loading: false, connected: false },
      { name: "Fennec", icon: fe, loading: false, connected: false },
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
export const connectEthereum = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case "Metamask":
      walletConnect(w);
      break;
    case "Rabby":
      walletConnect(w);
      break;
    case "WalletConnect":
      walletConnect(w);
      break;
    case "Coinbase":
      walletConnect(w);
      break;
  }
}

// Bitcoin
export const connectBitcoin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case "Unisat":
      unisat(w);
      break;
  }
}

// Mixin
export const connectMixin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case "MixinMessenger":
      mixin(w);
      break;
    case "Fennec":
      fennec(w);
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

export const unisat = async (w: Wallet) => {
  const cStore = useConnectStore()
  if (typeof window.unisat !== 'undefined') {
    console.log('UniSat Wallet is installed!');
  }
  try {
    const result = await window.unisat.requestAccounts()
    if (result.length != 0) {
      w.connected = true
      cStore.afterConnect()
    }
  } catch (e) {
    switch (e.code) {
      case 4001:
        console.log('unisat rejected');
        break;
      default: 
        console.log(e);
    }
  } finally {
    w.loading = false;
  }
}

export const mixin = (w: Wallet) => {

}

export const fennec = (w: Wallet) => {
  
}