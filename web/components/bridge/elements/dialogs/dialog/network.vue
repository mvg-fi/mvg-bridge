<template>
  <v-card
    class="d-flex justify-center network-selector h-100 px-1"
    elevation="0"
    style="position: relative"
    v-if="store.selectNetworkBar"
  >
    <div class="overflow-y-auto" style="width: 32px">
      <v-btn
        icon
        elevation="0"
        :key="chain.asset_id"
        style="position: relative"
        :class="clsx('network-btn my-1')"
        @click="setNetwork(chain)"
        v-for="chain in store.filteredChains"
      >
        <v-icon style="width: 32px; height: 32px">
          <v-img :src="chain.icon" :alt="chain.symbol" />
          <v-tooltip activator="parent" location="left" open-delay="50ms">
            {{ $t('only_show_network', {name: chain.name}) }}
          </v-tooltip>
        </v-icon>
      </v-btn>
    </div>
    <div style="position: absolute; right: 0px; top: 45%">
      <v-icon
        style="width: 10px; height: 32px; color: grey; opacity: 80%"
        @click.stop="toggle"
      >
        <ChevronLeftIcon />
      </v-icon>
    </div>
  </v-card>

  <div
    class="d-flex justify-center h-100"
    style="position: relative"
    elevation="0"
    v-else
  >
    <div style="position: absolute; left: 0px; top: 45%; z-index: 100">
      <v-icon
        style="width: 10px; height: 32px; color: grey; opacity: 80%"
        @click.stop="toggle"
      >
        <ChevronRightIcon />
      </v-icon>
    </div>
  </div>
</template>

<script setup>
import {
  Bars3Icon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
import clsx from "clsx";

const store = useBridgeStore();
const setNetwork = (network) => {
  // If select == current, change to undefined
  // If undefined, change
  // If select != current, change
  if (store.selectedNetwork == undefined) {
    store.setSelectedNetwork(network);
    return;
  }
  if (network.asset_id === store.selectedNetwork.asset_id) {
    store.setSelectedNetwork(undefined);
    return;
  }
  if (network != store.selectedNetwork) {
    store.setSelectedNetwork(network);
    return;
  }
};
const toggle = () => {
  store.mutateSelectNetworkBar(!store.selectNetworkBar);
};
</script>

<style lang="scss" scoped>
.network-btn {
  width: 36px;
  height: 36px;
}
.network-selector {
  width: 72px;
}
.v-tooltip :deep(.v-overlay__content) {
  font-size: 10px !important;
  opacity: 90%;
  padding: 4px;
  padding-left: 8px;
  padding-right: 8px;
}

* {
  scrollbar-width: none;
}
::-webkit-scrollbar {
  display: none;
}
</style>