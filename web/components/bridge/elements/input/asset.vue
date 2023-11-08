<template>
  <div :class="clsx('d-flex align-center', mobile ? 'mr-4' : 'mr-6')">
    <v-btn
      :class="
        clsx(
          'rounded-pill d-flex flex-row',
          mobile ? 'button-mobile px-4' : 'button-normal px-6 py-2'
        )
      "
      elevation="0"
      size="x-large"
      @click="store.mutateDialog(props.from, true)"
    >
      <v-icon :size="mobile ? 24 : 36" :class="mobile ? 'mr-2' : 'mr-4'">
        <v-img :src="from ? store.fromAsset.icon : store.toAsset.icon" />
        <v-img
          :src="from ? store.fromAsset.chain_icon : store.toAsset.chain_icon"
          v-if="
            from
              ? store.fromAsset.asset_id != store.fromAsset.chain_id
              : store.toAsset.asset_id != store.toAsset.chain_id
          "
          :class="
            clsx(mobile ? 'chain-icon-absolute-mobile' : 'chain-icon-absolute')
          "
        ></v-img>
      </v-icon>

      <div
        :class="
          clsx(
            mobile ? 'mr-2' : 'mr-4',
            'd-flex flex-column align-center justify-center'
          )
        "
      >
        <span style="" :class="clsx(mobile ? 'h7m' : 'h6m')">
          {{ from ? store.fromAsset.symbol : store.toAsset.symbol }}
        </span>
        <span :class="clsx(mobile ? 'chain-name-mobile' : 'chain-name')">
          {{ from ? store.fromAsset.chain_name : store.toAsset.chain_name }}
        </span>
      </div>

      <v-icon :size="mobile ? 12 : 20">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          height="24"
          width="24"
        >
          <path
            xmlns="http://www.w3.org/2000/svg"
            d="M17 10L12 16L7 10H17Z"
            fill="#0D0D0D"
          />
        </svg>
      </v-icon>
    </v-btn>
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";

const { mobile } = useDisplay();
const store = useBridgeStore();
const props = defineProps(["from"]);

console.log(store.fromAsset);
console.log(store.toAsset);
</script>

<style lang="scss" scoped>
.chain-icon-absolute {
  height: 14px;
  width: 14px;
  position: absolute;
  bottom: 0px;
  left: 0px;
}
.chain-icon-absolute-mobile {
  height: 8px;
  width: 8px;
  position: absolute;
  bottom: 0px;
  left: 0px;
}
.chain-name {
  color: var(--palette-black-50);
  font-size: 12px;
  line-height: 1rem;
}
.chain-name-mobile {
  color: var(--palette-black-50);
  font-size: 10px;
  line-height: 0.75rem;
}
.button-normal {
  height: 64px;
}
.button-mobile {
  height: 48px;
}
</style>