<template>
  <div class="px-2 d-flex bg-secondary rounded-xl">
    <TextField :from="props.from"/>
    <Asset :from="props.from"/>
  </div>
</template>

<script setup>
import Asset from "~/components/bridge/elements/input/asset.vue"
import TextField from "~/components/bridge/elements/input/textfield.vue"
const props = defineProps(['from'])
</script>