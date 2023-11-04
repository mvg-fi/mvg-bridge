<template>
  <v-sheet class="menu-card py-3 my-4 rounded-lg">
    <v-list class="px-5 py-0 d-flex flex-column align-center overflow-y-auto">
      <button
        :key="lang.key"
        v-for="lang in langs"
        @click="setLang(lang.key)"
        style="width: 80px;"
        class="mx-1 py-2 text-start d-flex align-center text-black justify-space-between"
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

const t = useI18n();
const activeLang = computed(() => {
  return t.locale.value;
});
const setLang = (lang: string) => {
  t.locale.value = lang;
};
</script>