<template>
  <v-overlay
    v-if="!mobile"
    location="bottom end"
    location-strategy="connected"
    scrim="false"
    transition="scale-transition"
    opacity="0.1"
  >
    <template v-slot:activator="{ props }">
      <v-btn icon class="menu-btn" v-bind="props">
        <v-icon>
          <EllipsisHorizontalIcon class="menu-icon" />
        </v-icon>
      </v-btn>
    </template>
    <v-sheet
      class="menu-card py-4 pt-3 rounded-xl"
      v-if="bStore.menuState == 0"
    >
      <!-- Settings -->
      <Language />

      <!-- Dark mode -->
      <Dark />

      <!-- Recent exchanges -->
      <RecentExchanges />

      <v-divider class="mx-5 my-4" style="color: var(--palette-black-50)" />

      <!-- Links -->
      <Links />
    </v-sheet>

    <v-sheet
      class="menu-card py-4 rounded-xl overflow-y-auto"
      v-else-if="bStore.menuState == 1"
    >
      <Lang />
    </v-sheet>
  </v-overlay>

  <div v-else>
    <v-btn icon class="menu-btn" @click="mobileDialog = !mobileDialog">
      <v-icon>
        <EllipsisHorizontalIcon class="menu-icon" />
      </v-icon>
    </v-btn>
    <v-bottom-sheet
      v-model="mobileDialog"
      content-class="elevation-0"
      transition="slide-y-reverse-transition"
      :class="clsx('d-flex justify-center dialog-blur')"
    >
      <v-sheet
        :class="
          clsx(
            'align-self-center rounded-t-xl',
            'menu-card-mobile', 'pb-6'
          )
        "
        v-if="bStore.menuState == 0"
        elevation="3"
      >
        <div class="pt-0 d-flex align-center justify-center">
          <button @click="mobileDialog = false" class="w-100">
            <v-icon size="16" style="opacity: 50%;">
              <ChevronDownIcon />
            </v-icon>
          </button>
        </div>
        <!-- Settings -->
        <Language />

        <!-- Dark mode -->
        <Dark />

        <!-- Recent exchanges -->
        <RecentExchanges />

        <v-divider class="mx-5 my-4" style="color: var(--palette-black-50)" />

        <!-- Links -->
        <Links />
      </v-sheet>
      <v-sheet
        class="menu-card-mobile py-4 rounded-t-xl overflow-y-auto"
        v-else-if="bStore.menuState == 1"
      >
        <Lang />
      </v-sheet>
    </v-bottom-sheet>
  </div>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import { useBridgeStore } from "~/stores/bridge/bridge";
import { VBottomSheet } from "vuetify/labs/VBottomSheet";
import {
  EllipsisHorizontalIcon,
  ChevronDownIcon,
} from "@heroicons/vue/20/solid";

import Lang from "./lang.vue";
import Dark from "~/components/toolbar/menu/dark.vue";
import Links from "~/components/toolbar/menu/links.vue";
import Language from "~/components/toolbar/menu/language.vue";
import RecentExchanges from "~/components/toolbar/menu/recentExchanges.vue";

const bStore = useBridgeStore();
const { mobile } = useDisplay();
let mobileDialog = ref(false);
</script>

<style lang="scss" scoped>
.menu-icon {
  width: 22px;
  height: 22px;
}
.menu-card {
  min-width: 240px;
  min-height: 220px;
  box-shadow: 0 2px 32px #0003 !important;
}
.menu-card-mobile {
  width: 100vw;
}
.setting-item {
  height: 48px;
}
</style>