<template>
  <v-row class="d-flex flex-row hero-height w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 break-text mb-4">
        {{ $t("privacy_guaranteed") }}
      </span>

      <span class="h6m mt-2 break-text">
        {{ $t("privacy_guaranteed_intro") }}
      </span>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <div style="width: 400px; height: 400px">
        <v-img :src="safe" style="width: 400px; height: 400px" />
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import safe from "@/assets/images/intro/safe.svg";
</script>
