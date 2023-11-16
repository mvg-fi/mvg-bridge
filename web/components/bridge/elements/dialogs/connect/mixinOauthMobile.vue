<template>
  <v-sheet
    :class="
      clsx(
        'align-self-center overflow-y-auto bg-background',
        'mixin-oauth-card-mobile rounded-t-xl',
        expanded ? 'pb-12' : 'pb-2'
      )
    "
    elevation="3"
  >
    <!-- Title -->
    <v-row no-gutters class="d-flex mt-6 justify-center align-center">
      <v-col cols="2" />
      <v-col cols="8" class="text-center justify-center"
        ><span class="oauth-title">
          {{ $t("connect_to_mixin") }}
        </span>
      </v-col>
      <v-col cols="2" class="d-flex align-center justify-center">
        <v-btn
          icon
          elevation="0"
          class="mr-4"
          style="width: 16px; height: 16px"
          @click="cStore.mutateDialog(false)"
        >
          <v-icon>
            <XMarkIcon />
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>

    <!-- Icon -->
    <div class="d-flex justify-center" v-if="!showQR">
      <div class="ma-8">
        <v-icon size="72" class="rounded-pill">
          <v-img :src="mx" />
        </v-icon>
      </div>
    </div>
    <!-- QRCode -->
    <div class="d-flex justify-center" v-else>
      <div class="ma-8" v-if="qrLoaded">
        <div style="position: relative">
          <v-img
            eager
            :src="mx"
            style="
              width: 48px;
              height: 48px;
              position: absolute;
              top: 50%;
              left: 50%;
              transform: translate(-50%, -50%);
              border-radius: 32px;
            "
          />
          <VueQrcode
            :key="qrURL"
            :value="qrURL"
            :options="{ width: 232, padding: 1, margin: 1 }"
          />
        </div>
      </div>
      <div v-else class="ma-8">
        <div
          style="height: 232px; width: 232px"
          class="d-flex justify-center align-center"
        >
          <v-progress-circular indeterminate color="primary" width="2" />
        </div>
      </div>
    </div>

    <!-- Helper text -->
    <div class="d-flex flex-column text-center justify-center align-center">
      <div :class="clsx('d-flex flex-row', expanded ? 'mb-10' : 'mb-6')">
        <!-- Open in Mixin -->
        <v-btn
          elevation="0"
          @click="showQR = !showQR"
          style="width: 128px; background-color: var(--palette-black-10)"
          class="rounded-pill ma-3 pa-0"
        >
          <span class="h7 text-black">{{ $t("show_qr_code") }}</span>
        </v-btn>
        <v-btn
          elevation="0"
          @click="openInMixin"
          style="width: 128px"
          :disabled="!qrLoaded"
          :loading="!qrLoaded || clicked"
          class="rounded-pill ma-3 pa-0 bg-primary"
        >
          <span class="h7 text-background">{{ $t("open_in_mixin") }}</span>
          <!-- <v-progress-circular
            color="background"
            size="16"
            indeterminate
            width="2"
            v-else
          /> -->
        </v-btn>
      </div>

      <div v-if="!expanded">
        <button @click="expanded = true">
          <v-icon size="14" class="stroke-black-darken-3">
            <ChevronDownIcon />
          </v-icon>
        </button>
      </div>

      <div class="d-flex flex-column" v-else>
        <!-- Install -->
        <span class="h7">
          {{ $t("don't_have_mixin_messenger") }}
        </span>
        <!-- <a :href="MixinMessengerLink"> -->
          <v-btn
            elevation="0"
            @click="openInMixin()"
            class="rounded-pill mt-3"
            style="background-color: var(--palette-black-10); width: 120px"
          >
            <span class="h7">{{ $t("install") }}</span>
          </v-btn>
        <!-- </a> -->
      </div>
    </div>
  </v-sheet>
</template>

<script setup>
import clsx from "clsx";
import authorize from "~/helpers/mixin/oauth";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import mx from "~/assets/images/wallets/mixin.png";
import { ChevronDownIcon, XMarkIcon } from "@heroicons/vue/24/outline";
import { mixinOauthSuccess, useConnectStore } from "~/stores/connect/connect";
import {
  BridgeBotID,
  MixinMessengerLink,
  OAuthScope,
} from "~/helpers/constants";

const cStore = useConnectStore();

let qrLoaded = ref(false);
let qrURL = ref("");
let clicked = ref(false);
let expanded = ref(false);
let showQR = ref(false);

const auth = () => {
  authorize(
    { clientId: BridgeBotID, scope: OAuthScope, pkce: true },
    {
      onShowUrl: (url) => {
        qrURL.value = url;
        qrLoaded.value = true;
        console.log("onShowURL:", url);
      },
      onError: (error) => {
        console.error(error);
        return;
      },
      onSuccess: async (data) => {
        await mixinOauthSuccess(data);
        return;
      },
    }
  );
};
onMounted(() => {
  auth();
});
onUnmounted(() => {
  qrLoaded = false;
  qrURL = "";
});

const openInMixin = () => {
  window.open(qrURL.value);
  clicked.value = true;
};
</script>

<style lang="scss" scoped>
.oauth-title {
  font-size: 18px;
  line-height: 32px;
  font-family: Satoshi-Bold;
}
.mixin-oauth-card {
  width: 305px;
  height: 464px;
}
.mixin-oauth-card-mobile {
  width: 100vw;
}
</style>