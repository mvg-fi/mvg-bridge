<template>
  <!-- State 1 shrink -->
  <v-sheet
    v-if="bStore.recentCardState == 1"
    elevation="3"
    class="recent-sheet-shrink px-4 bg-background"
  >
    <!-- Title -->
    <div class="d-flex justify-space-between align-center mt-6 mx-2">
      <span class="h7" style="font-size: 16px">
        {{ $t("recent_exchanges") }}
      </span>
      <div>
        <button
          class="bg-background px-0"
          @click="bStore.setRecentCardState(2)"
        >
          <v-icon size="20">
            <ChevronUpIcon />
          </v-icon>
        </button>
        <button
          class="bg-background px-0 ml-2"
          @click="bStore.setRecentCardState(0)"
        >
          <v-icon size="20">
            <XMarkIcon />
          </v-icon>
        </button>
      </div>
    </div>
  </v-sheet>

  <!-- State 2 extended -->
  <v-sheet
    v-if="bStore.recentCardState == 2"
    elevation="3"
    class="recent-sheet-extended px-4 bg-background overflow-y-auto"
  >
    <!-- Title -->
    <div class="d-flex justify-space-between align-center mt-6 mx-2">
      <span class="h7" style="font-size: 16px">
        {{ $t("recent_exchanges") }}
      </span>
      <button class="bg-background px-0" @click="bStore.setRecentCardState(1)">
        <v-icon size="20">
          <ChevronDownIcon />
        </v-icon>
      </button>
    </div>

    <!-- Items -->
    <div class="d-flex flex-column mt-5 mb-3">
      <!-- Item -->
      <div
        class="my-1 d-flex flex-row align-center"
        style="
          height: 64px;
          border: 1px solid var(--palette-black-10);
          border-radius: 32px;
        "
        v-for="(e, i) in example"
        :key="i"
      >
        <!-- Icon -->
        <div class="ml-4 mb-4" style="width: 24px; height: 24px">
          <v-icon
            style="
              height: 24px;
              width: 24px;
              z-index: 2;
            "
            class="mr-4"
          >
            <v-img :src="e.fromAsset.icon" />
            <v-img
              :src="e.fromAsset.chain_icon"
              v-if="e.fromAsset.asset_id != e.fromAsset.chain_id"
              style="
                height: 8px;
                width: 8px;
                position: absolute;
                bottom: 0px;
                left: 0px;
              "
            ></v-img>
          </v-icon>

          <!-- <v-img :src="e.fromAsset.icon" style="z-index: 2" /> -->

          <v-icon
            style="
              height: 24px;
              width: 24px;
              position: relative;
              bottom: 12px;
              left: 12px;
            "
            class="mr-4"
          >
            <v-img :src="e.toAsset.icon" />
            <v-img
              :src="e.toAsset.chain_icon"
              v-if="e.toAsset.asset_id != e.toAsset.chain_id"
              style="
                height: 8px;
                width: 8px;
                position: absolute;
                bottom: 0px;
                left: 0px;
              "
            ></v-img>
          </v-icon>
        </div>

        <!-- Amount -->
        <div class="ml-6 d-flex align-center justify-center">
          <span>
            {{ e.fromAmount }} {{ e.fromAsset.symbol }}
            <v-icon size="16" class="align-center">
              <ArrowRightIcon />
            </v-icon>
            {{ e.toAmount }} {{ e.toAsset.symbol }}
          </span>
        </div>

        <!-- State -->
        <div class="d-flex align-center flex-grow-1 justify-end mr-4">
          <v-progress-circular
            size="16"
            indeterminate
            color="primary"
            width="2"
            v-if="!e.completed"
          />
          <v-icon size="20" color="primary" v-else>
            <CheckIcon />
          </v-icon>
        </div>
      </div>
    </div>
  </v-sheet>
</template>

<script lang="ts" setup>
import {
  CheckIcon,
  XMarkIcon,
  ChevronUpIcon,
  ArrowRightIcon,
  ChevronDownIcon,
} from "@heroicons/vue/20/solid";
import { useBridgeStore } from "~/stores/bridge/bridge";
import type { Exchanges } from "~/types/history";
const bStore = useBridgeStore();

const example: Exchanges[] = [
  {
    fromAsset: bStore.fromAsset,
    toAsset: bStore.toAsset,
    fromAmount: 100,
    toAmount: 200,
    completed: false,
  },
  {
    fromAsset: bStore.toAsset,
    toAsset: bStore.fromAsset,
    fromAmount: 200,
    toAmount: 100,
    completed: false,
  },
  {
    fromAsset: bStore.toAsset,
    toAsset: bStore.fromAsset,
    fromAmount: 200,
    toAmount: 100,
    completed: true,
  },
  {
    fromAsset: bStore.toAsset,
    toAsset: bStore.fromAsset,
    fromAmount: 45,
    toAmount: 1,
    completed: true,
  },
  {
    fromAsset: bStore.toAsset,
    toAsset: bStore.fromAsset,
    fromAmount: 12.23123123,
    toAmount: 1523.523153212,
    completed: true,
  },
];
</script>

<style lang="scss" scoped>
.recent-sheet-shrink {
  position: absolute;
  right: 0px;
  bottom: 0px;
  height: 64px;
  width: 308px;
  border-top-left-radius: 32px;
  border-top-right-radius: 32px;
}
.recent-sheet-extended {
  position: absolute;
  right: 0px;
  bottom: 0px;
  height: 360px;
  width: 308px;
  border-top-left-radius: 32px;
  border-top-right-radius: 32px;
}

::-webkit-scrollbar {
  width: 8px;
  border-radius: 32px;
}
::-webkit-scrollbar-thumb {
  height: 16px;
  border-radius: 32px;
  background-color: var(--palette-black-10);
}
::-webkit-scrollbar-thumb:hover {
  background-color: var(--palette-black-25);
}
</style>