<template>
  <button
    :key="props.w.name"
    class="d-flex flex-row align-center mb-3"
    style="height: 40px; gap: 8px"
  >
    <v-img
      class="flex-grow-0 rounded-circle"
      :src="props.w.icon"
      style="width: 32px; height: 32px"
    />

    <span
      class="flex-grow-1 justify-start text-start"
      style="font-size: 16px; line-height: 32px; font-family: Satoshi-SemiBold"
    >
      {{ props.w.name }}
    </span>

    <div v-if="props.w.connected">
      <v-icon size="16" color="primary">
        <CheckIcon />
      </v-icon>
    </div>

    <div v-if="props.w.loading">
      <v-progress-circular
        indeterminate
        color="primary"
        size="16"
        width="2"
      ></v-progress-circular>
    </div>
  </button>
</template>

<script lang="ts" setup>
import { CheckIcon } from "@heroicons/vue/24/outline";
const props = defineProps(["w"]);
</script>


<style lang="scss" scoped>

</style>