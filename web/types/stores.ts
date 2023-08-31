import { Asset } from "./asset"

export interface BridgeState {
  fromAsset: Asset
  toAsset: Asset
  bridgeAmount: number | undefined
  receiveAmount: number | undefined
  fromDialog: boolean,
  toDialog: boolean,
  settingMode: boolean,
  receiver: string,
  selectedNetwork: Asset | undefined,
  searchAsset: string,
  selectNetworkBar: boolean
  // filteredItems: Asset[],
  // filteredChains: any,
}