<template>
  <v-dialog
    v-model="store.fromDialog"
    class="d-flex justify-center select-asset-dialog"
  >
    <v-card
      class="select-asset-card rounded-xl pt-3 overflow-y-hidden"
      elevation="3"
    >
      <Title :from="true" class="mb-3" />
      <Search class="mx-5" />
      <v-card class="d-flex flex-row mt-5 h-100" elevation="0">
        <Network/>
        <Assets :from="true"/>
      </v-card>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Title from "~/components/bridge/elements/dialogs/dialog/title.vue";
import Search from "~/components/bridge/elements/dialogs/dialog/search.vue";
import Network from "~/components/bridge/elements/dialogs/dialog/network.vue";
import Assets from "~/components/bridge/elements/dialogs/dialog/assets.vue";
const store = useBridgeStore();
</script>

<style lang="scss">
@import "assets/css/custom";
.select-asset-dialog {
  width: $select-asset-dialog-width;
  height: $select-asset-dialog-height;
}
.select-asset-card {
  height: $select-asset-dialog-height;
}
</style>