import { reactive } from 'vue'

export const themeState = reactive({
    currentTheme: 'default', // 初始主题
    availableThemes: [] as string[] // 可用主题列表
})

export function setTheme(theme: string) {
    themeState.currentTheme = theme
}
