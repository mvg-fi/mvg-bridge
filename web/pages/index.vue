<template>
  <Main />
  <Fast />
  <Price />
  <Privacy />
  <Multichain />
  <Safe />
</template>

<script setup>
import Main from "~/components/intro/main.vue";
import Fast from "~/components/intro/fast.vue";
import Price from "~/components/intro/price.vue";
import Privacy from "~/components/intro/privacy.vue";
import Safe from "~/components/intro/safe.vue";

definePageMeta({
  layout: "intro",
});
</script>

<style>
</style>