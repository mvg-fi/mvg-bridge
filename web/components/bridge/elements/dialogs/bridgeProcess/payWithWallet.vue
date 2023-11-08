<template>
  <!-- Title -->
  <div class="text-center pt-0 pb-4" v-if="payableWallets.length > 0">
    <span class="h7m" style="font-size: 16px">
      {{ $t("choose_a_wallet_to_pay") }}
    </span>
  </div>

  <div class="text-center pt-0 pb-4" v-else>
    <span class="h7m" style="font-size: 16px">
      {{ $t("no_wallet_found", {"wallet_name": bStore.fromAsset.symbol}) }}
    </span>
  </div>

  <div class="d-flex flex-column" v-if="payableWallets.length > 0">
    <div
      :key="w"
      v-for="w in payableWallets"
      @click="console.log(w)"
      style="border: 0.5px solid var(--palette-black-10)"
      class="px-4 py-2 my-3 d-flex flex-row align-center rounded-pill"
    >
      <img
        style="width: 40px; height: 40px"
        :src="w.icon"
        class="rounded-pill mr-4"
      />
      <div class="d-flex flex-column mr-4" style="min-width: 140px">
        <span class="h7" style="font-size: 14px">
          {{ w.name }}
        </span>
        <span class="h7m" style="font-size: 12px">
          {{ shortenAddress(w.address) || w.name }}
        </span>
      </div>
    </div>
  </div>
  <div class="py-8 d-flex flex-column align-center justify-center text-center" v-else>
    <!-- No wallet found -->
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
    <span
      class="mt-2 font-weight-bold unselectable"
      style="text-transform: uppercase; color: #d1d3d6"
      >{{ $t("no_wallet_found") }}</span
    >
  </div>

  <div
    class="d-flex align-center justify-center flex-row my-4 mt-4"
    style="gap: 24px"
  >
    <v-btn
      size="large"
      class="bg-background rounded-pill"
      elevation="0"
      style="border: 1px solid var(--palette-black-10)"
      @click="
        () => {
          bStore.setBridgeProcessState(0);
        }
      "
    >
      <span class="h7 text-black-darken-5">
        {{ $t("back") }}
      </span>
    </v-btn>

    <v-btn
      size="large"
      elevation="0"
      v-if="payableWallets.length === 0"
      class="bg-background rounded-pill confirm-btn"
      style="border: 1px solid var(--palette-black-10)"
      @click="
        () => {
          cStore.mutateDialog(true);
        }
      "
    >
      <span class="h7 text-background">
        {{ $t("connect_wallet") }}
      </span>
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import { shortenAddress } from "~/helpers/utils";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { useConnectStore } from "~/stores/connect/connect";
import type { ConnectedWallet } from "~/types/wallet";

const bStore = useBridgeStore();
const cStore = useConnectStore();
const address = "asjdfiasjdfkasjdfkljaslkdfjqlwkjerkqwjeoijoi";

const payableWallets = computed(() =>
  cStore.connectedWallets.filter(
    (e: ConnectedWallet) => e.chain_id === bStore.fromAsset.chain_id
  )
);
</script>