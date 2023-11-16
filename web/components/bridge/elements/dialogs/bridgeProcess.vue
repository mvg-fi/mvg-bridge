<template>
  <v-dialog
    persistent
    v-if="!mobile"
    no-click-animation
    v-model="bStore.bridgeProcessDialog"
    @keyup.esc="bStore.mutateBridgeProcess(false)"
    :transition="'fade-transition'"
    :class="clsx('d-flex justify-center dialog-blur')"
  >
    <v-sheet
      elevation="5"
      v-if="bStore.bridgeProcessState == 0"
      class="rounded-xl align-self-center overflow-y-auto bg-background bridge-process-card py-4"
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

    <v-sheet
      elevation="5"
      v-else-if="bStore.bridgeProcessState == 2"
      class="rounded-xl align-self-center overflow-y-auto bg-background bridge-process-card pa-4"
    >
      <PayWithWallet />
    </v-sheet>
  </v-dialog>

  <v-bottom-sheet
    v-model="bStore.bridgeProcessDialog"
    class="dialog-blur"
    content-class="elevation-0"
    v-else
  >
    <v-sheet
      elevation="5"
      v-if="bStore.bridgeProcessState == 0"
      :class="
        clsx(
          'rounded-t-xl align-self-center overflow-y-auto bg-background bridge-process-card-mobile py-4'
        )
      "
    >
      <Title />
      <Route class="mt-2" />
      <Details class="my-8" />
      <Confirm class="py-4 mb-2" />
    </v-sheet>

    <v-sheet
      elevation="5"
      v-else-if="bStore.bridgeProcessState == 1"
      :class="
        clsx(
          'rounded-t-xl align-self-center overflow-y-auto bg-background bridge-process-card-mobile',
          'py-4'
        )
      "
    >
      <ViewAddress />
    </v-sheet>

    <v-sheet
      elevation="5"
      v-else-if="bStore.bridgeProcessState == 2"
      :class="
        clsx(
          'rounded-t-xl align-self-center overflow-y-auto bg-background bridge-process-card-mobile',
          'py-4'
        )
      "
    >
      <PayWithWallet />
    </v-sheet>
  </v-bottom-sheet>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { VBottomSheet } from "vuetify/labs/VBottomSheet";
import Title from "~/components/bridge/elements/dialogs/bridgeProcess/title.vue";
import Route from "~/components/bridge/elements/dialogs/bridgeProcess/route.vue";
import Details from "~/components/bridge/elements/dialogs/bridgeProcess/details.vue";
import Confirm from "~/components/bridge/elements/dialogs/bridgeProcess/confirm.vue";
import ViewAddress from "~/components/bridge/elements/dialogs/bridgeProcess/viewAddress.vue";
import PayWithWallet from "~/components/bridge/elements/dialogs/bridgeProcess/payWithWallet.vue";
const bStore = useBridgeStore();
const { mobile } = useDisplay();
</script>

<style lang="scss" scoped>
.bridge-process-card {
  min-width: 340px;
}
.bridge-process-card-mobile {
  width: 100vw;
}
</style>