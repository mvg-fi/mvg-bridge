<template>
  <div class="px-3 d-flex align-center">
    <v-col cols="1"/> 

    <v-col cols="10" class="text-center py-0">
      <span class="font-weight-semibold"> {{ $t("select_asset") }}</span>
    </v-col>

    <v-col cols="1" class="py-0">
      <v-btn icon elevation="0" @click="store.mutateDialog(props.from, false)">
        <v-icon>
          <XMarkIcon />
        </v-icon>
      </v-btn>
    </v-col>
  </div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import { XMarkIcon } from "@heroicons/vue/24/solid";
const store = useBridgeStore();
const props = defineProps(["from"]);
</script>

<style lang="scss" scoped>
</style>