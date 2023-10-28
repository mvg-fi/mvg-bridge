<template>
  <div class="d-flex mt-6 ml-8 justify-center align-center">
    <span class="title flex-grow-1"> {{ $t("connect_a_wallet") }} </span>
    <v-btn
      icon
      elevation="0"
      class="mr-8 d-flex align-center justify-center"
      style="width: 16px; height: 16px"
      @click="cStore.mutateDialog(false)"
    >
      <v-icon>
        <close-icon />
      </v-icon>
    </v-btn>
  </div>
</template>

<script setup>
import CloseIcon from "@heroicons/vue/24/outline/XCircleIcon";
import { useConnectStore } from "~/stores/connect/connect";
const cStore = useConnectStore();
</script>

<style lang="scss" scoped>
.title {
  font-size: 16px;
  line-height: 32px;
  font-family: Satoshi-SemiBold;
}
</style>