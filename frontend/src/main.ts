import './assets/main.css'
import 'leaflet/dist/leaflet.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// Create the Vue application
const app = createApp(App)

// Use plugins
app.use(createPinia())
app.use(router)

// Mount the application
app.mount('#app')
