<template>
  <div class=" mx-2">
    <div class="w-100" style="position: relative; top: -24px; height: 0px;">
    <div class="d-flex flex-row align-center">
      <hr class="line" />
      <v-icon @click="store.switch()" class="down pa-3" size="48px">
        <ArrowSmallDownIcon />
      </v-icon>
      <hr class="line" />
    </div>
  </div>
  </div>
</template>

<script setup>
import { ArrowSmallDownIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
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
  background-color: var(--palette-black-5);
  border: 3px solid var(--palette-background-1);
}
.line {
  border: 2px solid var(--palette-background-1);
  width: 258px;
}
</style>