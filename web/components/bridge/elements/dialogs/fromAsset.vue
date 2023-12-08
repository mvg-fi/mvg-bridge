<template>
  <v-dialog
    v-if="!mobile"
    v-model="store.fromDialog"
    :transition="'scale-transition'"
    :class="clsx('d-flex justify-center dialog-blur')"
  >
    <v-card
      :class="
        clsx(
          'select-asset-card align-self-center rounded-xl pt-3 overflow-y-hidden'
        )
      "
      elevation="3"
    >
      <Title :from="true" class="mb-3" />
      <Search class="mx-5" />
      <v-card class="d-flex flex-row mt-5 h-100" elevation="0">
        <Network />
        <Assets :from="true" />
      </v-card>
    </v-card>
  </v-dialog>

  <v-bottom-sheet v-model="store.fromDialog" class="dialog-blur" content-class="elevation-0" v-else>
    <v-card
      :class="
        clsx(
          'align-self-center rounded-t-xl mobile-card pt-3 overflow-y-hidden'
        )
      "
      elevation="3"
    >
      <Title :from="true" class="mb-3" />
      <Search class="mx-5" />
      <v-card class="d-flex flex-row mt-5 h-100" elevation="0">
        <Network />
        <Assets :from="true" />
      </v-card>
    </v-card>
  </v-bottom-sheet>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { VBottomSheet } from "vuetify/labs/VBottomSheet";
import Title from "~/components/bridge/elements/dialogs/asset/title.vue";
import Search from "~/components/bridge/elements/dialogs/asset/search.vue";
import Network from "~/components/bridge/elements/dialogs/asset/network.vue";
import Assets from "~/components/bridge/elements/dialogs/asset/assets.vue";

const store = useBridgeStore();
const { mobile } = useDisplay();
</script>

<style lang="scss">
@import "assets/css/custom";
.select-asset-card {
  height: $select-asset-dialog-height;
  min-width: $select-asset-dialog-width;
}
.dialog-blur {
  backdrop-filter: blur(4px);
}
.mobile-card {
  height: 85vh;
  height: 85svh;
  width: 100vw;
}
</style>
