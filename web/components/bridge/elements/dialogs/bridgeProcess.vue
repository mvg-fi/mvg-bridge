<template>
  <v-dialog
    persistent
    no-click-animation
    v-model="bStore.bridgeProcessDialog"
    @keyup.esc="bStore.mutateBridgeProcess(false)"
    class="d-flex justify-center bridge-process-dialog"
  >
    <v-sheet
      class="rounded-xl align-self-start overflow-y-auto bridge-process-card"
      elevation="5"
    >
      <Title />
      <Route class="mt-2"/>
      <Details class="mt-8" style="margin-left: 120px; margin-right: 120px;"/>
      <Confirm class="mt-12"/>
    </v-sheet>
  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Title from "~/components/bridge/elements/dialogs/bridgeProcess/title.vue";
import Route from "~/components/bridge/elements/dialogs/bridgeProcess/route.vue";
import Details from "~/components/bridge/elements/dialogs/bridgeProcess/details.vue";
import Confirm from "~/components/bridge/elements/dialogs/bridgeProcess/confirm.vue";
const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.bridge-process-dialog {
  width: 700px;
}

.bridge-process-card {
  width: 680px;
  height: 370px;
  max-height: 600px;
  background-color: var(--palette-background-1);
}
</style>