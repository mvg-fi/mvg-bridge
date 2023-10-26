<template>
  <v-dialog
    v-model="store.connectDialog"
    class="d-flex justify-center connect-dialog"
  >
    <v-card
      class="rounded-xl overflow-y-hidden connect-card"
      elevation="3"
    >
      <Title />

      <div class="mt-5 mx-8">
        <Ethereum />
        <Bitcoin />
        <Mixin />
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { useConnectStore } from "~/stores/connect/connect";
import Title from "~/components/bridge/elements/dialogs/connect/title.vue";
import Ethereum from "~/components/bridge/elements/dialogs/connect/ethereum.vue";
import Bitcoin from "~/components/bridge/elements/dialogs/connect/bitcoin.vue";
import Mixin from "~/components/bridge/elements/dialogs/connect/mixin.vue";
const store = useConnectStore();
</script>

<style lang="scss" scoped>
.connect-dialog {
  width: 305px;
}
.connect-card {
  width: 305px;
  background-color: var(--palette-background-1);
}
</style>