import './styles/index.css'
import '../node_modules/vditor/src/assets/less/index.less'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { VDialogPlugin } from './plugins/dialog'
import { useTheme } from './hooks/useTheme'

const app = createApp(App)
app.use(router)
app.use(createPinia())

app.use(VDialogPlugin)

fetch('/configs/config.json')
    .then((response) => response.json())
    .then((configs) => {
        app.mount('#app')

        const { theme, themes, switchTheme, config } = useTheme()
        Object.assign(config, configs)

        themes.push(...Object.keys(config.themes))

        if (themes.length > 0 && theme.value == '') {
            theme.value = themes[0]
            switchTheme(theme.value)
        }
    })

    .catch((error) => {
        console.error('Failed to load config:', error)
    })
