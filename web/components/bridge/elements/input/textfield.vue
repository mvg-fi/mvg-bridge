<template>
  <input
    v-model="input"
    v-maska:[options]
    @keyup="inputFunc"
    :placeholder="$t('placeholder')"
    class="text-field font-weight-bold w-100 h5m mr-3"
  />
</template>

<script setup lang="ts">
import { useBridgeStore } from "~/stores/bridge/bridge";
import { filterInputEvents } from "~/helpers/utils";
import { vMaska } from "maska";
const store = useBridgeStore();
const options = {
  mask: "A.BBBBBBBB",
  tokens: {
    A: { pattern: /[0-9]/, multiple: true },
    B: { pattern: /[0-9]/, optional: true },
  },
};

const props = defineProps(["from"]);
const input = props.from ? ref(store.bridgeAmount) : ref(store.receiveAmount);

let timeout: any = null;
const inputFunc = (event: KeyboardEvent) => {
  if (!filterInputEvents(event)) return;
  clearTimeout(timeout);
  timeout = setTimeout(async function () {
    store.setAmount(props.from ? true : false, input);
  }, 500);
};
</script>

<style lang="scss" scoped>
.text-field {
  color: var(--palette-black-100);
  box-shadow: none !important;
  height: 64px;
}

input::placeholder {
  color: var(--palette-black-25) !important;
}

.text-field::placeholder{
  color: grey;
}
</style>