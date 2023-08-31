<template>
  <v-dialog
    v-model="store.toDialog"
    class="d-flex justify-center select-asset-dialog"
  >
    <v-card
      class="select-asset-card rounded-xl py-2 pt-3 overflow-y-hidden"
      elevation="3"
    >
      <Title :from="false" class="mb-3" />
      <Search class="mx-5" />
      <div class="fill-height d-flex flex-row mt-5">
        <Network class="px-3 mx-2" />
        <v-divider vertical></v-divider>
        <Assets :from="false"/>
      </div>
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