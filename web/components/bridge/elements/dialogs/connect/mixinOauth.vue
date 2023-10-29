<template>
  <v-sheet
    class="rounded-xl align-self-start overflow-y-auto mixin-oauth-card"
    elevation="3"
  >
    <!-- Title -->
    <div class="d-flex mt-6 ml-6 justify-center align-center">
      <span class="oauth-title flex-grow-1"> {{ $t("scan_to_connect") }} </span>
      <v-btn
        icon
        elevation="0"
        class="mr-6 d-flex align-center justify-center"
        style="width: 16px; height: 16px"
        @click="cStore.mutateDialog(false)"
      >
        <v-icon>
          <close-icon />
        </v-icon>
      </v-btn>
    </div>

    <!-- QRcode -->
    <div class="d-flex justify-center">
      <div class="ma-8" v-if="qrLoaded">
        <div>
          <v-img
            eager
            :src="mx"
            style="
              width: 48px;
              height: 48px;
              top: 180px;
              left: 128px;
              border-radius: 32px;
              position: absolute;
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
          <v-progress-circular indeterminate color="primary" />
        </div>
      </div>
    </div>

    <!-- Helper text -->
    <div class="d-flex flex-column text-center justify-center align-center">
      <!-- Open in Mixin -->
      <!-- <v-btn
        elevation="0"
        class="rounded-pill mt-3 pa-0"
        style="background-color: var(--palette-main-secondary); width: 128px;"
      >
        <span class="h7" style="color: var(--palette-background-1);">{{ $t("open_in_mixin") }}</span>
      </v-btn> -->

      <!-- Install -->
      <span class="helper-text">
        {{ $t("don't_have_mixin_messenger") }}
      </span>
      <a :href="MixinMessengerLink" target="_blank">
        <v-btn
          elevation="0"
          class="rounded-pill mt-3"
          style="background-color: var(--palette-black-10); width: 170px"
        >
          <span class="helper-text">{{ $t("install") }}</span>
        </v-btn>
      </a>
    </div>
  </v-sheet>
</template>

<script setup>
import authorize from "~/helpers/mixin/oauth";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import mx from "~/assets/images/wallets/mixin.png";
import CloseIcon from "@heroicons/vue/24/outline/XCircleIcon";
import { appendConnected, mixinOauthSuccess, useConnectStore } from "~/stores/connect/connect";
import {
  BridgeBotID,
  MixinChainName,
  MixinMessengerLink,
  MixinMessengerName,
  OAuthScope,
} from "~/helpers/constants";
import { userMe } from "~/helpers/mixin/user";
const cStore = useConnectStore();

let qrLoaded = ref(false);
let qrURL = ref("");

const auth = () => {
  authorize(
    { clientId: BridgeBotID, scope: OAuthScope, pkce: true },
    {
      onShowUrl: (url) => {
        qrLoaded.value = true;
        qrURL.value = url;
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
  background-color: var(--palette-background-1);
}
.helper-text {
  font-size: 12px;
  font-family: Satoshi-Bold;
}
</style>