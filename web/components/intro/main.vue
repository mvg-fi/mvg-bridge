<template>
  <v-row class="d-flex flex-row h-100 w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 break-text">
        {{ $t("making_blockchain_borderless") }}
      </span>

      <span class="h5m mt-2 break-text">
        {{ $t("making_blockchain_borderless_intro") }}
      </span>

      <div class="d-flex align-start mt-8" style="width: 488px">
        <NuxtLink to="/bridge">
          <v-btn class="enter-app-btn" elevation="8">
            <span class="text-background h5">{{ $t("enter_app") }} </span>
          </v-btn>
        </NuxtLink>
      </div>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <v-img :src="main" style="width: 666px; height: 436px;" />
    </v-col>
  </v-row>
</template>

<script setup>
import main from "@/assets/images/intro/hero-main.png"
</script>

<style scoped>
.break-text {
  width: 488px;
  word-break: break-word;
}
.enter-app-btn {
  width: 288px;
  height: 84px;
  border-radius: 100px;
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
}
</style>