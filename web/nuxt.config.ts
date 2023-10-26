// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  css: ['vuetify/lib/styles/main.sass', '~/assets/css/main.scss', '~/assets/css/custom.scss'],
  build: {
    transpile: ['vuetify'],
  },
  runtimeConfig: {
    ETHERSCAN_TOKEN: process.env.ETHERSCAN_TOKEN,
    INFURA_PROJECT_ID: process.env.INFURA_PROJECT_ID,
    WALLETCONNECT_ID: process.env.WALLETCONNECT_ID,
  },
  app: {
    head: {
      meta: [
        { charset: "utf-8" },
        {
          name: "viewport",
          content:
            "width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no, viewport-fit=cover",
        },
      ],
    },
  },
  alias: {

  },
  modules: [
    '@pinia/nuxt',
  ]
})