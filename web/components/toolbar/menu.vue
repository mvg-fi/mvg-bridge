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
      <!-- Language -->
      <div
        class="d-flex justify-space-between align-center px-5 py-0 setting-item"
      >
        <span class="h7m">
          {{ $t("language") }}
        </span>
        <button class="px-0 justify-end align-center" @click="moveToLang()">
          {{ activeLang || "EN" }}
          <v-icon size="16">
            <ChevronRightIcon />
          </v-icon>
        </button>
      </div>

      <!-- Dark mode -->
      <div
        class="d-flex justify-space-between align-center px-5 py-0 setting-item"
      >
        <span class="h7m flex-grow-1">
          {{ $t("dark_mode") }}
        </span>
        <div class="justify-end align-center">
          <v-switch
            inset
            v-model="dark"
            hide-details
            color="blue"
            flat
          ></v-switch>
        </div>
      </div>

      <!-- Recent exchanges -->
      <div
        class="d-flex justify-space-between align-center px-5 py-0 setting-item"
      >
        <span class="h7m flex-grow-1">
          {{ $t("show_recent_exchanges") }}
        </span>
        <div class="justify-end align-center">
          <v-switch
            inset
            v-model="showRecentExchanges"
            hide-details
            color="blue"
            flat
          ></v-switch>
        </div>
      </div>

      <v-divider class="mx-5 my-4" style="color: var(--palette-black-50)" />

      <!-- Links -->
      <div
        class="d-flex flex-column justify-start text-start px-5 py-0 h7m"
        style=""
      >
        <a
          :key="l.name"
          :href="l.link"
          target="_blank"
          v-for="l in Links"
          class="py-1"
          style="text-decoration: none; color: var(--palette-black-75)"
        >
          {{ l.name }}
          <v-icon size="16">
            <ExternalIcon />
          </v-icon>
        </a>
      </div>
    </v-sheet>

    <Lang v-if="bStore.menuState == 1" />
  </v-overlay>

  <div v-else>
    <v-btn icon class="menu-btn" @click="mobileDialog = !mobileDialog">
      <v-icon>
        <EllipsisHorizontalIcon class="menu-icon" />
      </v-icon>
    </v-btn>
    <v-dialog
      fullscreen
      v-model="mobileDialog"
      transition="slide-y-reverse-transition"
      :class="clsx('d-flex justify-center dialog-blur')"
    >
      <v-sheet
        :class="
          clsx(
            'select-asset-card align-self-center rounded-xl pt-3 overflow-y-hidden',
            'mobile-card menu-card-mobile'
          )
        "
        elevation="3"
      >
      </v-sheet>
    </v-dialog>
  </div>
</template>

<script setup>
import {
  EllipsisHorizontalIcon,
  ChevronRightIcon,
} from "@heroicons/vue/20/solid";
import { useI18n } from "vue-i18n";
import { useDisplay } from "vuetify";
import { AppURL, DeveloperURL, DocumentationURL } from "~/helpers/constants";
import { useBridgeStore } from "~/stores/bridge/bridge";

import Lang from "./lang.vue";
import ExternalIcon from "./externalIcon.vue";
import { getNameByKey } from "~/plugins/02-i18n";
import clsx from "clsx";

const t = useI18n();
const Links = computed(() => {
  return [
    { name: t.t("documentation"), link: DocumentationURL },
    { name: t.t("developer"), link: DeveloperURL },
    { name: t.t("about_us"), link: AppURL + "/345" },
  ];
});

const bStore = useBridgeStore();
const { mobile } = useDisplay();
let dark = ref(false);
let mobileDialog = ref(false);
const moveToLang = () => {
  bStore.setMenuState(1);
};
const activeLang = computed(() => {
  return getNameByKey(t.locale.value);
});
const showRecentExchanges = computed({
  get() {
    return bStore.recentCardState != 0;
  },
  set(value) {
    if (value) bStore.setRecentCardState(2);
    else bStore.setRecentCardState(0);
  },
});
</script>

<style lang="scss" scoped>
.menu-icon {
  width: 22px;
  height: 22px;
}
.menu-card {
  // border-radius: 2vh;
  min-width: 240px;
  min-height: 220px;
  box-shadow: 0 2px 32px #0003 !important;
}
.menu-card-mobile {
  padding-top: 80%;
  height: 20%;
}
.setting-item {
  height: 48px;
}
</style>