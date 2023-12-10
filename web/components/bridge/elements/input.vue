<template>
  <div class="mx-2">
    <div
      :class="
        clsx(
          'd-flex',
          'flex-column',
          'custom-input-bg',
          from ? (mobile ? 'input-top-bg-mobile' : 'input-top-bg') : '',
          !from ? (mobile ? 'input-down-bg-mobile' : 'input-down-bg') : ''
        )
      "
    >
      <div :class="clsx(mobile ? 'ml-4 mt-3' : 'ml-6 mt-3', 'h7m')">
        <span style="opacity: 50%">
          {{ from ? $t("you_pay") : $t("you_receive") }}
        </span>
      </div>

      <div :class="clsx(mobile ? 'ml-4' : 'ml-6', 'd-flex flex-row')">
        <TextField :from="from" />
        <Asset :from="from" />
      </div>

      <div :class="clsx(mobile ? 'ml-4 mb-3' : 'ml-6 mb-3', 'h7m')">
        <span
          style="opacity: 50%"
          v-if="from ? store.bridgeAmount : store.receiveAmount"
        >
          {{ from ? formatUSMoney(store.fromTotalPrice) : formatUSMoney(store.toTotalPrice) }}
        </span>
        <span style="opacity: 0" v-else> $ </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import Asset from "~/components/bridge/elements/input/asset.vue";
import TextField from "~/components/bridge/elements/input/textfield.vue";
import { formatUSMoney } from "~/helpers/utils";
import { useBridgeStore } from "~/stores/bridge/bridge";
const props = defineProps(["from"]);
const from = toRef(props, "from");
const store = useBridgeStore();
const { mobile } = useDisplay();

watch(
  store.bridgeAmount,
  () => {
    store.fetchUSD(true);
  },
  { immediate: true }
);
watch(
  store.receiveAmount,
  () => {
    store.fetchUSD(false);
  },
  { immediate: true }
);
</script>

<style lang="scss" scoped>
.input-top-bg {
  border-top-left-radius: 32px;
  border-top-right-radius: 32px;
}
.input-down-bg {
  border-bottom-left-radius: 32px;
  border-bottom-right-radius: 32px;
}
.input-top-bg-mobile {
  border-top-left-radius: 24px;
  border-top-right-radius: 24px;
}
.input-down-bg-mobile {
  border-bottom-left-radius: 24px;
  border-bottom-right-radius: 24px;
}
</style>
