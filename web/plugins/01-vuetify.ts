import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const Custom = {
  dark: false,
  colors: {
    background: '#FFFFFF',
    primary: '#5959D7',
    secondary: '#557DED',
    third: "#F9F9F9",
    error: '#FF3333',
    correct: '#75D658',
    black: '#261F22',
    "black-darken-1": "#D8D8D8",  //5%
    "black-darken-2": "#CFCECE",  //10%
    "black-darken-3": "#B2B1B1",  //25%
    "black-darken-4": "#837F81",  //50%
    "black-darken-5": "#433F41",  //75%
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
    },
    display: {
      mobileBreakpoint: 'md',
      thresholds: {
        xs: 0,
        sm: 340,
        md: 540,
        lg: 800,
        xl: 1280,
      },
    },
  })
  nuxtApp.vueApp.use(vuetify)
})