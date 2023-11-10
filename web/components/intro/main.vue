<template>
  <v-row
    :class="
      clsx(
        'w-100 hero-height',
        customMobile
          ? 'flex-column align-center text-center'
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
              ? 'd-flex flex-column align-center text-center mt-12'
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
          {{ $t("making_blockchain_borderless") }}
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
          {{ $t("making_blockchain_borderless_intro") }}
        </span>
      </div>

      <div
        class="d-flex align-start mt-8"
        style="width: 488px"
        v-if="!customMobile"
      >
        <NuxtLink to="/bridge">
          <v-btn class="enter-app-btn" elevation="8">
            <span class="text-background h6">{{ $t("enter_app") }} </span>
          </v-btn>
        </NuxtLink>
      </div>
    </v-col>
    <v-col
      :cols="xl ? 6 : lg ? 6 : customMobile ? 8 : 4"
      :class="clsx('pa-0', customMobile && 'mt-6 d-flex justify-center')"
    >
      <v-img
        :src="main"
        :class="clsx(customMobile ? 'hero-img-mobile' : 'hero-img')"
      />
    </v-col>
    <div
      class="d-flex justify-center mt-8"
      style="width: 300px"
      v-if="customMobile"
    >
      <NuxtLink to="/bridge">
        <v-btn class="enter-app-btn-mobile" elevation="8">
          <span class="text-background h7">{{ $t("enter_app") }} </span>
        </v-btn>
      </NuxtLink>
    </div>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import main from "@/assets/images/intro/hero-main.png";
const { lg, xl, xxl, width } = useDisplay();

const customMobile = computed(() => width.value <= 800);
</script>

<style>
.hero-height {
  min-height: 780px;
  max-height: 1080px;
}
.desktop-padding {
  padding-left: 10%;
  padding-right: 10%;
}
.hero-img {
  min-width: 240px;
  max-width: 666px;
  max-height: 436px;
}
.hero-img-mobile {
  min-width: 300px;
  min-height: 200px;
  width: 300px;
  height: 200px;
}
.break-text {
  width: 488px;
  display: block;
  word-break: break-word;
}
.break-text-lg {
  width: 320px;
  display: block;
  word-break: break-word;
}
.break-text-mobile {
  width: 240px;
  display: block;
  word-break: break-word;
}
.break-text-xxl {
  width: 680px;
  display: block;
  word-break: break-word;
}
</style>

<style scoped>
.enter-app-btn {
  min-width: 184px;
  height: 64px;
  border-radius: 32px;
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
}
.enter-app-btn-mobile {
  min-width: 144px;
  height: 48px;
  border-radius: 32px;
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
}
</style>