<template>
  <v-dialog
    :fullscreen="mobile"
    v-model="store.toDialog"
    :transition="mobile ? 'slide-y-reverse-transition' : 'scale-transition'"
    :class="
      clsx('d-flex justify-center dialog-blur', mobile && 'mobile-dialog')
    "
  >
    <v-card
      :class="
        clsx(
          'select-asset-card align-self-center rounded-xl pt-3 overflow-y-hidden',
          mobile && 'mobile-card'
        )
      "
      elevation="3"
    >
      <Title :from="false" class="mb-3" />
      <Search class="mx-5" />
      <v-card class="d-flex flex-row mt-5 h-100" elevation="0">
        <Network/>
        <Assets :from="false"/>
      </v-card>
    </v-card>
  </v-dialog>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";
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
  width: $select-asset-dialog-width;
}
.dialog-blur {
  backdrop-filter: blur(4px);
}
</style>