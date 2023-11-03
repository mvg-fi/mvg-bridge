<template>
  <v-row class="d-flex flex-row hero-height w-100" no-gutters>
    <v-col
      cols="6"
      class="d-flex flex-column align-center justify-center mb-16"
    >
      <span class="h2 mb-4 break-text">
        {{ $t("multichain_as_it_should") }}
      </span>

      <span class="h6m mt-2 break-text">
        {{ $t("multichain_as_it_should_intro") }}
      </span>
    </v-col>

    <v-col cols="6" class="d-flex justify-center align-center">
      <div style="width: 300px; height: 312px">
        <v-row no-gutters style="transform: rotate(-45deg);">
          <v-col
            cols="3"
            :key="c"
            v-for="c in cs"
            style="transform: rotate(45deg);"
            class="d-flex align-center justify-center"
          >
            <v-icon size="72" class="pa-2">
              <v-img :src="c.icon" />
            </v-icon>
          </v-col>
        </v-row>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import chains from "~/assets/constants/miniumchainlist.json";
const cs = chains.slice(0, 16);
</script>

<style scoped>
.break-text {
  width: 488px;
  word-break: break-word;
}
</style>