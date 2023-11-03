<template>
  <v-row class="d-flex flex-row hero-height w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 break-text-special mb-4">
        {{ $t("safe_decentralized") }}
      </span>

      <span class="h6m mt-2 break-text">
        {{ $t("safe_decentralized_intro") }}
      </span>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <div style="width: 400px; height: 400px">
        <v-img :src="security" style="width: 400px; height: 400px" />
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import security from "@/assets/images/intro/security.svg";
</script>

<style scoped>
.break-text-special {
  margin-left: calc(600px - 488px);
  width: 600px;
  word-break: break-word;
}
</style>