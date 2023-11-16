<template>
  <v-row no-gutters class="fill-height d-flex justify-center w-100">
    <div>
      <v-card
        :class="
          clsx(
            !mobile && 'bridge-card mt-10',
            mobile && 'mobile-bridge-card mt-6',
            'rounded-xl pa-5 pb-12 d-flex flex-column overflow-y-auto'
          )
        "
        elevation="0"
      >
        <Title />
        <Custom class="mt-9" />
        <v-divider class="my-4" />
        <Dex />
        <v-divider class="my-4" />
        <Privacy />
      </v-card>
    </div>
  </v-row>
</template>

<script setup>
import Title from "~/components/bridge/settings/title.vue";
import Custom from "~/components/bridge/settings/custom.vue";
import Dex from "~/components/bridge/settings/dex.vue";
import Privacy from "~/components/bridge/settings/privacy.vue";
import { useDisplay } from "vuetify";
import clsx from "clsx";

const { mobile } = useDisplay();
</script>

<style scoped>
.maxheight {
  max-height: 560px;
}
</style>