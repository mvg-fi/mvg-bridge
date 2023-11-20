<template>
  <input
    v-model="input"
    v-maska:[options]
    :placeholder="$t('placeholder')"
    class="text-field font-weight-bold w-100 h5m mr-3 text-black"
  />
<!--     @keyup="inputFunc"  -->
</template>

<script setup lang="ts">
import { vMaska } from "maska";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { filterInputEvents } from "~/helpers/utils";

const store = useBridgeStore();
const options = {
  mask: "A.BBBBBBBB",
  tokens: {
    A: { pattern: /[0-9]/, multiple: true },
    B: { pattern: /[0-9]/, optional: true },
  },
};

const props = defineProps(["from"]);
const input = computed({
  get(){
    return props.from ? store.bridgeAmount : store.receiveAmount
  },
  set(value){
    if (props.from) store.bridgeAmount = value
    else store.receiveAmount = value
  },
})
// const input = ;

// let timeout: any = null;
// const inputFunc = (event: KeyboardEvent) => {
//   if (!filterInputEvents(event)) return;
//   clearTimeout(timeout);
//   timeout = setTimeout(async function () {
//     store.setAmount(props.from ? true : false, input);
//   }, 500);
// };
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
</style>