import { defineStore } from "pinia";
import { BridgeState } from "~/types/stores";
import assets from "~/assets/constants/miniumlist.json"

export const useBridgeStore = defineStore('bridge', {
  state: () => ({
    fromAsset: assets[0],
    toAsset: assets[1],
    bridgeAmount: 0,
    receiveAmount: 0,
    fromDialog: false,
    toDialog: false,
  } as BridgeState),
  getters: {
  
  },
  actions: {
    switch() {
      const from = this.fromAsset
      this.fromAsset = this.toAsset
      this.toAsset = from
    },
    mutateDialog(from: string, open: boolean) {
      if (from) {
        this.fromDialog = open
        return
      }
      this.toDialog = open
    }
  }
})