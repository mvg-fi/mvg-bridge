<template>
  <div
    class="mx-2 py-2 justify-center align-center rounded-xl bg-background d-flex align-center"
    style="height: 48px"
  >
    <!-- <v-icon style="width: 20px; height: 20px" class="mx-3">
      <InboxArrowDownIcon />
    </v-icon> -->
    <input
      class="search w-100 mx-3"
      :placeholder="$t('receiver_placeholder')"
    />
    <v-icon
      style="width: 24px; height: 24px"
      @click.stop="console.log()"
      class="mx-2"
    >
      <XMarkIcon />
    </v-icon>
  </div>
</template>

<script setup>
import { InboxArrowDownIcon, XMarkIcon } from "@heroicons/vue/24/outline";
</script>

<style>
.search:focus {
  outline: none;
}
.search {
  font-size: 14px;
  height: 28px;
}

.search::placeholder{
  color: grey;
}

</style>