<template>
  <v-text-field
    single-line
    variant="plain"
    v-model="input"
    v-maska:[options]
    @keyup="inputFunc"
    :placeholder="$t('placeholder')"
    class="text-field font-weight-bold ml-3"
  >
  </v-text-field>
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
    store.setAmount(props.from ? "from" : "to", input);
  }, 500);
};
</script>

<style lang="scss" scoped>
.text-field {
  box-shadow: none !important;
}
.text-field .v-input_slot{
  font-size: 1.875rem !important;
}
</style>