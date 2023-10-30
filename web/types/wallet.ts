export type SupportedWallets = 'Metamask' | 'Rabby' | 'WalletConnect' | 'Coinbase' | 'Unisat' | 'Mixin Messenger' | 'Fennec'
export type SupportedChains  = 'Ethereum' | 'Bitcoin' | 'Mixin'

export interface Wallet {
  name: string,
  icon: string,
  loading: boolean,
  connected: boolean
}

export interface ConnectedWallet {
  name: string,
  icon: string,
  chain: string,
  chain_id: string, 
  address: string,
  token?: string
}