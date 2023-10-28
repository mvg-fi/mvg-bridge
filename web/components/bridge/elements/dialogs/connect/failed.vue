<template>
  <v-sheet
    class="rounded-xl align-self-start overflow-y-auto failed-card"
    elevation="3"
  >
    <!-- Title -->
    <div class="d-flex mt-6 ml-6 justify-center align-center">
      <span class="failed-title flex-grow-1"> {{ $t("canceled") }} </span>
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
        @click="cStore.setConnectState(0)"
        class="rounded-pill mr-4"
      >
        <span class="h7"> {{ $t("back") }} </span>
      </v-btn>

      <v-btn
        elevation="0"
        style="
          background-color: var(--palette-black-10);
          width: 120px;
          height: 38px;
        "
        class="rounded-pill"
      >
        <span class="h7"> {{ $t("retry") }} </span>
      </v-btn>
    </div>
  </v-sheet>
</template>

<script lang="ts" setup>
import { useConnectStore } from "~/stores/connect/connect";
import { XCircleIcon } from "@heroicons/vue/24/outline";
import { XMarkIcon } from "@heroicons/vue/20/solid";

const cStore = useConnectStore();
</script>

<style lang="scss" scoped>
.failed-title {
  font-size: 18px;
  line-height: 32px;
  font-family: Satoshi-Bold;
}
.failed-card {
  width: 328px;
  height: 360px;
  background-color: var(--palette-background-1);
}
</style>