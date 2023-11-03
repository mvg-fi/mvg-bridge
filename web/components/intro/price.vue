<template>
  <v-row class="d-flex flex-row hero-height w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 mb-6 break-text">
        {{ $t("best_price_as_always") }}
      </span>

      <span class="h6m mt-2 break-text">
        {{ $t("best_price_as_always_intro") }}
      </span>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <div style="width: 300px; height: 312px;">
        <v-row no-gutters>
          <v-col
            v-for="i in icons"
            :key="i"
            cols="4"
            class="d-flex align-center justify-center "
          >
            <v-icon size="80" class="my-3">
              <v-img :src="i" />
            </v-icon>
          </v-col>
        </v-row>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
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
</script>
