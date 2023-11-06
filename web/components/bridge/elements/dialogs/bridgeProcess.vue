<template>
  <v-dialog
    persistent
    no-click-animation
    v-model="bStore.bridgeProcessDialog"
    @keyup.esc="bStore.mutateBridgeProcess(false)"
    class="d-flex justify-center align-center dialog-blur"
  >
    <v-sheet
      class="rounded-xl align-self-center overflow-y-auto bg-background bridge-process-card py-4"
      elevation="5"
      v-if="bStore.bridgeProcessState == 0"
    >
      <Title />
      <Route class="mt-2" />
      <Details class="my-8" />
      <Confirm class="py-4 mb-2" />
    </v-sheet>

    <v-sheet
      elevation="5"
      v-else-if="bStore.bridgeProcessState == 1"
      class="rounded-xl align-self-center overflow-y-auto bg-background bridge-process-card pa-4"
    >
      <ViewAddress />
    </v-sheet>
  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Title from "~/components/bridge/elements/dialogs/bridgeProcess/title.vue";
import Route from "~/components/bridge/elements/dialogs/bridgeProcess/route.vue";
import Details from "~/components/bridge/elements/dialogs/bridgeProcess/details.vue";
import Confirm from "~/components/bridge/elements/dialogs/bridgeProcess/confirm.vue";
import ViewAddress from "~/components/bridge/elements/dialogs/bridgeProcess/viewAddress.vue";
const bStore = useBridgeStore();
</script>

<style lang="scss" scoped>
.bridge-process-card {
  min-width: 360px;
}
</style>