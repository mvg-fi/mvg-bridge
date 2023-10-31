<template>
  <div class="d-flex justify-center">
    <v-btn class="confirm-btn rounded-pill" @click="confirmBridge">
      <span class="h6" style="color: var(--palette-background-1)">
        {{ $t("bridge") }}
      </span>
    </v-btn>
  </div>
</template>

<script setup>
const confirmBridge = () => {
  // 1. Choose wallet to pay or scan
  // 2. Invoke sdk to pay
  // 3. Listen to payment
  // 4. Once paid, redirect to status page
}
</script>

<style lang="scss" scoped>
.confirm-btn {
  width: 188px;
  height: 56px;
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
}
</style>