<template>
  <v-btn icon elevation="0">
    <v-icon>
      <Bars3Icon />
    </v-icon>

    <v-menu activator="parent" class="mt-3 rounded-xl">
      <v-list>
        <v-list-item v-for="(item, index) in items" :key="index" :value="index" style="width: 88px;">
          <v-list-item-title>{{ item.name }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-btn>

  <div></div>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n";
import { Bars3Icon } from "@heroicons/vue/24/solid";

const t = useI18n()
const items = computed(()=>[
  {name: t.t('bridge'), link: '/bridge'},
  // {name: '', link: ''},
])
</script>