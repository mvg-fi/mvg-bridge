<template>
  <v-overlay location="bottom end" location-strategy="connected" scrim="false" transition="scale-transition" opacity="0.1">
    <template v-slot:activator="{ props }">
      <v-btn icon class="menu-btn" v-bind="props">
        <v-icon>
          <EllipsisHorizontalIcon class="menu-icon" />
        </v-icon>
      </v-btn>
    </template>
    <v-card class="menu-card pa-6">
      <v-list />
    </v-card>
  </v-overlay>
</template>

<script setup>
import { EllipsisHorizontalIcon } from "@heroicons/vue/20/solid";
</script>

<style lang="scss" scoped>
.menu-icon {
  width: 22px;
  height: 22px;
}
.menu-card {
  margin-top: 5px;
  border-radius: 2vh;
  min-width: 240px;
  min-height: 300px;
  box-shadow: 0 2px 32px #0003 !important;
}
</style>