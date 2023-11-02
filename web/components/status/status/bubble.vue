<template>
  <div class="d-flex flex-column justify-center align-center">
    <div
      class="bg-white d-flex align-center justify-center rounded-pill"
      style="width: 56px; height: 56px"
    >
      <v-icon
        size="24"
        style="color: var(--palette-main-correct)"
        v-if="props.state == 2"
      >
        <CheckIcon />
      </v-icon>
      <v-progress-circular
        size="20"
        width="2"
        color="primary"
        indeterminate
        v-else-if="props.state == 1"
      />
      <v-icon size="56" class="bg-black-darken-2 rounded-pill" v-if="props.state == 0">
      </v-icon>
    </div>
    <span class="text-center mt-3">
      {{ props.text }}
    </span>
  </div>
</template>

<script setup>
import { CheckIcon } from "@heroicons/vue/24/solid";

const props = defineProps(["text", "state"]);
// state 0 == not init, 1 == loading, 2 == correct, 3 == error
</script>