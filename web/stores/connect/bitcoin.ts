import type { Wallet } from "~/types/wallet";
import us from "~/assets/images/wallets/unisat.png";
import { appendConnected, useConnectStore } from "./connect";
import { BTCUUID, BitcoinChainName, UnisatName } from "~/helpers/constants";

// Bitcoin
export const connectBitcoin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case UnisatName:
      ConnectUnisat(w);
      break;
  }
}

export const defaultBitcoin = [
  { name: UnisatName, icon: us, loading: false, connected: false },
]

export const ConnectUnisat = async (w: Wallet) => {
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
export const PayUnisat = (toAddress: string, satoshis: string) => {
  if (typeof window.unisat !== 'undefined') {
    console.log('UniSat Wallet is installed!');
  }
  try {
    const tx = window.unisat.sendBitcoin(toAddress, satoshis)
    console.log('unisat pay tx:', tx)
  } catch (e) {
    console.log(e)
  }
}

const unisatLastConnected = () => {
  // check if unisat in the list
  // if so, try to connect unisat
  // works = connected, vice versa
}
