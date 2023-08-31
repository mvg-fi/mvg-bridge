<template>
  <v-card
    class="d-flex justify-center network-selector h-100 px-1"
    elevation="0"
    style="position: relative"
    v-if="barOpen"
  >
    <div class="overflow-y-auto" style="width: 32px">
      <v-btn
        icon
        class="network-btn my-1"
        elevation="0"
        v-for="chain in store.filteredChains"
        :key="chain.asset_id"
        @click="setNetwork(chain)"
      >
        <v-icon style="width: 32px; height: 32px" class="">
          <v-img :src="chain.icon" :alt="chain.symbol" />
        </v-icon>
      </v-btn>
    </div>
    <div style="position: absolute; right: 0px; top: 45%">
      <v-icon style="width: 10px; height: 32px; color: grey; opacity: 80%;" @click.stop="toggle">
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
    <div style="position: absolute; left: 0px; top: 45%; z-index: 100;">
      <v-icon style="width: 10px; height: 32px; color: grey; opacity: 80%;" @click.stop="toggle">
        <ChevronRightIcon />
      </v-icon>
    </div>
  </div>
</template>

<script setup>
import { Bars3Icon, ChevronLeftIcon, ChevronRightIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();

const barOpen = ref(false);
const setNetwork = (network) => {
  store.setSelectedNetwork(network);
};
const toggle = () => {
  barOpen.value = !barOpen.value;
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
* {
  scrollbar-width: none;
}
::-webkit-scrollbar {
  display: none;
}
</style>