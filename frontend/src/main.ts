import './assets/style/index.css'
import '../node_modules/vditor/src/assets/less/index.less'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import { IconPlugin } from './plugins/icon'

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(router)
app.use(createPinia())
app.use(IconPlugin)

app.mount('#app')
