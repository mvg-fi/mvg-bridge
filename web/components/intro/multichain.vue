<template>
  <v-row
    :class="
      clsx(
        'w-100 hero-height',
        customMobile
          ? 'flex-column align-center justify-center text-center'
          : 'flex-row desktop-padding align-center'
      )
    "
    no-gutters
  >
    <v-col :cols="xl ? 6 : lg ? 6 : customMobile ? 12 : 8" class="pa-0">
      <div
        :class="
          clsx(
            customMobile
              ? 'd-flex flex-column align-center text-center'
              : 'mb-8'
          )
        "
      >
        <span
          :class="
            clsx(
              xxl
                ? 'h1 break-text-xxl'
                : xl
                ? 'h2 break-text'
                : lg
                ? 'h3 break-text-lg'
                : 'h4 break-text-mobile'
            )
          "
        >
          {{ $t("multichain_as_it_should") }}
        </span>

        <span
          :class="
            clsx(
              'mt-2',
              xxl
                ? 'h5m break-text-xxl'
                : xl
                ? 'h6m break-text'
                : lg
                ? 'h6m break-text-lg mt-4'
                : 'h7m break-text-mobile mt-4'
            )
          "
        >
          {{ $t("multichain_as_it_should_intro") }}
        </span>
      </div>
    </v-col>
    <v-col
      :cols="xl ? 6 : lg ? 6 : customMobile ? 8 : 4"
      :class="
        clsx('pa-0  d-flex justify-center align-center', customMobile && 'mt-6 d-flex justify-center align-center mt-12')
      "
    >
      <div :class="clsx(customMobile ? 'icons-mobile': 'icons')">
        <v-row no-gutters style="transform: rotate(-45deg)">
          <v-col
            cols="3"
            :key="c"
            v-for="c in cs"
            style="transform: rotate(45deg)"
            class="d-flex align-center justify-center"
          >
            <v-icon :size="customMobile ? 48 : 72" :class="customMobile ? 'pa-1': 'pa-2'">
              <v-img :src="c.icon" />
            </v-icon>
          </v-col>
        </v-row>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
const { lg, xl, xxl, width } = useDisplay();
import chains from "~/assets/constants/miniumchainlist.json";
const cs = chains.slice(0, 16);
const customMobile = computed(() => width.value <= 800);
</script>

<style scoped>
.icons {
  width: 288px; 
}
.icons-mobile {
  width: 192px;
}
</style>