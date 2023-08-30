<template>
  <div
    class="pa-2 justify-center align-center rounded-xl bg-secondary d-flex align-center"
    style="height:36px"
  >
    <v-icon style="width:16px; height:16px" class="mx-2"> 
      <SearchIcon />
    </v-icon>
    <input class="search w-100" :placeholder="$t('search_asset_placeholder')" style="height: 28px" />
  </div>
</template>

<script setup>
import { MagnifyingGlassIcon as SearchIcon } from "@heroicons/vue/24/outline";
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