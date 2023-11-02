<template>
  <div class="h-100 w-100">
    <Title class="pt-12 mb-2" />
    <Route class="my-6" />
    <State class="my-8"/>
    <Details class="my-8" />
    <Buttons class="my-8"/>
  </div>
</template>

<script setup>
import Title from "~/components/status/status/title";
import State from "~/components/status/status/state";
import Details from "~/components/status/status/details";
import Buttons from "~/components/status/status/buttons";
import Route from "~/components/bridge/elements/dialogs/bridgeProcess/route";
</script>

<style>
</style>