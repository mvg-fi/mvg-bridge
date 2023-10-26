<template>
  <v-app>
    <NuxtLoadingIndicator v-if="loading"/>
    <NuxtLayout v-else>
      <NuxtPage />
    </NuxtLayout>
  </v-app>
</template>

<script lang="ts" setup>
import { initWeb3Modal } from "~/helpers/web3modal/web3modal";
import { BridgeFullName, AppFullName } from "~/helpers/constants";

useHead({
  title: BridgeFullName,
  titleTemplate: `%s - ${AppFullName}`,
});

initWeb3Modal()

const loading = ref(true)
onMounted(()=>loading.value=false)
</script>
<style>

</style>