<template>
  <v-row no-gutters class="fill-height d-flex justify-center">
    <div>
      <v-card
        :class="
          clsx(
            !mobile && 'bridge-card mt-10',
            mobile && 'mobile-bridge-card mt-6',
            'rounded-xl pa-0 d-flex flex-column'
          )
        "
        elevation="0"
      >
        <Title class="mx-6 mt-4" />

        <div class="mt-4">
          <Input :from="true" />
          <Switch />
          <Input :from="false" />
        </div>

        <Receiver class="my-2" />
        <Fees class="my-2" />
        <Confirm class="my-2" />
      </v-card>
    </div>
  </v-row>
</template>

<script setup>
import Title from "~/components/bridge/elements/title.vue";
import Input from "~/components/bridge/elements/input.vue";
import Switch from "~/components/bridge/elements/switch.vue";
import Receiver from "~/components/bridge/elements/receiver.vue";
import Fees from "~/components/bridge/elements/fees.vue";
import Confirm from "~/components/bridge/elements/confirm.vue";
import { useDisplay } from "vuetify";
import clsx from "clsx";

const { mobile } = useDisplay();
</script>

<style lang="scss">
@import "assets/css/custom";
.bridge-card {
  width: 540px;
  box-shadow: 0 1px 8px #0003 !important;
}
.mobile-bridge-card {
  width: calc(100vw - 16px);
  box-shadow: 0 1px 8px #0003 !important;
}
</style>