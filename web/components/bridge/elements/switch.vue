<template>
  <div class="w-100 d-flex justify-center">
  <v-icon @click="store.switch()" class="down" size="20px">
    <ArrowSmallDownIcon />
  </v-icon>
</div>
</template>

<script setup>
import { ArrowSmallDownIcon } from "@heroicons/vue/24/outline"
import { useBridgeStore } from "~/stores/bridge/bridge"
const store = useBridgeStore()
</script>

<style lang="scss" scoped>
.down:hover {
	transform: rotate(-180deg);
	transition: transform .5s;
}
.down:not(hover){
  transform: rotate(0deg);
	transition: transform .5s;
}
</style>