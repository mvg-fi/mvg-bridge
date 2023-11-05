<template>
  <div class="d-flex align-center justify-center flex-row" style="gap: 24px">
    <v-btn
      size="large"
      class="bg-background rounded-pill"
      elevation="0"
      style="border: 1px solid var(--palette-black-10)"
      @click="viewAddress"
    >
      <span class="h7 text-black-darken-4">
        {{ $t("view_address") }}
      </span>
    </v-btn>

    <v-btn class="confirm-btn rounded-pill" size="large" @click="confirmBridge">
      <span class="h7 text-background">
        {{ $t("pay_with_wallet") }}
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
};
const viewAddress = () => {};
</script>

<style lang="scss" scoped>
.confirm-btn {
  background: linear-gradient(
    90deg,
    var(--palette-main-primary),
    var(--palette-main-secondary)
  );
  text-transform: capitalize;
}
</style>