import {  reactive } from 'vue'

const themes = reactive<string[]>([])

function switchTheme(_theme: string) {
    if (themes.includes(_theme)) {
        document.documentElement.setAttribute('data-theme', _theme)
    } else {
        console.warn(`Theme ${_theme} not found in global config`)
    }
}

export function useTheme(_themes: string[]) {
    themes.splice(0, themes.length, ..._themes)
    return { themes, switchTheme }
}
