<template>
  <Setting v-if="store.settingMode" />
  <Card v-else />

  <Recent />

  <FromAssetDialog />
  <ToAssetDialog />
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Card from "~/components/bridge/card.vue";
import Recent from "~/components/bridge/recent.vue";
import Setting from "~/components/bridge/setting.vue";
import FromAssetDialog from "~/components/bridge/elements/dialogs/fromAsset.vue";
import ToAssetDialog from "~/components/bridge/elements/dialogs/toAsset.vue";

const store = useBridgeStore();
</script>

<style>
</style>