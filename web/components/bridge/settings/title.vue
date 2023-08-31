<template>
  <div>
  <v-row>
    <v-col cols="2">
      <v-icon
        style="width: 24px; height: 24px"
        @click.stop="store.switchSettingMode()"
      >
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" height="24" width="24"><path xmlns="http://www.w3.org/2000/svg" d="M14.7071 5.29289C15.0976 5.68342 15.0976 6.31658 14.7071 6.70711L9.41421 12L14.7071 17.2929C15.0976 17.6834 15.0976 18.3166 14.7071 18.7071C14.3166 19.0976 13.6834 19.0976 13.2929 18.7071L7.29289 12.7071C6.90237 12.3166 6.90237 11.6834 7.29289 11.2929L13.2929 5.29289C13.6834 4.90237 14.3166 4.90237 14.7071 5.29289Z"></path></svg>
      </v-icon>
    </v-col>
    <v-col cols="8" class="text-center font-weight-semibold">
      {{ $t("advanced_settings") }}
    </v-col>
    <v-col cols="2" />
  </v-row>
</div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
const store = useBridgeStore();
</script>

<style lang="scss" scoped>
</style>