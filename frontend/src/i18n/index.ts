import { createI18n } from 'vue-i18n'
import en from './locales/en.json'
import sv from './locales/sv.json'

const i18n = createI18n({
  legacy: false, // Set to false to use Composition API
  locale: 'en', // set default locale
  fallbackLocale: 'en', // set fallback locale
  messages: {
    en,
    sv
  }
})

export default i18n 