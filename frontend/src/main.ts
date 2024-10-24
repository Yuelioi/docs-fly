import './styles/index.css'
// import '../node_modules/vditor/src/assets/less/index.less'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { VDialogPlugin } from './plugins/dialog'

const app = createApp(App)
app.use(router)
app.use(createPinia())

app.use(VDialogPlugin)
app.mount('#app')
