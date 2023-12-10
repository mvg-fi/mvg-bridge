<template>
  <div class=" mx-2">
    <div class="w-100" style="position: relative; top: -24px; height: 0px;">
    <div class="d-flex flex-row align-center">
      <hr class="line" />
      <v-icon @click="store.switch()" class="down custom-input-bg pa-2" size="40px">
        <ArrowSmallDownIcon />
      </v-icon>
      <hr class="line" />
    </div>
  </div>
  </div>
</template>

<script setup>
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { ArrowSmallDownIcon } from "@heroicons/vue/24/outline";
const store = useBridgeStore();
</script>

<style lang="scss" scoped>
.down:hover {
  transform: rotate(-180deg);
  transition: transform 0.5s;
}
.down:not(hover) {
  transform: rotate(0deg);
  transition: transform 0.5s;
}
.down {
  border-radius: 100%;
  border: 3px solid var(--palette-background-0);
}
.line {
  border: 2px solid var(--palette-background-0);
  width: 258px;
}
</style>