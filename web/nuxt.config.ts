// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  css: ['vuetify/lib/styles/main.sass', '~/assets/css/main.scss', '~/assets/css/custom.scss'],
  build: {
    transpile: ['vuetify'],
  },  
  alias: {

  },
  modules: [
    '@pinia/nuxt',
  ]
})