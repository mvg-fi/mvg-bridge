<template>
  <v-main>
    <v-app-bar
      elevation="0"
      :class="clsx(!mobile && 'px-8', mobile && 'px-6', 'bg-transparent')"
    >
      <Logo />
      <v-spacer />
      <EnterApp />
    </v-app-bar>
    <slot />
  </v-main>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import Logo from "~/components/toolbar/logo.vue";
import EnterApp from "~/components/toolbar/enterApp.vue";
const { mobile } = useDisplay();
</script>