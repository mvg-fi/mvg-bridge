import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const Custom = {
  dark: false,
  colors: {
    background: '#F9F9F9',
    primary: '#5959D7',
    secondary: '#557DED',
    error: '#FF3333',
    correct: '#75D658',
    black: '#261F22',
    "black-darken-1": "#D8D8D8",
    "black-darken-2": "#CFCECE",
    "black-darken-3": "#B2B1B1",
    "black-darken-4": "#837F81",
    "black-darken-5": "#433F41",
    // 'blacksf': '#433F41',
    // 'blackfifty': '#837F81',
    // 'blacktf': '#B2B1B1',
    // 'blackten': '#CFCECE',
    // 'blackfive': '#D8D8D8',
  }
} 

export default defineNuxtPlugin(nuxtApp => {
  const vuetify = createVuetify({
    ssr: true,
    components,
    directives,
    theme: {
      defaultTheme: "Custom",
      themes: {
        Custom
      }
    }
  })
  nuxtApp.vueApp.use(vuetify)
})