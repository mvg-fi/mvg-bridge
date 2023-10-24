<template>
  <v-app>
    <NuxtLoadingIndicator v-if="loading"/>
    <NuxtLayout v-else>
      <NuxtPage />
    </NuxtLayout>
  </v-app>
</template>

<script lang="ts" setup>
import { BridgeFullName, AppFullName } from "~/helpers/constants";

useHead({
  title: BridgeFullName,
  titleTemplate: `%s - ${AppFullName}`,
});

const loading = ref(true)
onMounted(()=>loading.value=false)
</script>
<style>

</style>