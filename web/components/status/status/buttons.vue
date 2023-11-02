<template>
  <div class="d-flex justify-center align-center" style="gap: 128px">
    <v-btn
      class="rounded-pill status-button h7"
      elevation="0"
      @click="router.push({ path: '/bridge' })"
    >
      {{ $t("back") }}
    </v-btn>
    <v-btn class="rounded-pill status-button h7" elevation="0">
      {{ $t("refresh") }}
    </v-btn>
  </div>
</template>

<script setup>
const router = useRouter();
</script>

<style scoped>
.status-button {
  width: 128px;
  height: 56px;
  background-color: var(--palette-black-10);
}
</style>