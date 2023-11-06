<template>
  <!-- Title -->
  <div class="text-center pt-0 pb-4">
    <span class="h7m" style="font-size: 16px">
      {{ $t("transfer") }}
      {{ bStore.bridgeAmount }}
      {{ bStore.fromAsset.symbol }}
    </span>
  </div>

  <div class="flex flex-col">
    <!-- Network -->
    <div class="deposit-title flex flex-row bg-background rounded-lg">
      <div class="d-flex flex-grow-1 flex-column pa-2 px-3">
        <div class="text-black-darken-5">
          <span> {{ $t("network") }} </span>
        </div>

        <div class="flex flex-row items-center select-text mt-1">
          <span class="h7" style="overflow-wrap: break-word">
            {{ bStore.fromAsset.chain_name }}
          </span>
        </div>
      </div>
    </div>

    <!-- QrCode -->
    <div class="d-flex justify-center my-4">
      <div
        style="
          border: 1px solid var(--palette-black-10);
          border-radius: 16px;
          position: relative;
        "
        class="pa-3"
      >
        <VueQrcode
          :key="address"
          :value="address"
          :options="{ width: 130, padding: 0, margin: 0 }"
        />
        <img
          alt=""
          :src="bStore.fromAsset.chain_icon"
          style="
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            height: 36px;
            width: 36px;
          "
        />
      </div>
    </div>

    <!-- Address -->
    <div class="d-flex flex-row bg-background rounded-lg">
      <div class="d-flex flex-column pa-2 px-3">
        <div class="text-black-darken-5">
          <span> {{ $t("address") }} </span>
        </div>

        <div class="d-flex flex-row align-center mt-1">
          <span class="h7" style="overflow-wrap: break-word">
            {{ address }}
          </span>
        </div>
      </div>
      <Copy
        :copyText="address"
        class="d-flex px-4 justify-center align-center"
      />
    </div>

    <!-- Memo -->
    <div
      style="font-size: 12px; width: 340px;"
      class="deposit-info d-flex flex-column my-2 mx-2 h7m text-black-darken-4"
    >
      <!-- <span v-if="bStore.fromAsset.deposit_entries[0].tag">
        - {{ $t("memo_required") }}
      </span> -->
      <span> - {{ $t("view_address_intro1", { number: 0 }) }} </span>
      <span> - {{ $t("view_address_intro2") }} </span>
    </div>
  </div>
  <div
    class="d-flex align-center justify-center flex-row my-4 mt-8"
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

    <v-btn class="confirm-btn rounded-pill" size="large" @click="confirmPaid">
      <span class="h7 text-background">
        {{ $t("ive_transferred") }}
      </span>
    </v-btn>
  </div>
</template>

<script lang="ts" setup>
import Copy from "./copy.vue";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import { useBridgeStore } from "~/stores/bridge/bridge";
const bStore = useBridgeStore();
const address = "asjdfiasjdfkasjdfkljaslkdfjqlwkjerkqwjeoijoi";

const confirmPaid = () => {};
</script>