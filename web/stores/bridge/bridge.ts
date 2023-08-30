import { defineStore } from "pinia";
import { BridgeState } from "~/types/stores";
import { Asset } from "~/types/asset"
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
    receiver: '',
    selectedNetwork: undefined,
  } as BridgeState),
  getters: {
  
  },
  actions: {
    switch() {
      const from = this.fromAsset
      this.fromAsset = this.toAsset
      this.toAsset = from
    },
    switchSettingMode() {
      this.settingMode = !this.settingMode
    },
    mutateDialog(from: string, open: boolean) {
      if (from) {
        this.fromDialog = open
        return
      }
      this.toDialog = open
    },
    setAsset(from: string, value: Asset) {
      if (from) {
        this.fromAsset = value
        this.fromDialog = false;
        return
      }
      this.toAsset = value
      this.toDialog = false;
    },
    setAmount(from: string, value: number) {
      if (from) {
        this.bridgeAmount = value
        return
      }
      this.receiveAmount = value
    },
    setSelectedNetwork(value: Asset) { 
      this.selectedNetwork = value
    },
  }
})