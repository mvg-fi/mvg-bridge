<template>
  <Main />
  <Fast class="bg-third" />
  <Price />
  <Privacy class="bg-third" />
  <Multichain />
  <Safe class="bg-third" />
  <JoinUs />
</template>

<script setup>
import Main from "~/components/intro/main.vue";
import Fast from "~/components/intro/fast.vue";
import Price from "~/components/intro/price.vue";
import Privacy from "~/components/intro/privacy.vue";
import Multichain from "~/components/intro/multichain.vue";
import Safe from "~/components/intro/safe.vue";
import JoinUs from "~/components/intro/joinus.vue";

definePageMeta({
  layout: "intro",
});
</script>

<style>
</style>