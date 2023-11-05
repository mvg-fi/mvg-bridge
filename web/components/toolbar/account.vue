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
        style="border: 0.5px solid var(--palette-main-primary)"
      >
        <span style="color: var(--palette-black-100)" class="h7m"
          >{{ cStore.connectedWallets.length }}
          {{
            $t("wallet")
          }}</span
        >
      </v-btn>
    </template>

    <!-- Connected account card -->

    <v-card class="account-card py-3">
      <v-list class="px-3">
        <div class="d-flex mb-2 mx-3">
          <span class="h6m d-flex align-center flex-grow-1"
            >{{ $t("wallet") }}
          </span>
          <v-btn
            icon
            style="
              width: 36px;
              height: 36px;
              color: var(--palette-black-75);
              border: 0.5px solid var(--palette-main-primary);
            "
            elevation="0"
            @click="cStore.connectMore()"
            class="d-flex justify-center align-center rounded-pill mr-4"
          >
            <v-tooltip
              activator="parent"
              location="bottom"
              transition="fade-transition"
            >
              {{ $t("connect_more_wallet") }}
            </v-tooltip>
            <v-icon size="16">
              <PlusIcon />
            </v-icon>
          </v-btn>

          <v-btn
            icon
            style="width: 36px; height: 36px; color: var(--palette-black-75); border: 0.5px solid var(--palette-main-primary);"
            elevation="0"
            @click="cStore.disconnectAll"
            class="d-flex justify-center align-center rounded-pill"
          >
            <v-tooltip
              activator="parent"
              location="bottom"
              transition="fade-transition"
            >
              {{ $t("disconnect_all") }}
            </v-tooltip>
            <v-icon size="16">
              <PowerIcon />
            </v-icon>
          </v-btn>
        </div>

        <v-list class="mt-2 py-0">
          <div
            v-for="w in cStore.connectedWallets"
            :key="w.name"
            style="border: 0.5px solid var(--palette-main-primary)"
            class="px-4 py-2 my-3 d-flex flex-row align-center rounded-pill"
          >
            <img
              style="width: 40px; height: 40px"
              :src="w.icon"
              class="rounded-pill mr-4"
            />
            <div
              class="flex-grow-1 d-flex flex-column mr-4"
              style="min-width: 140px"
            >
              <span class="h7" style="font-size: 14px">
                {{ w.name }}
              </span>
              <span class="h7m" style="font-size: 12px">
                {{ shortenAddress(w.address) || w.name }}
              </span>
            </div>

            <v-btn
              icon
              style="width: 16px; height: 16px; color: var(--palette-black-50)"
              elevation="0"
              @click="navigator.clipboard.writeText(w.address)"
              class="d-flex justify-center align-center rounded-pill mr-3"
            >
              <v-tooltip
                activator="parent"
                location="bottom"
                transition="fade-transition"
              >
                {{ $t("copy_address") }}
              </v-tooltip>
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
              <v-tooltip
                activator="parent"
                location="bottom"
                transition="fade-transition"
              >
                {{ $t("disconnect") }}
              </v-tooltip>
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
import { shortenAddress } from "~/helpers/utils";
import { useConnectStore } from "~/stores/connect/connect";
import { PlusIcon, PowerIcon, ClipboardIcon } from "@heroicons/vue/24/outline";

const cStore = useConnectStore();
const ConnectBtnClass = computed(() =>
  clsx("mx-4", "rounded-pill", "text-capitalize", "font-weight-bold")
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
  min-width: 344px;
  box-shadow: 0 3px 32px #0003 !important;
}
</style>