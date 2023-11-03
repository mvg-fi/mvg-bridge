<template>
  <v-row class="d-flex flex-row hero-height w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 break-text">
        {{ $t("fast_like_a_rocket") }}
      </span>

      <span class="h6m mt-2 break-text">
        {{ $t("fast_like_a_rocket_intro") }}
      </span>

      <!-- <div class="d-flex align-start mt-8" style="width: 488px">
        <NuxtLink to="/bridge">
          <v-btn class="enter-app-btn" elevation="8">
            <span class="text-background h5">{{ $t("enter_app") }} </span>
          </v-btn>
        </NuxtLink>
      </div> -->
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <div style="width: 400px; height: 400px">
        <v-img :src="fast" style="width: 400px; height: 400px" />
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import fast from "@/assets/images/intro/fast.svg";
</script>