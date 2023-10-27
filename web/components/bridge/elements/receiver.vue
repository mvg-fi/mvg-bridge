<template>
  <div
    class="mx-2 justify-center align-center rounded-xl d-flex align-center"
    style="height: 48px; background-color: var(--palette-black-5);"
  >
    <input
      v-model="store.receiver"
      class="search w-100 ml-6"
      :placeholder="$t('receiver_placeholder')"
    />
    <v-icon
      style="width: 24px; height: 24px"
      @click.stop="clearReceiver"
      class="mx-4"
    >
      <XMarkIcon />
    </v-icon>
  </div>
</template>

<script setup>
import { InboxArrowDownIcon, XMarkIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore()

const clearReceiver = () => {
  store.clearReceiver()
}
</script>

<style scoped>
.search:focus {
  outline: none;
}
.search {
  height: 56px;
}

.search::placeholder{
  color: grey;
}

</style>