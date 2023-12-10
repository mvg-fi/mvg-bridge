<template>
  <div
    :class="
      clsx(
        'mx-2 rounded-xl d-flex align-center justify-center custom-input-bg',
        mobile ? 'recipient-mobile' : 'recipient'
      )
    "
  >
    <input
      v-model="store.receiver"
      :placeholder="$t('receiver_placeholder')"
      :class="clsx('w-100', mobile ? 'ml-4 search-mobile' : 'ml-6 search')"
    />
    <v-icon :size="mobile ? 16 : 24" @click.stop="clearReceiver" class="mx-4">
      <XMarkIcon />
    </v-icon>
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();
const { mobile } = useDisplay();

const clearReceiver = () => {
  store.clearReceiver();
};
</script>

<style scoped>
.search:focus {
  outline: none;
}
.search {
  height: 56px;
  font-size: 16px;
}
.search-mobile {
  height: 36px;
  font-size: 14px;
}
.recipient {
  height: 48px;
}
.recipient-mobile {
  height: 40px;
}
.search::placeholder {
  color: grey;
}
</style>