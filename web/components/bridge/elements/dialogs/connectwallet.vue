<template>
  <v-dialog
    persistent
    no-click-animation
    :fullscreen="mobile"
    v-model="cStore.connectDialog"
    @keyup.esc="cStore.mutateDialog(false)"
    :class="clsx('d-flex justify-center dialog-blur', mobile && 'mobile-dialog')"
    :transition="mobile ? 'slide-y-reverse-transition' : 'fade-transition'"
  >
    <v-sheet
      v-if="cStore.connectState === 0"
      class="rounded-xl align-self-center overflow-y-auto connect-card bg-background"
      elevation="3"
    >
      <Title />

      <div class="mt-3 mx-5 mb-6">
        <Ethereum />
        <Bitcoin />
        <Mixin />
      </div>
    </v-sheet>

    <MixinOauth v-else-if="cStore.connectState === 1" />
    <Success v-else-if="cStore.connectState === 2" />
    <Failed v-else-if="cStore.connectState === 3" />
  </v-dialog>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useConnectStore } from "~/stores/connect/connect";
import { useWeb3ModalAccount } from "@web3modal/ethers5/vue";
import Title from "~/components/bridge/elements/dialogs/connect/title.vue";
import Ethereum from "~/components/bridge/elements/dialogs/connect/ethereum.vue";
import Bitcoin from "~/components/bridge/elements/dialogs/connect/bitcoin.vue";
import Mixin from "~/components/bridge/elements/dialogs/connect/mixin.vue";
import MixinOauth from "~/components/bridge/elements/dialogs/connect/mixinOauth.vue";
import Success from "~/components/bridge/elements/dialogs/connect/success.vue";
import Failed from "~/components/bridge/elements/dialogs/connect/failed.vue";

const { mobile } = useDisplay();

const cStore = useConnectStore();
watch(cStore.connected, () => {
  try {
    const { address, chainId, isConnected } = useWeb3ModalAccount();
    console.log(address, chainId, isConnected);
  } catch (e) {
    console.log(e);
  }
});
</script>

<style lang="scss" scoped>
.connect-card {
  width: 305px;
  max-height: 600px;
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