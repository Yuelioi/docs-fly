import './styles/index.css'
import '../node_modules/vditor/src/assets/less/index.less'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(router)
app.use(createPinia())

import { themeState, setTheme } from './utils/themeManager'

app.provide('themeState', themeState)
app.provide('setTheme', setTheme)

fetch('/configs/themeConfig.json')
    .then((response) => response.json())
    .then((config) => {
        themeState.availableThemes = config.themes
        return import(`./themes/${themeState.currentTheme}/main.css`)
    })
    .then(() => {
        return import(`./themes/${themeState.currentTheme}/dark.css`)
    })
    .then(() => {
        app.mount('#app')
    })
