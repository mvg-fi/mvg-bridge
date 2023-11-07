<template>
  <v-app-bar
    elevation="0"
    :class="clsx(!mobile && 'px-8', mobile && 'px-6', 'bg-background')"
  >
    <Logo />

    <v-spacer />

    <Menu :class="clsx(!mobile && 'mx-4')" />
    <!-- <Network class="mx-4" /> -->
    <ConnectWallet class="mx-4" v-if="!cStore.connected" />
    <Account class="" v-else />
  </v-app-bar>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import Menu from "~/components/toolbar/menu.vue";
import Logo from "~/components/toolbar/logo.vue";
import Network from "~/components/toolbar/network.vue";
import Account from "~/components/toolbar/account.vue";
import ConnectWallet from "~/components/toolbar/connect.vue";
import { useConnectStore } from "~/stores/connect/connect";

const { mobile } = useDisplay();
const cStore = useConnectStore();
</script>