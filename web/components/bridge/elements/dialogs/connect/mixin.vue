<template>
  <div class="d-flex flex-column">
    <div class="h7m mb-2">
      <span style="color: var(--palette-black-50)">{{ $t("mixin") }}</span>
    </div>

    <button
      v-for="w in mixinWallets"
      :key="w.name"
      class="d-flex flex-row align-center mb-3"
      style="height: 40px; gap: 8px"
      @click="connect(w.name)"
    >
      <v-img
        class="flex-grow-0 rounded-circle"
        :src="w.icon"
        style="width: 32px; height: 32px"
      />

      <span
        class="flex-grow-1 justify-start text-start"
        style="font-size: 16px; line-height: 32px; font-family: Satoshi-Bold"
      >
        {{ w.name }}
      </span>

      <div v-if="false">
        <!-- If connected -->
      </div>

      <div v-if="false">
        <!-- If loading -->
      </div>
    </button>
  </div>
</template>

<script setup>
import mm from "~/assets/images/wallets/mixin.png";
import fe from "~/assets/images/wallets/fennec.png";

const mixinWallets = [
  { name: "MixinMessenger", icon: mm },
  { name: "Fennec", icon: fe },
];

const connect = (w) => {
  switch (w) {
    case "MixinMessenger":
      mixinMessenger()
      break;
    case "Fennec":
      fennec()
      break;
  }
}

const mixinMessenger = () => {

}

const fennec = () => {

}
</script>