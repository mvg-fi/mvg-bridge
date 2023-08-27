<template>
  <v-overlay
    location="bottom end"
    location-strategy="connected"
    scrim="false"
    transition="scale-transition"
    opacity="0.1"
  >
    <template v-slot:activator="{ props }">
      <v-btn size="large" v-bind="props" :class="NetworkClass">
        <v-icon class="mr-1">
          <v-img :src="currentNetwork.icon"></v-img>
        </v-icon>
        <v-icon>
          <ChevronDownIcon class="network-icon" />
        </v-icon>
      </v-btn>
    </template>

    <v-card class="network-card pa-2">
      <button
        :key="n.evmId"
        @click="()=>{currentNetwork=n}"
        v-for="n in Networks"
        class="d-flex justify-start align-center network-item-btn pa-2"
      >
        <v-icon class="mr-3 network-icon">
          <v-img :src="n.icon" />
        </v-icon>
        <span class="text-capitalize font-weight-light" style="font-size: 14px;">
          {{ n.name }}
        </span>

        <v-spacer />
        <v-icon v-if="n.evmId == currentNetwork.evmId">
          <CheckIcon style="width: 14px; height: 14px; color: blue" />
        </v-icon>
      </button>
    </v-card>
  </v-overlay>
</template>

<script setup>
import clsx from "clsx";
import { Networks } from "~/helpers/networks";
import ethereum from "~/assets/images/networks/ethereum.png";
import { ChevronDownIcon, CheckIcon } from "@heroicons/vue/20/solid";

const currentNetwork = ref(Networks[0])
const NetworkClass = computed(() =>
  clsx("bg-transparent", "rounded-pill", "text-capitalize", "font-weight-bold")
);
</script>

<style scoped>
.network-icon {
  width: 20px;
  height: 20px;
}
.v-btn:hover:before {
  background-color: transparent;
}
.network-item-btn {
  width: 240px;
  padding: 0px 8px;
  line-height: 20px;
  border-radius: 12px;
  transition: background-color 250ms ease 0s;
}
.network-item-btn:hover {
  background-color: rgb(228, 232, 244);
}
.network-card {
  margin-top: 5px;
  border-radius: 2vh;
  box-shadow: 0 2px 32px #0003 !important;
}
</style>