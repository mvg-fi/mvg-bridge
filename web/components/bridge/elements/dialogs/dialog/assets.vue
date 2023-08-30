<template>
  <div
    class="w-100 h-100 overflow-y-auto"
    style="min-height: 400px; padding-bottom: 110px;"
  >
    <div class="w-100" v-for="item in Assets" :key="item.asset_id">
      <v-btn
        block
        elevation="0"
        height="44px"
        class="d-flex justify-start px-4"
        @click="store.setAsset(item); console.log('clicked')"
      >
        <v-icon style="width: 32px; height: 32px; position: reactive">
          <v-img :src="item.icon" />
          <v-img :src="item.icon"
            style="
              height: 14px;
              width: 14px;
              position: absolute;
              bottom: 0px;
              left: 0px;
            "
          ></v-img>
        </v-icon>

        <div class="d-flex flex-column justify-start text-start ml-3">
          <span class="font-weight-bold" style="letter-spacing: 0px">
            {{ item.symbol }}
          </span>

          <span class="font-weight-light" style="font-size: 10px; color: grey">
            {{ item.name }}
            <!-- replace with chain name -->
          </span>
        </div>
      </v-btn>
    </div>
  </div>
</template>

<script setup>
import { useBridgeStore } from "~/stores/bridge/bridge";
import Assets from "~/assets/constants/miniumlist.json";
const store = useBridgeStore()
</script>

<style lang="scss" scoped>
* {
  scrollbar-width: none;
}
::-webkit-scrollbar {
  display: none;
}
</style>