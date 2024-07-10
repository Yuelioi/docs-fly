import './styles/index.css'
import '../node_modules/vditor/src/assets/less/index.less'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { VDialogPlugin } from './plugins/dialog'

import ToastService from 'primevue/toastservice'

import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'

const app = createApp(App)
app.use(router)
app.use(createPinia())
app.use(ToastService)
app.use(PrimeVue, {
    theme: {
        preset: Aura,
        options: {
            prefix: 'p',
            darkModeSelector: 'dark'
        }
    }
})

app.use(VDialogPlugin)
app.mount('#app')
