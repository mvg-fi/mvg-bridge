<template>
  <v-sheet class="menu-card py-3 rounded-xl">
    <button class="px-4 mb-4 d-flex align-center" @click="bStore.setMenuState(0)">
      <v-icon size="16">
        <ChevronLeftIcon />
      </v-icon>

      <span class="ml-2"> {{ $t('language') }} </span>
    </button>
    <v-list class="px-5 py-0 d-flex flex-column align-center overflow-y-auto  ">
      <button
        :key="lang.key"
        v-for="lang in langs"
        @click="setLang(lang.key)"
        style="color: var(--palette-black-100)"
        class="px-0 pb-4 w-100 text-start d-flex align-center justify-space-between"
      >
        <span class="h7m">
          {{ lang.name }}
        </span>

        <v-icon size="16" v-if="activeLang == lang.key">
          <CheckIcon />
        </v-icon>
      </button>
    </v-list>
  </v-sheet>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n";
import { langs } from "~/plugins/02-i18n";
import { CheckIcon, ChevronLeftIcon } from "@heroicons/vue/24/outline";
import { useBridgeStore } from "~/stores/bridge/bridge";

const t = useI18n();
const bStore = useBridgeStore()
const activeLang = computed(() => {
  return t.locale.value;
});
const setLang = (lang: string) => {
  t.locale.value = lang;
  // bStore.setMenuState(0);
};
</script>