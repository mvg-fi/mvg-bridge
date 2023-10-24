import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const Custom = {
  dark: false,
  variations: {
    colors: ['black'],
    
  },
  colors: {
    background: '#F9F9F9',
    primary: '#5959D7',
    secondary: '#557DED',
    error: '#FF3333',
    correct: '#75D658',
    black: '#261F22',
    'blacksf': '#433F41',
    'blackfifty': '#837F81',
    'blacktf': '#B2B1B1',
    'blackten': '#CFCECE',
    'blackfive': '#D8D8D8',
  }
} 

export default defineNuxtPlugin(nuxtApp => {
  const vuetify = createVuetify({
    ssr: true,
    components,
    directives,
    theme: {
      themes: {
        Custom
      }
    }
  })
  nuxtApp.vueApp.use(vuetify)
})