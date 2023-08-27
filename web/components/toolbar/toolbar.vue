<template>
  <v-app-bar elevation="0" class="px-8">
    <Logo/>
    
    <v-spacer />
    
    <Menu class="mx-4"/>
    <Network class="mx-4" />
    <ConnectWallet class="mx-4"/>
  </v-app-bar>
</template>

<script setup>
import Menu from '~/components/toolbar/menu.vue';
import Logo from '~/components/toolbar/logo.vue';
import ConnectWallet from '~/components/toolbar/connect.vue';
import Network from '~/components/toolbar/network.vue';
</script>