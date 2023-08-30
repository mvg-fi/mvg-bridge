import { Asset } from "./asset"

export interface BridgeState {
  fromAsset: Asset
  toAsset: Asset
  bridgeAmount: number | undefined
  receiveAmount: number | undefined
  fromDialog: Boolean,
  toDialog: Boolean,
}