import { createI18n } from 'vue-i18n'
import zh from '~/locales/zh.json'
import en from '~/locales/en.json'
import zh_hk from '~/locales/zh-hk.json'

export const langs = [
  {name: "English", key: "en"},
  {name: "简体中文", key: "zh"},
  {name: "繁體中文", key: "zh_hk"},
]

export const getNameByKey = (k: string) : string => {
  switch(k) {
  case "en":
    return "English"
  case "zh":
    return "简体中文"
  case "zh_hk":
    return "繁體中文"
  default:
    return "English"
  }
} 

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: 'en',
  messages: {
    en,
    zh,
    zh_hk
  }
})

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.vueApp.use(i18n)
})