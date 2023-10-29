<template>
  <v-app-bar elevation="0" class="px-8">
    <Logo/>
    
    <v-spacer />
    
    <Menu class="mx-4"/>
    <!-- <Network class="mx-4" /> -->
    <ConnectWallet class="mx-4" v-if="!cStore.connected"/>
    <Account class="mx-4" v-else/> 
  </v-app-bar>
</template>

<script setup>
import Menu from '~/components/toolbar/menu.vue';
import Logo from '~/components/toolbar/logo.vue';
import Network from '~/components/toolbar/network.vue';
import Account from '~/components/toolbar/account.vue';
import ConnectWallet from '~/components/toolbar/connect.vue';
import { useConnectStore } from '~/stores/connect/connect';
const cStore = useConnectStore()
</script>