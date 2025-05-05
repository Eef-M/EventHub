import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { addIcons, OhVueIcon } from 'oh-vue-icons'
import * as AllIcons from "oh-vue-icons/icons";

const Icons = Object.values({ ...AllIcons });
addIcons(...Icons);

const app = createApp(App)
app.use(router)
app.component("v-icon", OhVueIcon)
app.mount('#app')
