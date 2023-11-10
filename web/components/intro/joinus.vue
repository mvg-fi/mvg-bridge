<template>
  <v-row class="d-flex flex-row pa-16 w-100" no-gutters>
    <v-col
      cols="12"
      class="d-flex flex-column align-center text-center justify-center"
    >
      <span :class="clsx('mb-8', customMobile ? 'h5' : 'h4')">
        {{ $t("join_us_now") }}
      </span>

      <div
        :class="
          clsx(
            'd-flex flex-row justify-center align-center',
            customMobile ? 'mt-1' : 'mt-2'
          )
        "
      >
        <div
          :class="customMobile ? 'mx-2' : 'mx-3'"
          v-for="i in socialMedias"
          :key="i"
        >
          <a :href="i.link" target="_blank">
            <v-icon :size="customMobile ? 36 : 48">
              <v-img :src="i.icon" :alt="i.name" />
            </v-icon>
          </a>
        </div>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import clsx from "clsx";
import { useDisplay } from "vuetify";
import di from "~/assets/images/intro/social/discord.png";
import gi from "~/assets/images/intro/social/github.png";
import mi from "~/assets/images/intro/social/mixin.png";
import te from "~/assets/images/intro/social/telegram.png";
import x from "~/assets/images/intro/social/x.png";
import {
  MVGxLink,
  MVGMixinLink,
  MVGGithubLink,
  MVGDiscordLink,
  MVGTelegramLink,
} from "~/helpers/constants";

const socialMedias = [
  { name: "Discord", icon: di, link: MVGDiscordLink },
  { name: "Telegram", icon: te, link: MVGTelegramLink },
  { name: "Mixin", icon: mi, link: MVGMixinLink },
  { name: "X", icon: x, link: MVGxLink },
  { name: "Github", icon: gi, link: MVGGithubLink },
];

const { width } = useDisplay();
const customMobile = computed(() => width.value <= 800);
</script>

<style scoped>
.break-text-special {
  margin-left: calc(600px - 488px);
  width: 600px;
  word-break: break-word;
}
</style>