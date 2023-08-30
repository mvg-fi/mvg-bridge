<template>
  <v-text-field
    v-maska
    data-maska="0.99999999"
    data-maska-tokens="0:\d:multiple|9:\d:optional"
    single-line
    variant="plain"
    @keyup="inputFunc"
    class="text-field font-weight-bold ml-3"
  >
  </v-text-field>
</template>

<script setup lang="ts">
import { useBridgeStore } from "~/stores/bridge/bridge";
import { filterInputEvents } from "~/helpers/utils";
const store = useBridgeStore();

const from = defineProps(["from"]);
const input = from
  ? store.bridgeAmount
  : store.receiveAmount;

let timeout: any = null;
const inputFunc = (event: KeyboardEvent) => {
  if (!filterInputEvents(event)) return;
  clearTimeout(timeout);
  timeout = setTimeout(async function () {
    await console.log();
  }, 500);
};
</script>

<style lang="scss" scoped>
.text-field {
  box-shadow: none !important;
}
</style>