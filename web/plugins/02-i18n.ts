import { createI18n } from 'vue-i18n'
import zh from '~/locales/zh.json'
import en from '~/locales/en.json'
import zhhk from '~/locales/zh-hk.json'

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: 'en',
  messages: {
    en,
    zh,
    zhhk
  }
})

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.vueApp.use(i18n)
})