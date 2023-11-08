<template>
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
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import ExternalIcon from "~/components/toolbar/externalIcon.vue";
import { AppURL, DeveloperURL, DocumentationURL } from "~/helpers/constants";

const t = useI18n();
const Links = computed(() => {
  return [
    { name: t.t("documentation"), link: DocumentationURL },
    { name: t.t("developer"), link: DeveloperURL },
    { name: t.t("about_us"), link: AppURL + "/345" },
  ];
});
</script>