<template>
  <div class="d-flex justify-center">
    <Bubble :text="t('sent')" :state="2" />
    <div class="d-flex align-center mx-6">
      <v-icon size="20">
        <ArrowRightIcon />
      </v-icon>
    </div>
    <Bubble :text="t('processing')" :state="1" />
    <div class="d-flex align-center mx-6">
      <v-icon size="20">
        <ArrowRightIcon />
      </v-icon>
    </div>
    <Bubble :text="t('finished')" :state="0" />
  </div>
</template>

<script setup>
import { useI18n } from "vue-i18n";
import { ArrowRightIcon } from "@heroicons/vue/24/solid";
import Bubble from "~/components/status/status/bubble.vue";

const { t } = useI18n();
</script>

<style>
</style>