<template>
  <v-overlay
    location="bottom end"
    location-strategy="connected"
    scrim="false"
    transition="scale-transition"
    opacity="0.1"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        size="large"
        v-bind="props"
        :class="ConnectBtnClass"
        style="background-color: var(--palette-black-5)"
      >
        <span style="color: var(--palette-black-100)" class="h7m"
          >{{ cStore.connectedWallets.length }}
          {{
            $t("wallet", {
              multi: cStore.connectedWallets.length == 1 ? "" : "s",
            })
          }}</span
        >
      </v-btn>
    </template>

    <!-- Connected account card -->

    <v-card class="account-card py-4">
      <v-list class="px-6">
        <div class="d-flex mb-2">
          <span class="h6m d-flex align-center flex-grow-1">{{ $t("wallet", {multi: 's'}) }}</span>
          <v-btn
            icon
            style="
              width: 36px;
              height: 36px;
              color: var(--palette-black-75);
              background-color: var(--palette-black-10);
            "
            elevation="0"
            @click="cStore.connectMore()"
            class="d-flex justify-center align-center rounded-pill mr-4"
          >
            <v-icon size="16">
              <PlusIcon />
            </v-icon>
          </v-btn>

          <v-btn
            icon
            style="
              width: 36px;
              height: 36px;
              color: var(--palette-black-75);
              background-color: var(--palette-black-10);
            "
            elevation="0"
            @click="cStore.disconnectAll"
            class="d-flex justify-center align-center rounded-pill"
          >
            <v-icon size="16">
              <PowerIcon />
            </v-icon>
          </v-btn>
        </div>

        <v-list class="mt-2 py-0">
          <div
            v-for="w in cStore.connectedWallets"
            :key="w.name"
            style="border: 1px solid var(--palette-black-10); background-color: var(--palette-black-5);"
            class="px-0 py-0 my-4 d-flex flex-row align-center rounded-pill"
          >
            <img
              style="width: 36px; height: 36px"
              :src="w.icon"
              class="rounded-pill mr-4"
            />
            <span class="h7m mr-4" style="font-size: 16px; width: 184px">
              {{ shortenAddress(w.address) || w.name }}
            </span>

            <v-btn
              icon
              style="width: 16px; height: 16px; color: var(--palette-black-50)"
              elevation="0"
              @click="navigator.clipboard.writeText(w.address)"
              class="d-flex justify-center align-center rounded-pill mr-3"
            >
              <v-icon size="16">
                <ClipboardIcon />
              </v-icon>
            </v-btn>

            <v-btn
              icon
              style="width: 16px; height: 16px; color: var(--palette-black-50)"
              elevation="0"
              @click="cStore.disconnectSpecific(w)"
              class="d-flex justify-center align-center rounded-pill mr-3"
            >
              <v-icon size="16">
                <PowerIcon />
              </v-icon>
            </v-btn>
          </div>
        </v-list>
      </v-list>
    </v-card>
  </v-overlay>
</template>

<script setup>
import clsx from "clsx";
import { PlusIcon, PowerIcon, ClipboardIcon } from "@heroicons/vue/24/outline";
import { useConnectStore } from "~/stores/connect/connect";
import { shortenAddress } from "~/helpers/utils";

const cStore = useConnectStore();
const ConnectBtnClass = computed(() =>
  clsx(
    "rounded-pill",
    "text-capitalize",
    "font-weight-bold",
    "elevation-1",
    "mx-4"
  )
);

const connectMoreWallet = () => {
  console.log(cStore.connectedWallets);
};
</script>

<style lang="scss" scoped>
.btn-hover {
  opacity: 0.85;
}
.account-card {
  margin-top: 8px;
  border-radius: 32px;
  width: 344px;
  box-shadow: 0 3px 32px #0003 !important;
}
</style>