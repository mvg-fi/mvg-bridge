<template>
  <div class="d-flex mt-6 ml-8 align-center">
    <span class="h6 flex-grow-1"> {{ $t("connect_a_wallet") }} </span>
    <v-btn
      icon
      elevation="0"
      class="mr-6"
      style="width: 24px; height: 24px"
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