import type { ConnectedWallet, Wallet } from "~/types/wallet";
import { userMe } from "~/helpers/mixin/user";
import mx from "~/assets/images/wallets/mixin.png";
import fe from "~/assets/images/wallets/fennec.png";
import { appendConnected, useConnectStore } from "./connect";
import { FennecName, MixinChainName, MixinMessengerName } from "~/helpers/constants";

export const defaultMixin = [
  { name: MixinMessengerName, icon: mx, loading: false, connected: false },
  { name: FennecName, icon: fe, loading: false, connected: false },
]

// Mixin
export const connectMixin = (w: Wallet) => {
  w.loading = true;
  switch (w.name) {
    case MixinMessengerName:
      ConnectMixinMessenger(w);
      break;
    case FennecName:
      ConnectFennec(w);
      break;
  }
  w.loading = false;
}

export const ConnectMixinMessenger = (w: Wallet) => {
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

export const ConnectFennec = (w: Wallet) => {

}

export const PayMixinMessenger = (address: string, amount: string) => {
  // generate payment link
}

// Last connected

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