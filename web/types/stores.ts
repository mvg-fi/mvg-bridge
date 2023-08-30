import { Asset } from "./asset"

export interface BridgeState {
  fromAsset: Asset
  toAsset: Asset
  bridgeAmount: number
  receiveAmount: number
  fromDialog: Boolean,
  toDialog: Boolean,
}