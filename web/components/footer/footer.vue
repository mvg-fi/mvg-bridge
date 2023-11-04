<template>
  <v-row
    no-gutters
    class="d-flex flex-row px-16 pt-12"
    style="padding-bottom: 160px"
  >
    <v-col class="d-flex flex-column">
      <div class="d-flex justify-start align-center">
        <v-icon class="mr-4" size="48">
          <v-img :src="Icon"></v-img>
        </v-icon>
        <span class="logo-text h6"> {{ AppFullName }} </span>
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
      class="d-flex flex-column align-start mt-3"
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
</template>

<script setup>
import Icon from "assets/images/mvg.png";
import { useI18n } from "vue-i18n";
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

const products = reactive();
const developers = reactive();
const resources = reactive();

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
</script>

<style>
</style>