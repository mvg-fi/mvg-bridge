<template>
  <v-row no-gutters class="fill-height d-flex justify-center w-100">
    <div>
      <v-card
        class="bridge-card rounded-xl mt-10 pa-0 d-flex flex-column"
        elevation="0"
      >
        <Title class="mx-6 mt-4" />

        <div class="my-2">
          <Input class="my-2" :from="true" />
          <Switch class="my-1" />
          <Input class="my-2" :from="false" />
        </div>
        
        <Receiver class="my-2" />
        <Confirm class="my-3" />
      </v-card>
    </div>
  </v-row>
</template>

<script setup>
import Title from "~/components/bridge/elements/title.vue";
import Input from "~/components/bridge/elements/input.vue";
import Switch from "~/components/bridge/elements/switch.vue";
import Receiver from "~/components/bridge/elements/receiver.vue";
import Confirm from "~/components/bridge/elements/confirm.vue";
</script>

<style lang="scss" scoped>
@import "assets/css/custom";
.bridge-card {
  width: 540px;
  box-shadow: 0 1px 8px #0003 !important;
}
</style>