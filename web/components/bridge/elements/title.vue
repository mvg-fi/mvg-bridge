<template>
  <div class="d-flex justify-space-between">
    <span> {{ $t("bridge") }} </span>
    <v-btn
      icon
      plain
      elevation="0"
      @click="ToggleSettingMode"
      class="bg-transparent setting-btn"
      style="width: 24px; height: 24px"
    >
      <v-icon>
        <Cog8ToothIcon />
      </v-icon>
    </v-btn>
  </div>
</template>

<script setup>
import { Cog8ToothIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore()
const ToggleSettingMode = ()=>{
  store.switchSettingMode()
}
</script>

<style lang="scss" scoped>
.setting-btn:hover{
  transform: rotate(30deg);
  transition: transform .6s;
}
</style>