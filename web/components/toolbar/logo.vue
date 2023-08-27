<template>
  <div class="d-flex align-center" style="cursor:pointer;" @click="navigateTo('/')">
      <v-app-bar-nav-icon>
        <v-img height="28px" width="28px" :src="Icon"></v-img>
      </v-app-bar-nav-icon>
      <v-app-bar-title :text="AppName" class="logo-text unselectable" />
    </div>
</template>

<script lang="ts" setup>
import Icon from "assets/images/mvg.png";
import { AppName } from "~/helpers/constants";
</script>

<style scoped>
.logo-text {
  font-size: 22px;
  font-weight: 600;
  font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
    Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, sans-serif,
    "Apple Color Emoji", "Segoe UI Emoji", Segoe UI Symbol, "Noto Color Emoji";
}
</style>