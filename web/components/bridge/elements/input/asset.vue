<template>
  <div class="d-flex align-center">
    <v-btn
      class="rounded-xl d-flex flex-row pa-2 mr-2"
      elevation="0"
      size="x-large"
      @click="store.mutateDialog('from', true)"
    >
      <v-icon style="height: 32px; width: 32px" class="mx-2">
        <v-img
          :src="
            from
              ? store.fromAsset.icon
              : store.toAsset.icon
          "
        />
      </v-icon>

      <div class="d-flex flex-column align-center justify-center mx-1">
        <span style="font-size: 16px; font-weight: 600">
          {{
            from
              ? store.fromAsset.symbol
              : store.toAsset.symbol
          }}
        </span>
        <span
          style="
            font-size: 10px;
            opacity: 40;
            color: grey;
            font-weight: 100;
            line-height: 1rem;
            letter-spacing: 0.1px !important;
          "
        >
          {{
            from
              ? store.fromAsset.name
              : store.toAsset.name
          }}
        </span>
      </div>

      <v-icon style="width: 20px; height: 20px" class="mx-1">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          height="24"
          width="24"
        >
          <path
            xmlns="http://www.w3.org/2000/svg"
            d="M17 10L12 16L7 10H17Z"
            fill="#0D0D0D"
          />
        </svg>
      </v-icon>
    </v-btn>
    <FromAssetDialog/>
  </div>
</template>

<script setup>
import FromAssetDialog from "~/components/bridge/elements/dialogs/fromAsset.vue"
import EthereumIcon from "~/assets/images/networks/ethereum.png";
import { ChevronDownIcon } from "@heroicons/vue/20/solid";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();
const props = defineProps(["from"]);
</script>

<style lang="scss" scoped>
</style>