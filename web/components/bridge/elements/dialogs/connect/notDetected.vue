<template>
  <v-sheet
    :class="
      clsx(
        'align-self-center overflow-y-auto',
        mobile ? 'detect-card-mobile rounded-t-xl' : 'detect-card rounded-xl'
      )
    "
    elevation="3"
  >
    <!-- Title -->
    <div class="d-flex mt-6 ml-6 justify-center align-center">
      <span class="detect-title flex-grow-1">
        {{ $t("haven't_detected") }}
      </span>
      <v-btn
        icon
        elevation="0"
        class="mr-6 d-flex align-center justify-center"
        style="width: 16px; height: 16px"
        @click="cStore.mutateDialog(false)"
      >
        <v-icon>
          <XCircleIcon />
        </v-icon>
      </v-btn>
    </div>

    <!-- X Icon -->
    <div class="d-flex justify-center align-center my-12">
      <div
        class="rounded-circle d-flex justify-center align-center"
        style="
          width: 140px;
          height: 140px;
          background-color: var(--palette-main-error);
        "
      >
        <v-icon size="100" style="color: var(--palette-background-1)">
          <XMarkIcon />
        </v-icon>
      </div>
    </div>

    <!-- Actions -->
    <div class="d-flex flex-row justify-center align-center">
      <v-btn
        elevation="0"
        style="
          background-color: var(--palette-black-5);
          width: 120px;
          height: 38px;
        "
        @click="cStore.setConnectState(0); cStore.setConnecting({})"
        class="rounded-pill mr-4"
      >
        <span class="h7"> {{ $t("back") }} </span>
      </v-btn>

      <v-btn
        elevation="0"
        style="
          background-color: var(--palette-black-5);
          width: 120px;
          height: 38px;
        "
        class="rounded-pill"
        @click="goInstall"
      >
        <span class="h7"> {{ $t("install") }} </span>
      </v-btn>
    </div>
  </v-sheet>
</template>

<script lang="ts" setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useConnectStore } from "~/stores/connect/connect";
import { XCircleIcon } from "@heroicons/vue/24/outline";
import { XMarkIcon } from "@heroicons/vue/20/solid";
import { GetInstallLink } from "~/stores/connect/install";

const { mobile } = useDisplay();
const cStore = useConnectStore();
const goInstall = () => {
  window.open(GetInstallLink(cStore.connecting));
};
</script>

<style lang="scss" scoped>
.detect-title {
  font-size: 18px;
  line-height: 32px;
  font-family: Satoshi-Bold;
}
.detect-card {
  width: 328px;
  height: 360px;
  background-color: var(--palette-background-1);
}
.detect-card-mobile {
  width: 100vw;
  height: 360px;
  background-color: var(--palette-background-1);
}
</style>