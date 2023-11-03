<template>
  <v-main>
    <v-app-bar elevation="0" class="px-8 bg-transparent">
      <Logo />
      <v-spacer />
      <EnterApp />
    </v-app-bar>
    <slot />
  </v-main>
</template>

<script setup>
import Logo from "~/components/toolbar/logo.vue";
import EnterApp from "~/components/toolbar/enterApp.vue";
</script>