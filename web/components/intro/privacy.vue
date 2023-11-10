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
    <v-col :cols="xl ? 6 : lg ? 6 : customMobile ? 12: 8" class="pa-0">
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
        {{ $t("privacy_guaranteed") }}
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
        {{ $t("privacy_guaranteed_intro") }}
      </span>
      </div>
    </v-col>
    <v-col
      :cols="xl ? 6 : lg ? 6 : customMobile ? 8 : 4"
      :class="clsx('pa-0', customMobile && 'mt-6 d-flex justify-center')"
    >
      <v-img
        :src="safe"
        :class="clsx(customMobile ? 'hero-img-mobile' : 'hero-img')"
      />
    </v-col>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
const { lg, xl, xxl, width } = useDisplay();
import safe from "@/assets/images/intro/safe.svg";

const customMobile = computed(() => width.value <= 800);
</script>