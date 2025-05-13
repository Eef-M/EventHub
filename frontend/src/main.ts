import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { addIcons, OhVueIcon } from 'oh-vue-icons'
import * as AllIcons from "oh-vue-icons/icons";
import { createPinia } from 'pinia'
import { initAuth } from './utils/auth'

const Icons = Object.values({ ...AllIcons });
addIcons(...Icons);

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)

initAuth().then(() => {
  app.use(router)
  app.component("v-icon", OhVueIcon)
  app.mount('#app')
})