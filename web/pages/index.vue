<template>
  <Main />
  <Fast class="custom-input-bg" />
  <Price />
  <Privacy class="custom-input-bg" />
  <Multichain />
  <Safe class="custom-input-bg" />
  <JoinUs />
  <Footer class="custom-input-bg"/>
</template>

<script setup>
import Main from "~/components/intro/main.vue";
import Fast from "~/components/intro/fast.vue";
import Price from "~/components/intro/price.vue";
import Privacy from "~/components/intro/privacy.vue";
import Multichain from "~/components/intro/multichain.vue";
import Safe from "~/components/intro/safe.vue";
import JoinUs from "~/components/intro/joinus.vue";
import Footer from "~/components/footer/footer.vue";

definePageMeta({
  layout: "intro",
});
</script>

<style>
</style>