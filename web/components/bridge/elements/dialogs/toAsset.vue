<template>
  <v-dialog v-model="store.toDialog" width="auto">
    <v-card>
      <Title from="to"/>
      
    </v-card>
  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from '~/stores/bridge/bridge';
import Title from "~/components/bridge/elements/dialogs/dialog/assets"
const store = useBridgeStore()
</script>

<style lang="scss" scoped>
</style>