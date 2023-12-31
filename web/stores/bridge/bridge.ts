import { defineStore } from "pinia";
import { BridgeState } from "~/types/stores";
import { Asset } from "~/types/asset"
import assets from "~/assets/constants/miniumlist.json"
import chains from "~/assets/constants/miniumchainlist.json"
import { networkAsset } from "~/helpers/mixin/asset";
import { BN2 } from "~/helpers/bignumber/bn";

export const useBridgeStore = defineStore('bridge', {
  state: () => ({
    fromAsset: assets[0],
    toAsset: assets[1],
    bridgeUsdPrice: undefined,
    bridgeAmount: undefined,
    receiveUsdPrice: undefined,
    receiveAmount: undefined,
    fromDialog: false,
    toDialog: false,
    settingMode: false,
    receiver: '',
    selectedNetwork: undefined,
    searchAsset: '',
    selectNetworkBar: false,
    priceLoading: false,
    priceLoaded: true,
    costUSD: 0,

    bridgeProcessDialog: false,
    menuState: 0,               // 0 == Default, 1 == Language select
    recentCardState: 0,         // 0 == Not shown, 1 == show, 2 == extended
    bridgeProcessState: 0,      // 0 == opened, 1 == show address, 2 == select wallet, 3 == paid
  } as BridgeState),
  getters: {
    filteredItems: (state) => {
      // if (state.selectedNetwork != undefined) {
      //   return assets.filter((asset) => {
      //     if (asset.chain_id != state.selectedNetwork?.asset_id) return false
      //     return (
      //       asset.symbol.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //       asset.name.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //       asset.chain_name?.toUpperCase().match(state.searchAsset.toUpperCase())
      //     );
      //   })
      // }
      // return assets.filter((asset) => {
      //   return (
      //     asset.symbol.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //     asset.name.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //     asset.chain_name?.toUpperCase().match(state.searchAsset.toUpperCase())
      //   );
      // })
      const searchQuery = state.searchAsset.toUpperCase();
      const selectedNetworkAssetId = state.selectedNetwork?.asset_id;

      return assets.filter((asset) => {
        if (selectedNetworkAssetId !== undefined && asset.chain_id !== selectedNetworkAssetId) {
          return false;
        }

        const assetName = asset.name.toUpperCase();
        const assetSymbol = asset.symbol.toUpperCase();
        const chainName = asset.chain_name?.toUpperCase();

        return assetSymbol.includes(searchQuery) || assetName.includes(searchQuery) || (chainName && chainName.includes(searchQuery));
      });
    },
    filteredChains: (state) => {
      // const assetList = assets.filter((asset) => {
      //   return (
      //     asset.symbol.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //     asset.name.toUpperCase().match(state.searchAsset.toUpperCase()) ||
      //     asset.chain_name?.toUpperCase().match(state.searchAsset.toUpperCase())
      //   )
      // })
      // return chains.filter(chain => {
      //   return assetList.some(asset => chain.asset_id === asset.chain_id);
      // });

      const searchQuery = state.searchAsset.toUpperCase();
      const assetList = assets.filter((asset) => {
        const assetName = asset.name.toUpperCase();
        const assetSymbol = asset.symbol.toUpperCase();
        const chainName = asset.chain_name?.toUpperCase();

        return assetSymbol.includes(searchQuery) || assetName.includes(searchQuery) || (chainName && chainName.includes(searchQuery));
      });

      const chainAssetIds = assetList.map((asset) => asset.chain_id);
      return chains.filter((chain) => chainAssetIds.includes(chain.asset_id));
    },
    fromTotalPrice: (state) => {
      return BN2(state.bridgeUsdPrice).multipliedBy(BN2(state.bridgeAmount)).toFixed(2)
    },
    toTotalPrice: (state) => {
      return BN2(state.receiveUsdPrice).multipliedBy(BN2(state.receiveAmount)).toFixed(2)
    },
    confirmEnabled: (state) => {
      return state.bridgeAmount > 0 && state.receiveAmount > 0;
    }
  },
  actions: {
    switch() {
      const from = this.fromAsset
      const fromUsd = this.bridgeUsdPrice
      this.fromAsset = this.toAsset
      this.bridgeUsdPrice = this.receiveUsdPrice
      this.toAsset = from
      this.receiveUsdPrice = fromUsd
    },
    switchSettingMode() {
      this.settingMode = !this.settingMode
    },
    mutateDialog(from: boolean, open: boolean) {
      if (from) {
        this.fromDialog = open
        this.cleanUpStates()
        console.log(`mutateDialog(${from}, ${open})`)
        return
      }
      this.toDialog = open
      this.cleanUpStates()
      console.log(`mutateDialog(${from}, ${open})`)
    },
    mutateBridgeProcess(open: boolean) {
      this.bridgeProcessDialog = open
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
        this.fetchUSD(true)
        return
      }
      this.toAsset = value
      this.toDialog = false;
      this.fetchUSD(false)
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
    setMenuState(v: number) {
      this.menuState = v
    },
    setRecentCardState(v: number) {
      this.recentCardState = v
    },
    setBridgeProcessState(v: number) {
      this.bridgeProcessState = v
    },
    clearReceiver() {
      this.receiver = ""
    },
    cleanUpStates() {
      this.searchAsset = ''
      this.selectNetworkBar = false
      this.selectedNetwork = undefined
    },
    async fetchUSD(from: boolean) {
      if (from) this.bridgeUsdPrice = await fetchUSDPrice(this.fromAsset);
      else this.receiveUsdPrice = await fetchUSDPrice(this.toAsset);
    }
  }
})

export const fetchUSDPrice = async (a: Asset) => {
  const aa = await networkAsset(a)
  return aa.price_usd || 0
}