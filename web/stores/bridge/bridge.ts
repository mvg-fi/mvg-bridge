import { defineStore } from "pinia";
import { BridgeState } from "~/types/stores";
import { Asset } from "~/types/asset"
import assets from "~/assets/constants/miniumlist.json"
import chains from "~/assets/constants/miniumchainlist.json"

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
    searchAsset: '',
    selectNetworkBar: false,
  } as BridgeState),
  getters: {
    filteredItems: (state) => {
      return assets.filter((asset) => {
        return (
          asset.symbol.toUpperCase().match(state.searchAsset.toUpperCase()) ||
          asset.name.toUpperCase().match(state.searchAsset.toUpperCase())
          // asset.chainName?.toUpperCase().match(state.searchAsset.toUpperCase()) ||
          // asset.chainSymbol?.toUpperCase().match(state.searchAsset.toUpperCase())
        );
      })
    },
    filteredChains: (state) => {
      let filteredChain: Asset[] = chains.filter((chain) => {
        return (
          chain.symbol.toUpperCase().match(state.searchAsset.toUpperCase()) ||
          chain.name.toUpperCase().match(state.searchAsset.toUpperCase())
        )
      })
      const itemsChains: Asset[] = state.filteredItems.filter((asset: Asset) => {
        return asset.asset_id === asset.chain_id
      })
      // itemsChains.forEach(element => {
      //   if (!(element.asset_id in filteredChain)) filteredChain.push(element)
      // });
      return filteredChain;
    }
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
    mutateDialog(from: boolean, open: boolean) {
      if (from) {
        this.fromDialog = open
        console.log(`mutateDialog(${from}, ${open})`)
        return
      }
      console.log(`mutateDialog(${from}, ${open})`)
      this.toDialog = open
    },
    mutateSelectNetworkBar(open: boolean) {
      this.selectNetworkBar = open
      console.log(`mutateSelectNetworkBar(${open})`)
    },
    setAsset(from: boolean, value: Asset) {
      if (from) {
        this.fromAsset = value
        this.fromDialog = false;
        console.log(`setAsset(${from}, ${value.symbol})`)
        return
      }
      this.toAsset = value
      this.toDialog = false;
      console.log(`setAsset(${from}, ${value.symbol})`)
    },
    setAmount(from: boolean, value: number) {
      if (from) {
        this.bridgeAmount = value
        return
      }
      this.receiveAmount = value
    },
    setSearchAsset(value: string) {
      this.searchAsset = value
      console.log('setSearchAsset:', value)
    },
    setSelectedNetwork(value: Asset) {
      this.selectedNetwork = value
      console.log('setSelectedNetwork:', this.selectedNetwork)
    },
  }
})