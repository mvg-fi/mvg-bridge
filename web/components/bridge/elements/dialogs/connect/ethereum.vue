<template>
  <div class="d-flex flex-column">
    <div class="h7m mb-2 px-3">
      <span style="color: var(--palette-black-50)">{{ $t("ethereum") }}</span>
    </div>

    <Single
      :w="w"
      :key="w.name"
      class="px-3 b rounded-xl"
      @click="connectEthereum(w)"
      v-for="w in cStore.ethereumWallets"
    />
  </div>
</template>

<script lang="ts" setup>
import Single from "~/components/bridge/elements/dialogs/connect/single.vue";
import {
  connectEthereum,
  useConnectStore,
} from "~/stores/connect/connect";
const cStore = useConnectStore();

</script>

<style lang="scss" scoped>
.b:hover {
  background-color: var(--palette-black-5);
}
</style>