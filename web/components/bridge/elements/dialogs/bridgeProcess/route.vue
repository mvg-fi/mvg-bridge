<template>
  <div :class="clsx('d-flex', 'flex-column', 'justify-center')">
    <!-- From -->
    <div class="ml-6 mt-4">
      <span class="h7" style="color: var(--palette-black-50)">
        {{ $t("you_pay") }}
      </span>
    </div>

    <div class="d-flex justify-space-between mx-6 mr-8">
      <span class="h6 mr-3 my-2">
        {{ bStore.bridgeAmount }} {{ bStore.fromAsset.symbol }}
      </span>
      <div class="d-flex align-center">
        <v-icon size="40">
          <v-img :src="bStore.fromAsset.icon" :alt="bStore.fromAsset.name" />
          <v-img
            :src="bStore.fromAsset.chain_icon"
            v-if="bStore.fromAsset.asset_id != bStore.fromAsset.chain_id"
            :class="
              clsx(
                'chain-icon-absolute'
              )
            "
          ></v-img>
        </v-icon>
      </div>
    </div>

    <div class="h7m ml-6 mb-4">
      <span style="color: var(--palette-black-50)" v-if="bStore.bridgeAmount">
        ${{ bStore.fromTotalPrice }}
      </span>
      <span style="opacity: 0" v-else> $ </span>
    </div>

    <!-- To -->
    <div class="ml-6 mt-4">
      <span class="h7" style="color: var(--palette-black-50)">
        {{ $t("you_receive") }}
      </span>
    </div>

    <div class="d-flex justify-space-between mx-6 mr-8">
      <span class="h6 mr-3 my-2">
        {{ bStore.receiveAmount }} {{ bStore.toAsset.symbol }}
      </span>
      <div class="d-flex align-center">
        <v-icon size="40">
          <v-img :src="bStore.toAsset.icon" :alt="bStore.toAsset.name" />
          <v-img
            :src="bStore.toAsset.chain_icon"
            v-if="bStore.toAsset.asset_id != bStore.toAsset.chain_id"
            :class="
              clsx(
                 'chain-icon-absolute'
              )
            "
          ></v-img>
        </v-icon>
      </div>
    </div>

    <div class="h7m ml-6 mb-4">
      <span style="color: var(--palette-black-50)" v-if="bStore.receiveAmount">
        ${{ bStore.toTotalPrice }}
      </span>
      <span style="opacity: 0" v-else> $ </span>
    </div>
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { ArrowRightIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";

const { mobile } = useDisplay()
const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.img-class {
  width: 24px;
  height: 24px;
}
.chain-icon-absolute {
  height: 14px;
  width: 14px;
  position: absolute;
  bottom: 0px;
  left: 0px;
}
</style>