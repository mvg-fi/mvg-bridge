import type { Asset } from "./asset";

export type Exchanges = {
  fromAsset: Asset,
  toAsset: Asset,
  fromAmount: string,
  toAmount: string,
  completed: boolean,
}