<template>
  <div class="d-flex">
    <span class="flex-grow-1"> {{ $t("bridge") }} </span>
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
const ToggleSettingMode = ()=>{
  // toggle switching setting
}
</script>

<style lang="scss" scoped>
.setting-btn:hover{
  transform: rotate(-30deg);
  transition: transform .6s;
}
</style>