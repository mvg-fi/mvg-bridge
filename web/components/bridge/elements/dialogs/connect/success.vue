<template>
  <v-sheet
    class="rounded-xl align-self-start overflow-y-auto success-card"
    elevation="3"
  >
    <!-- Title -->
    <div class="d-flex mt-6 ml-6 justify-center align-center">
      <span class="success-title flex-grow-1"> {{ $t("connected") }} </span>
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

    <!-- Check Icon -->
    <div class="d-flex justify-center align-center my-16">
      <div
        class="rounded-circle d-flex justify-center align-center"
        style="width: 140px; height: 140px; background-color: var(--palette-main-correct)"
      >
        <v-icon size="100" style="color: var(--palette-background-1)">
          <CheckIcon />
        </v-icon>
      </div>
    </div>
  </v-sheet>
</template>

<script lang="ts" setup>
import { useConnectStore } from "~/stores/connect/connect";
import { XCircleIcon } from "@heroicons/vue/24/outline";
import { CheckIcon } from "@heroicons/vue/20/solid";

const cStore = useConnectStore();
</script>

<style lang="scss" scoped>
.success-title {
  font-size: 18px;
  line-height: 32px;
  font-family: Satoshi-Bold;
}
.success-card {
  width: 328px;
  height: 345px;
  background-color: var(--palette-background-1);
}
</style>