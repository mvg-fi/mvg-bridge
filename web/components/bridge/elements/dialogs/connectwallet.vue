<template>
  <v-dialog
    persistent
    no-click-animation
    v-model="cStore.connectDialog"
    class="d-flex justify-center connect-dialog"
  >
    <v-sheet
      class="rounded-xl align-self-start overflow-y-auto connect-card"
      elevation="3"
    >
      <Title />

      <div class="mt-3 mx-3">
        <Ethereum />
        <Bitcoin />
        <Mixin />
      </div>
    </v-sheet>
  </v-dialog>
</template>

<script setup>
import { useConnectStore } from "~/stores/connect/connect";
import Title from "~/components/bridge/elements/dialogs/connect/title.vue";
import Ethereum from "~/components/bridge/elements/dialogs/connect/ethereum.vue";
import Bitcoin from "~/components/bridge/elements/dialogs/connect/bitcoin.vue";
import Mixin from "~/components/bridge/elements/dialogs/connect/mixin.vue";
const cStore = useConnectStore();
</script>

<style lang="scss" scoped>
.connect-dialog {
  width: 350px;
}
.connect-card {
  width: 305px;
  height: 600px;
  background-color: var(--palette-background-1);
}

::-webkit-scrollbar {
  width: 8px;
  border-radius: 32px;
}
::-webkit-scrollbar-thumb {
  height: 16px;
  border-radius: 32px;
  background-color: var(--palette-black-10);
}
::-webkit-scrollbar-thumb:hover {
  background-color: var(--palette-black-25);
}
</style>