import { defineStore } from "pinia";
import type { ConnectedWallet, Wallet } from "~/types/wallet";
import mm from "~/assets/images/wallets/metamask.png";
import rb from "~/assets/images/wallets/rabby.png";
import wc from "~/assets/images/wallets/walletconnect.png";
import cb from "~/assets/images/wallets/coinbase.png";
import us from "~/assets/images/wallets/unisat.png";
import mx from "~/assets/images/wallets/mixin.png";
import fe from "~/assets/images/wallets/fennec.png";
import { userMe } from "~/helpers/mixin/user";
import { useWeb3Modal, useWeb3ModalEvents, useWeb3ModalState } from "@web3modal/ethers5/vue";
import { BTCUUID, BitcoinChainName, CoinbaseName, ETHUUID, EthereumChainName, FennecName, MetamaskName, MixinChainName, MixinMessengerName, RabbyName, UnisatName, WalletConnectName } from "~/helpers/constants";

const defaultEthereum = [
  { name: MetamaskName, icon: mm, loading: false, connected: false },
  { name: RabbyName, icon: rb, loading: false, connected: false },
  { name: WalletConnectName, icon: wc, loading: false, connected: false },
  { name: CoinbaseName, icon: cb, loading: false, connected: false },
]
const defaultBitcoin = [
  { name: UnisatName, icon: us, loading: false, connected: false },
]
const defaultMixin = [
  { name: MixinMessengerName, icon: mx, loading: false, connected: false },
  { name: FennecName, icon: fe, loading: false, connected: false },
]

export const useConnectStore = defineStore('connect', {
  state: () => ({
    connected: false,
    connectDialog: false,
    connectState: 3,      // 0 = connect view, 1 = mixin oauth view, 2 = success, 3 = failed/canceled
    connectedWallets: [],

    ethereumWallets: defaultEthereum,
    bitcoinWallets: defaultBitcoin,
    mixinWallets: defaultMixin,
  }),
  getters: {
  },
  actions: {
    setConnected(c: boolean) {
      this.connected = c
    },
    mutateDialog(open: boolean) {
      this.connectDialog = open
      if (!open) this.clearStates()
    },
    afterConnect() {
      this.connectDialog = false
      this.setConnected(true)
      this.clearStates()
    },
    afterDisconnect() {
      this.setConnected(false)
      this.clearStates()
    },
    disconnectAll() {
      this.connectedWallets = []
      this.afterDisconnect()
      localStorage.removeItem("connected_wallets")
    },
    disconnectSpecific(c: ConnectedWallet) {
      const after = this.connectedWallets.filter((e: ConnectedWallet) => { return e.name != c.name })
      this.connectedWallets = after
      localStorage.setItem("connected_wallets", JSON.stringify(after))
    },
    connectMore() {
      this.connectDialog = true;
    },
    setConnectState(e: number) {
      this.connectState = e
    },
    clearStates() {
      setTimeout(() => {
        this.connectState = 0
        this.ethereumWallets = defaultEthereum
        this.bitcoinWallets = defaultBitcoin
        this.mixinWallets = defaultMixin
      }, 200);
    }
  }
})

// Ethereum
export const connectEthereum = (w: Wallet) => {
  // This was done within specific component
  w.loading = true;
  switch (w.name) {
    case MetamaskName:
      walletConnect(w);
      break;
    case RabbyName:
      walletConnect(w);
      break;
    case WalletConnectName:
      walletConnect(w);
      break;
    case CoinbaseName:
      walletConnect(w);
      break;
  }
}

// Bitcoin
export const connectBitcoin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case UnisatName:
      unisat(w);
      break;
  }
}

// Mixin
export const connectMixin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case MixinMessengerName:
      mixin(w);
      break;
    case FennecName:
      fennec(w);
      break;
  }
  w.loading = false;
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
  watchEffect(async () => {
    switch (events.data.event) {
      case "CONNECT_SUCCESS":
        cStore.afterConnect();
        appendConnected({
          name: WalletConnectName,
          icon: wc,
          chain: EthereumChainName,
          chain_id: ETHUUID,
          address: '' //await signer?.getAddress(),
        })
        break;
      case "DISCONNECT_SUCCESS":
        cStore.afterDisconnect();
        break;
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
      appendConnected({
        name: UnisatName,
        icon: us,
        chain: BitcoinChainName,
        chain_id: BTCUUID,
        address: result[0] || '',
      })
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
  const cStore = useConnectStore()
  cStore.setConnectState(1)
}
export const mixinOauthSuccess = async (token: string) => {
  const cStore = useConnectStore()
  const me = await userMe(token)
  console.log(me)

  appendConnected({
    name: MixinMessengerName,
    icon: mx,
    chain: MixinChainName,
    chain_id: '',
    address: me.identity_number || '',
    token: token,
  })

  cStore.setConnectState(2);
  setTimeout(() => {
    cStore.afterConnect();
  }, 1000);
}

export const fennec = (w: Wallet) => {

}

export const connectLastConnected = () => {
  if (!localStorage.getItem('connected_wallets')) {
    return
  }
  const connected_wallets = localStorage.getItem('connected_wallets');
  if (connected_wallets == null) {
    return
  }
  const connected = JSON.parse(connected_wallets)
  connected.forEach(e => {
    switch (e.name) {
      case MetamaskName:
        break;
      case RabbyName:
        break;
      case WalletConnectName:
        break;
      case CoinbaseName:
        break;
      case UnisatName:
        break;
      case MixinMessengerName:
        break;
      case FennecName:
        break;
      default:
    }
  });
}
const walletConnectLastConnected = () => {

}
const unisatLastConnected = () => {
  // check if unisat in the list
  // if so, try to connect unisat
  // works = connected, vice versa
}
const mixinMessengerLastConnected = async (c: ConnectedWallet) => {
  // check if mixin messenger in the list
  // if so, check if token works
  // works = connected, vice versa

  if (!c.token) {
    return false
  }
  const me = await userMe(c.token)
  if (!me) {
    return false
  }
  return true
}
const fennecLastConnected = () => {

}
export const appendConnected = (c: ConnectedWallet) => {
  // 0. Add if localStorage doesn't exist
  // 1. Remove and append if old exist
  // 3. Append if old doesn't exist
  const cStore = useConnectStore()
  let state;
  if (!localStorage.getItem('connected_wallets')) {
    localStorage.setItem('connected_wallets', JSON.stringify([c]))
    cStore.connectedWallets = [c]
    return;
  }

  // @ts-ignore
  const old = JSON.parse(localStorage.getItem('connected_wallets'))
  // console.log('old:', old)
  const oldd = old.filter((e) => { return e.name != c.name });
  // console.log('oldd:', oldd)
  const neww = oldd.concat(c)
  // console.log('new_appended:', neww)
  localStorage.setItem("connected_wallets", JSON.stringify(neww))
  cStore.connectedWallets = neww;
}

const connectToLastConnected = () => {
  // 1. Get localStorage, if not exist, return
  // 2. Exist, create state with value
  if (!localStorage.getItem("connected_wallets")) return;
  const connectedWallets = localStorage.getItem("connected_wallets");
}