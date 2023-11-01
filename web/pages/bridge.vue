<template>
  <Setting v-if="store.settingMode" />
  <Card v-else />

  <Recent style="position: absolute; right: 0px; bottom: 0px" />
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Card from "~/components/bridge/card.vue";
import Recent from "~/components/bridge/recent.vue";
import Setting from "~/components/bridge/setting.vue";
const store = useBridgeStore();
</script>

<style>
</style>