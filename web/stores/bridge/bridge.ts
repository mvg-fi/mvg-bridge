import { defineStore } from "pinia";
import { BridgeState } from "~/types/stores";
import assets from "~/assets/constants/miniumlist.json"

export const useBridgeStore = defineStore('bridge', {
  state: () => ({
    fromAsset: assets[0],
    toAsset: assets[1],
    bridgeAmount: undefined,
    receiveAmount: undefined,
    fromDialog: false,
    toDialog: false,
    settingMode: false,
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
    },
    setAmount(from: string, value: number) {
      if (from) {
        this.bridgeAmount = value
        console.log(this.bridgeAmount)
        return
      }
      this.receiveAmount = value
      console.log(this.receiveAmount)
    },
    switchSettingMode() {
      this.settingMode = !this.settingMode
    }
  }
})