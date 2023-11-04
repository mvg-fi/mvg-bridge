import { createI18n } from 'vue-i18n'
import zh from '~/locales/zh.json'
import en from '~/locales/en.json'
import fr from '~/locales/fr.json'
import es from '~/locales/es.json'
import de from '~/locales/de.json'
import ru from '~/locales/ru.json'
import ar from '~/locales/ar.json'
import zh_hk from '~/locales/zh-hk.json'

export const langs = [
  { name: "English", key: "en" },
  { name: "简体中文", key: "zh" },
  { name: "繁體中文", key: "zh_hk" },
  { name: "français", key: "fr" },
  { name: "Deutsch", key: "de" },
  { name: "Español", key: "es" },
  { name: "ру́сский язы́к", key: "ru" },
  { name: "اَلْعَرَبِيَّةُ", key: "ar" },  
]

export const getNameByKey = (k: string): string => {
  switch (k) {
    case "en":
      return "English"
    case "zh":
      return "简体中文"
    case "zh_hk":
      return "繁體中文"
    case "fr":
      return "français"
    case "de":
      return "Deutsch"
    case "es":
      return "Español"
    case "ru":
      return "ру́сский язы́к"
    case "ar":
      return "اَلْعَرَبِيَّةُ"
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
    zh_hk,
    fr,
    es,
    de,
    ru,
    ar,
  }
})

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.vueApp.use(i18n)
})