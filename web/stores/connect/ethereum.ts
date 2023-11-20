import a from "~/assets/constants/assetToEvm.json";
import mm from "~/assets/images/wallets/metamask.png";
import rb from "~/assets/images/wallets/rabby.png";
import wc from "~/assets/images/wallets/walletconnect.png";
import cb from "~/assets/images/wallets/coinbase.png";

import type { ethers } from "ethers";
import { toHex } from "~/helpers/utils";
import type { Wallet } from "~/types/wallet";
import { getAddress } from "ethers/lib/utils";
import { appendConnected, useConnectStore } from "./connect";
import { CoinbaseName, ETHUUID, EthereumChainName, MetamaskName, RabbyName, WalletConnectName } from "~/helpers/constants";
import { useWeb3Modal, useWeb3ModalAccount, useWeb3ModalEvents, useWeb3ModalSigner, useWeb3ModalState } from "@web3modal/ethers5/vue";
import type { EVMNetworkParam } from "~/types/network";

// URLs
export const MAINNET_RPC_URL = `https://eth-rpc.gateway.pokt.network`
export const MVM_RPC_URL = "https://geth.mvm.dev/";
export const BSC_RPC_URL = "https://bsc-dataseed1.ninicoin.io";
export const POLYGON_RPC_URL = "https://polygon.llamarpc.com";
export const OPTIMISIM_RPC_URL = "https://mainnet.optimism.io";
export const ARBITRUM_RPC_URL = "https://arb1.arbitrum.io/rpc";
export const FANTOM_RPC_URL = "https://rpc3.fantom.network";

export const ETHER_SCAN_URL = 'https://etherscan.io/'
export const MVM_SCAN_URL = "https://scan.mvm.dev/";
export const BSC_SCAN_URL = "https://bscscan.com/";
export const POLYGON_SCAN_URL = "https://polygonscan.com/";
export const OPTIMISM_SCAN_URL = "https://optimistic.etherscan.io/";
export const ARBITRUM_SCAN_URL = "https://arbiscan.io/";
export const FANTOM_SCAN_URL = "https://ftmscan.com/";

// ChainID
export const MAINNET_CHAIN_ID = 1;
export const BSC_CHAIN_ID = 56;
export const POLYGON_CHAIN_ID = 137;
export const MVM_CHAIN_ID = 73927;
export const OPTIMISM_CHAIN_ID = 10;
export const FANTOM_CHAIN_ID = 250;
export const ARBITRUM_CHAIN_ID = 42161;

export const MAINNET_CHAIN_HEX_ID = toHex(MAINNET_CHAIN_ID);
export const BSC_CHAIN_HEX_ID = toHex(BSC_CHAIN_ID);
export const POLYGON_CHAIN_HEX_ID = toHex(POLYGON_CHAIN_ID);
export const MVM_CHAIN_HEX_ID = toHex(MVM_CHAIN_ID);

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
export const defaultEthereum = [
  { name: MetamaskName, icon: mm, loading: false, connected: false },
  { name: RabbyName, icon: rb, loading: false, connected: false },
  { name: WalletConnectName, icon: wc, loading: false, connected: false },
  { name: CoinbaseName, icon: cb, loading: false, connected: false },
]

export const isEvmName = (name: string) => { return name === MetamaskName || name === RabbyName || name === WalletConnectName || name === CoinbaseName }

export const walletConnect = async (w: Wallet) => {
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
        const { walletProvider } = useWeb3ModalSigner()
        const { address } = useWeb3ModalAccount();
        console.log(walletProvider.value)
        cStore.afterConnect();
        console.log(address);
        appendConnected({
          name: WalletConnectName,
          icon: wc,
          chain: EthereumChainName,
          chain_id: ETHUUID,
          address: address.value || '',
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

export const PayWalletConnect = async (to: string, value: string, chainId: number) => {
  const { signer } = useWeb3ModalSigner()
  if (!signer.value) return
  const params = {
    from: getAddress(await signer.value.getAddress()),
    to,
    value,
    chainId,
  };
  let tx;
  try {
    tx = await signer.value.sendTransaction(params)
  } catch (e) {
    console.log(e)
  }
}

const walletConnectLastConnected = () => {

}

const SwitchNetwork = async (provider: ethers.providers.Web3Provider, network: number) => {
  const request = provider.provider.request;
  if (!request) return new Error('Web3Provider.provider.request must be defined');

  const hex = toHex(network);

  try {
    return await request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: hex }]
    });
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (switchError: any) {
    if (
      !(
        switchError?.code === 4902 ||
        switchError?.code === -32603 ||
        switchError?.data?.orginalError?.code === 4902 ||
        switchError?.data?.orginalError?.code === -32603
      )
    )
      return switchError;
    await request({
      method: 'wallet_addEthereumChain',
      params: [networkParams[hex]]
    });
    await SwitchNetwork(provider, network);
  }
};

export const CheckNetworkCorrect = async (payAssetID: string) => {
  const { walletProvider } = useWeb3ModalSigner()
  const { selectedNetworkId } = useWeb3ModalState()
  if (a[payAssetID] == selectedNetworkId) {
    return
  }
  if (!a[payAssetID]) return
  if (!walletProvider.value) return

  // switch to pay asset network
  SwitchNetwork(walletProvider.value, a[payAssetID])
}

// Network
export const networkParams: Record<string, EVMNetworkParam> = {
  [MAINNET_CHAIN_HEX_ID]: {
    chainId: MAINNET_CHAIN_HEX_ID,
    rpcUrls: [MAINNET_RPC_URL],
    chainName: 'Ethereum Mainnet',
    nativeCurrency: { name: 'Ethereum', decimals: 18, symbol: 'ETH' },
    blockExplorerUrls: [ETHER_SCAN_URL]
  },
  [MVM_CHAIN_HEX_ID]: {
    chainId: MVM_CHAIN_HEX_ID,
    rpcUrls: [MVM_RPC_URL],
    chainName: 'Mixin Virtual Machine',
    nativeCurrency: { name: 'Ethereum', decimals: 18, symbol: 'ETH' },
    blockExplorerUrls: [MVM_SCAN_URL]
  },
  [BSC_CHAIN_HEX_ID]: {
    chainId: BSC_CHAIN_HEX_ID,
    rpcUrls: [BSC_RPC_URL],
    chainName: 'Binance Smart Chain',
    nativeCurrency: { name: 'Binance Coin', decimals: 18, symbol: 'BNB' },
    blockExplorerUrls: [BSC_SCAN_URL]
  },
  [POLYGON_CHAIN_HEX_ID]: {
    chainId: POLYGON_CHAIN_HEX_ID,
    rpcUrls: [POLYGON_RPC_URL],
    chainName: 'Polygon',
    nativeCurrency: { name: 'Matic Token', decimals: 18, symbol: 'MATIC' },
    blockExplorerUrls: [POLYGON_SCAN_URL]
  }
};
