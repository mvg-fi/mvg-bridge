<template>
  <Setting v-if="store.settingMode"/>
  <Card v-else/>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge"
import Card from "~/components/bridge/card.vue"
import Setting from "~/components/bridge/setting.vue"
const store = useBridgeStore()
</script>

<style>

</style>