<template>
  <v-main>
    <v-app-bar
      elevation="0"
      :class="clsx(!customMobile ? 'px-8' : 'px-6', 'bg-transparent')"
    >
      <Logo />
      <v-spacer />
      <EnterApp v-if="!customMobile" />
      <MobileMenu v-else />
    </v-app-bar>
    <slot />
  </v-main>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import Logo from "~/components/toolbar/logo.vue";
import EnterApp from "~/components/toolbar/enterApp.vue";
import MobileMenu from "~/components/toolbar/mobileMenu.vue";
const { width } = useDisplay();
const customMobile = computed(() => width.value <= 800);
</script>