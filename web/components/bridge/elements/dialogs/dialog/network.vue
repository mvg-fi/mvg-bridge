<template>
  <div class="d-flex flex-column network-selector">
    <v-divider />
    <div class="overflow-y-auto">
      <div>
        <v-btn
          icon
          class="network-btn my-1"
          elevation="0"
          v-for="chain in chainList"
          :key="chain.asset_id"
          @click="setNetwork(chain)"
        >
          <v-icon style="width: 32px; height: 32px" class="">
            <v-img :src="chain.icon" />
          </v-icon>
        </v-btn>
      </div>
    </div>
    <v-divider />

    <div class="mt-1 pb-3 d-flex justify-center">
      <div class="bg-white h-100">
        <v-btn icon style="width: 32px; height: 32px" elevation="0" @click="console.log('click expend')">
          <v-icon style="width: 12px; height: 12px; opacity: 40">
            <ChevronDoubleDownIcon />
          </v-icon>
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script setup>
import chainList from "~/assets/constants/miniumchainlist.json";
import { Bars3Icon, ChevronDoubleDownIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";

const store = useBridgeStore();
const setNetwork = (network) => {
  store.setSelectedNetwork(network);
};
</script>

<style lang="scss" scoped>
@import "assets/css/custom";
$height: calc(#{$select-asset-dialog-height} - 64px - 52px - 20px);
.network-btn {
  width: 36px;
  height: 36px;
}
.network-selector {
  width: 72px;
  min-height: $height;
  padding-bottom: 100px;
}
* {
  scrollbar-width: none;
}
::-webkit-scrollbar {
  display: none;
}
</style>