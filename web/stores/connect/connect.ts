import { defineStore } from "pinia";
import type { ConnectedWallet, Wallet } from "~/types/wallet";
import { CheckNetworkCorrect, PayWalletConnect, defaultEthereum } from "./ethereum";
import { PayUnisat, defaultBitcoin } from "./bitcoin";
import { PayMixinMessenger, defaultMixin } from "./mixin";
import { FennecName, MixinMessengerName, UnisatName, WalletConnectName } from "~/helpers/constants";
import type { Asset } from "~/types/asset";

export * from './bitcoin';
export * from './ethereum';
export * from './mixin';

export const useConnectStore = defineStore('connect', {
  state: () => ({
    connected: false,
    connectDialog: false,
    connectState: 0,      // 0 = connect view, 1 = mixin oauth view, 2 = success, 3 = not detected
    connectedWallets: [],
    connecting: {},

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

// export const connectLastConnected = () => {
//   if (!localStorage.getItem('connected_wallets')) {
//     return
//   }
//   const connected_wallets = localStorage.getItem('connected_wallets');
//   if (connected_wallets == null) {
//     return
//   }
//   const connected = JSON.parse(connected_wallets)
//   connected.forEach(e => {
//     switch (e.name) {
//       case MetamaskName:
//         break;
//       case RabbyName:
//         break;
//       case WalletConnectName:
//         break;
//       case CoinbaseName:
//         break;
//       case UnisatName:
//         break;
//       case MixinMessengerName:
//         break;
//       case FennecName:
//         break;
//       default:
//     }
//   });
// }



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

export const PayWithWallet = (w: Wallet, address: string, amount: string, asset: Asset) => {
  switch (w.name) {
    case WalletConnectName:
      CheckNetworkCorrect(asset.asset_id)
      PayWalletConnect(address, amount, 1)
      break;
    case UnisatName:
      PayUnisat(address, amount)
      break;
    case MixinMessengerName:
      PayMixinMessenger(address, amount)
      break;
    case FennecName:
      break;
    default:
  }
}