<template>
  <v-dialog
    persistent
    no-click-animation
    v-model="bStore.bridgeProcessDialog"
    @keyup.esc="1"
    class="d-flex justify-center bridge-process-dialog"
  >
    <v-sheet
      class="rounded-xl align-self-start overflow-y-auto bridge-process-card"
      elevation="3"
    >
      <div class="mt-3 mx-5 mb-6">
      </div>
    </v-sheet>

  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import { useConnectStore } from "~/stores/connect/connect";
const bStore = useBridgeStore();
const cStore = useConnectStore();
</script>

<style lang="scss" scoped>
.bridge-process-dialog {
  width: 350px;
}

.bridge-process-card {
  width: 305px;
  max-height: 600px;
  background-color: var(--palette-background-1);
}
</style>