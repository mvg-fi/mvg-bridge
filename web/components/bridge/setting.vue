<template>
  <v-row no-gutters class="fill-height d-flex justify-center w-100">
    <v-card
      class="bridge-card rounded-xl mt-10 pa-5 d-flex flex-column"
      elevation="0"
    >
      <Title />
      <Custom class="mt-9"/>
      <v-divider class="my-4"/>
      <Dex/>
      <v-divider class="my-4"/>
      <Privacy/>
    </v-card>
  </v-row>
</template>

<script setup>
import Title from "~/components/bridge/settings/title.vue"
import Custom from "~/components/bridge/settings/custom.vue"
import Dex from "~/components/bridge/settings/dex.vue"
import Privacy from "~/components/bridge/settings/privacy.vue"
</script>

<style lang="scss" scoped>
.bridge-card {
  max-height: 560px;
  width: 540px;
  box-shadow: 0 1px 8px #0003 !important;
}
</style>