<template>
  <div
    class="pa-2 justify-center align-center rounded-xl bg-secondary d-flex align-center"
    style="height: 36px"
  >
    <v-icon style="width: 16px; height: 16px" class="mx-2">
      <SearchIcon />
    </v-icon>
    <input
      class="search w-100"
      style="height: 28px"
      @keyup="inputFunc"
      v-model="input"
      :placeholder="$t('search_asset_placeholder')"
    />
  </div>
</template>

<script setup lang="ts">
import { MagnifyingGlassIcon as SearchIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore()
let timeout: any = null;
const input = ref('')
const inputFunc = (event: KeyboardEvent) => {
  clearTimeout(timeout);
  timeout = setTimeout(async function () {
    store.setSearchAsset(input.value)
  }, 300);
};
</script>

<style lang="scss" scoped>
.search:focus {
  outline: none;
}
.search {
  font-size: 16px;
}
::placeholder {
  opacity: 0.2;
  font-size: 14px;
}
</style>