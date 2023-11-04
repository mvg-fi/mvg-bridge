<template>
  <v-row class="d-flex flex-row w-100 hero-height" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 break-text mb-2">
        {{ $t("making_blockchain_borderless") }}
      </span>

      <span class="h6m mt-2 break-text mb-10">
        {{ $t("making_blockchain_borderless_intro") }}
      </span>

      <div class="d-flex align-start" style="width: 488px">
        <NuxtLink to="/bridge">
          <v-btn class="enter-app-btn" elevation="8">
            <span class="text-background h6">{{ $t("enter_app") }} </span>
          </v-btn>
        </NuxtLink>
      </div>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center mb-16">
      <v-img :src="main" style="width: 666px; height: 436px;" />
    </v-col>
  </v-row>
</template>

<script setup>
import main from "@/assets/images/intro/hero-main.png"
</script>

<style>
.hero-height {
  height: calc(100vh - 16px);
}
.break-text {
  width: 488px;
  word-break: break-word;
}
</style>

<style scoped>
.enter-app-btn {
  min-width: 184px;
  height: 64px;
  border-radius: 32px;
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
}
</style>