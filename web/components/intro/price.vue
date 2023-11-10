






























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
          {{ $t("best_price_as_always") }}
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
          {{ $t("best_price_as_always_intro") }}
        </span>
      </div>
    </v-col>
    <v-col
      :cols="xl ? 6 : lg ? 6 : customMobile ? 8 : 4"
      :class="
        clsx('pa-0 d-flex justify-center', customMobile && 'mt-6 d-flex justify-center align-center')
      "
    >
      <div :class="clsx(customMobile ? 'icons-mobile': 'icons')">
        <v-row no-gutters>
          <v-col
            v-for="i in icons"
            :key="i"
            cols="4"
            class="d-flex align-center justify-center"
          >
            <v-icon :size="customMobile ? 54 : 80" :class="clsx(customMobile? 'my-2': 'my-3')">
              <v-img :src="i" />
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
import bg from "~/assets/images/intro/exchanges/bigone.png";
import bi from "~/assets/images/intro/exchanges/binance.png";
import by from "~/assets/images/intro/exchanges/bybit.png";
import ex from "~/assets/images/intro/exchanges/exinone.png";
import ga from "~/assets/images/intro/exchanges/gate.png";
import hb from "~/assets/images/intro/exchanges/huobi.png";
import mv from "~/assets/images/intro/exchanges/mvg.png";
import ok from "~/assets/images/intro/exchanges/okx.png";
import pa from "~/assets/images/intro/exchanges/pando.png";

const icons = [bi, ok, ga, by, bg, hb, pa, ex, mv];
const customMobile = computed(() => width.value <= 800);
</script>

<style scoped>
.icons {
  width: 300px; 
  height: 312px
}
.icons-mobile {
  width: 224px;
}
</style>