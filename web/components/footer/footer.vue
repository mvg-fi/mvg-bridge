<template>
  <v-row
    no-gutters
    class="d-flex flex-row px-16 pt-12"
    style="padding-bottom: 160px"
    v-if="!customMobile"
  >
    <v-col class="d-flex flex-column">
      <div class="d-flex justify-start align-center">
        <v-icon class="mr-4" size="48">
          <v-img :src="Icon"></v-img>
        </v-icon>
        <span class="logo-text h6" style="font-size: 18px; min-width: 134px">
          {{ AppFullName }}
        </span>
      </div>

      <div class="mt-8 ml-2 d-flex align-center">
        <v-menu>
          <template v-slot:activator="{ props }">
            <button v-bind="props">
              <v-icon size="18">
                <MessageIcon />
              </v-icon>
              <span class="ml-2 mt-1"> {{ activeLang }} </span>
            </button>
          </template>

          <LangMenu />
        </v-menu>
      </div>
    </v-col>

    <v-col
      class="d-flex flex-column align-start mt-3 px-3"
      v-for="(l, i) in Object.keys(links)"
      :key="l"
    >
      <span class="h6m justify-center">
        {{ $t(l) }}
      </span>
      <div class="mt-4 flex-column">
        <div v-for="p in Object.values(links)[i]" :key="p" class="mb-2">
          <a :href="p.link" target="_blank">
            {{ p.name }}
          </a>
        </div>
      </div>
    </v-col>
  </v-row>

  <v-row
    no-gutters
    class="d-flex flex-row justify-center align-center px-16 py-12"
    v-else
  >
    <v-col class="d-flex flex-column">
      <div class="d-flex justify-center align-center">
        <v-icon class="mr-2" size="24">
          <v-img :src="Icon"></v-img>
        </v-icon>
        <span class="logo-text h7" style="font-size: 18px">
          {{ AppFullName }}
        </span>
      </div>

      <div class="mt-10 ml-2 d-flex flex-column justify-center align-center">
        <span class="h7 justify-center">
          {{ $t("resources") }}
        </span>
        <div class="mt-4 flex-column">
          <div v-for="p in mobileLinks" :key="p" class="h7m mb-2 text-center">
            <a :href="p.link" target="_blank">
              {{ p.name }}
            </a>
          </div>
        </div>
      </div>

      <div class="mt-8 ml-2 d-flex justify-center align-center">
        <v-menu>
          <template v-slot:activator="{ props }">
            <button v-bind="props">
              <v-icon size="12">
                <MessageIcon />
              </v-icon>
              <span class="ml-2 mt-1 h7m" style="font-size: 12px">
                {{ activeLang }}
              </span>
            </button>
          </template>
          <LangMenu />
        </v-menu>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useI18n } from "vue-i18n";
import { useDisplay } from "vuetify";
import Icon from "assets/images/mvg.png";
import {
  APIDocURL,
  BridgeURL,
  MediaKitURL,
  AppFullName,
  DocumentationURL,
  SupportCenterURL,
} from "~/helpers/constants";
import { getNameByKey } from "~/plugins/02-i18n";
import LangMenu from "~/components/footer/langMenu.vue";
import MessageIcon from "~/components/footer/messageIcon.vue";
const t = useI18n();

const activeLang = computed(() => {
  return getNameByKey(t.locale.value);
});

const links = computed(() => {
  return {
    products: [{ name: t.t("bridge_noun"), link: BridgeURL }],
    developers: [
      { name: t.t("documentation"), link: DocumentationURL },
      { name: t.t("api_reference"), link: APIDocURL },
    ],
    resources: [
      { name: t.t("media_kit"), link: MediaKitURL },
      { name: t.t("support_center"), link: SupportCenterURL },
    ],
  };
});

const mobileLinks = computed(() => {
  return [
    { name: t.t("documentation"), link: DocumentationURL },
    { name: t.t("api_reference"), link: APIDocURL },
    { name: t.t("support_center"), link: SupportCenterURL },
    { name: t.t("media_kit"), link: MediaKitURL },
  ];
});

const { width } = useDisplay();
const customMobile = computed(() => width.value <= 800);
</script>

<style>
</style>