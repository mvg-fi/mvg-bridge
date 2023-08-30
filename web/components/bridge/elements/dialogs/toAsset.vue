<template>
  <v-dialog v-model="store.state.bridge.fromDialog" width="auto">
    <v-card>
      <v-card-text> </v-card-text>
      <v-card-actions>
        <v-btn
          color="primary"
          block
          @click="store.mutateDialog('to', false)"
          >Close Dialog</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { useBridgeStore } from '~/stores/bridge/bridge';
const store = useBridgeStore()
</script>

<style lang="scss" scoped>
</style>