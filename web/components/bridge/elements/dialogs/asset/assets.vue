<template>
  <div class="w-100 overflow-y-auto" v-if="store.filteredItems.length != 0">
    <v-list class="pt-0">
      <v-list-item
        :disabled="
          item.asset_id == store.fromAsset?.asset_id || 
          item.asset_id == store.toAsset?.asset_id
        "
        class="w-100 px-0"
        v-for="item in store.filteredItems"
        :key="item.asset_id"
      >
        <v-btn
          block
          elevation="0"
          height="54px"
          :class="
            clsx(
              'd-flex justify-start',
              store.selectNetworkBar && 'px-5',
              !store.selectNetworkBar && 'px-7'
            )
          "
          @click="
            store.setAsset(props.from, item);
          "
        >
          <v-icon style="width: 36px; height: 36px; position: reactive">
            <v-img :src="item.icon" alt="item.symbol" style="transition:opacity 0.15s linear 0s; opacity: 1;"/>
            <v-img
              :src="item.chain_icon"
              v-if="item.asset_id != item.chain_id"
              style="
                height: 14px;
                width: 14px;
                position: absolute;
                bottom: 0px;
                left: 0px;
              "
            ></v-img>
          </v-icon>

          <div class="d-flex flex-column justify-start text-start ml-4">
            <span
              class="font-weight-semibold"
              style="letter-spacing: 0px; font-size: 14px"
            >
              {{ item.symbol }}
            </span>

            <span
              class="font-weight-light"
              style="font-size: 12px; color: grey; opacity: 60%"
            >
              {{ item.chain_name }}
            </span>
          </div>

          <v-icon
            class="mr-5"
            style="position:absolute;right:0px;"
            v-if="
              props.from
                ? store.fromAsset?.asset_id === item.asset_id
                : store.toAsset?.asset_id === item.asset_id
            "
          >
            <CheckIcon />
          </v-icon>
        </v-btn>
      </v-list-item>
    </v-list>
  </div>
  <div
    class="w-100 h-100 d-flex flex-column align-center justify-center text-center"
    v-else
  >
    <div>
      <svg
        width="80"
        height="80"
        viewBox="0 0 80 80"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          opacity=".1"
          d="M48 36H24m15 14H24m31-28H24m44 16.75V21.8c0-5.88 0-8.82-1.144-11.067a10.5 10.5 0 0 0-4.59-4.589C60.022 5 57.082 5 51.2 5H28.8c-5.88 0-8.82 0-11.067 1.144a10.5 10.5 0 0 0-4.589 4.59C12 12.978 12 15.918 12 21.8v36.4c0 5.88 0 8.82 1.144 11.067a10.5 10.5 0 0 0 4.59 4.589C19.978 75 22.918 75 28.8 75h13.45M75 75l-5.25-5.25m3.5-8.75c0 6.766-5.484 12.25-12.25 12.25S48.75 67.766 48.75 61 54.234 48.75 61 48.75 73.25 54.234 73.25 61Z"
          stroke="#000"
          stroke-width="5"
          stroke-linecap="round"
          stroke-linejoin="round"
        ></path>
      </svg>
    </div>
    <span
      class="mt-2 font-weight-bold unselectable"
      style="text-transform: uppercase; color: #d1d3d6"
      >{{ $t("no_result") }}</span
    >
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { CheckIcon } from "@heroicons/vue/24/outline";
const store = useBridgeStore();
const props = defineProps(["from"]);
</script>

<style lang="scss" scoped>
* {
  scrollbar-width: none;
}
::-webkit-scrollbar {
  display: none;
}
</style>