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
              : 'd-flex flex-column mb-8'
          )
        "
      >
        <span
          :class="
            clsx(
              xxl
                ? 'h1 break-text-special-xxl'
                : xl
                ? 'h2 break-text-special'
                : lg
                ? 'h3 break-text-special-lg'
                : 'h4 break-text-special-mobile'
            )
          "
        >
          {{ $t("safe_decentralized") }}
        </span>

        <span
          :class="
            clsx(
              'mt-2',
              xxl
                ? 'h5m break-text-special-xxl'
                : xl
                ? 'h6m break-text-special'
                : lg
                ? 'h6m break-text-special-lg mt-4'
                : 'h7m break-text-special-mobile mt-4'
            )
          "
        >
          {{ $t("safe_decentralized_intro") }}
        </span>
      </div>
    </v-col>
    <v-col
      :cols="xl ? 6 : lg ? 6 : customMobile ? 8 : 4"
      :class="clsx('pa-0', customMobile && 'mt-6 d-flex justify-center')"
    >
      <v-img
        :src="security"
        :class="clsx(customMobile ? 'hero-img-mobile' : 'intro-imgs')"
      />
    </v-col>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
const { lg, xl, xxl, width } = useDisplay();
import security from "@/assets/images/intro/security.svg";

const customMobile = computed(() => width.value <= 800);
</script>

<style scoped>
.break-text-special-xxl {
  width: 600px;
  word-break: break-word;
}
.break-text-special {
  width: 600px;
  word-break: break-word;
}
.break-text-special-lg {
  width: 400px;
  word-break: break-word;
}
.break-text-special-mobile {
  width: 310px;
  word-break: break-word;
}
</style>
