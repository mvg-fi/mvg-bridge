<template>
  <div class="d-flex flex-column network-selector">
    <div class="overflow-y-auto">
      <div class="">
        <v-btn
          icon
          class="network-btn my-1"
          elevation="0"
          v-for="chain in chainList"
          :key="chain.assetid"
        >
          <v-icon style="width: 32px; height: 32px" class="">
            <v-img :src="chain.icon" />
          </v-icon>
        </v-btn>
      </div>
    </div>

    <div class="my-3 mt-1">
      <div style="" class="bg-white">
        <v-btn icon class="network-btn" elevation="0" @click="detailOpened.value=!(detailOpened.value)">
          <v-icon style="width: 16px; height: 16px; opacity: 40">
            <ChevronDoubleLeftIcon v-if="detailOpened.value" />
            <ChevronDoubleRightIcon v-else/>
          </v-icon>
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script setup>
import chainList from "~/assets/constants/miniumchainlist.json";
import { Bars3Icon, ChevronDoubleRightIcon, ChevronDoubleLeftIcon } from "@heroicons/vue/24/outline";

const detailOpened = ref(false);
const selectedNetwork = ref(chainList[0]);
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
</style>