import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { initAuth } from './utils/auth'


const app = createApp(App)
const pinia = createPinia()

app.use(pinia)

initAuth().then(() => {
  app.use(router)
  app.mount('#app')
})