import { ref, reactive } from 'vue'

// 直接用全局变量, 因为就只有一套主题配置

const theme = ref('')
const config = reactive<any>({})
export const themes = reactive<string[]>([])

function loadCss(href: string, _type: string) {
    // 检查是否已经存在相同的样式表
    const existingLink = document.querySelector(`link[href="${href}"]`)
    if (existingLink) {
        return // 样式表已存在，直接返回
    }

    // 移除所有 id 以 theme- 开头的其他 CSS 样式表
    const existingThemeLinks = document.querySelectorAll('link[id^="theme-"]')
    existingThemeLinks.forEach((link) => {
        if (link && !link.id.startsWith(`theme-${theme.value}`)) {
            link.parentNode?.removeChild(link)
        }
    })

    // 创建并插入新的样式表
    const link = document.createElement('link')
    link.rel = 'stylesheet'
    link.type = 'text/css'
    link.href = href
    link.id = `theme-${theme.value}-${_type}`
    document.head.appendChild(link)
}

function loadTheme(_theme: string, cssFiles: string[]) {
    cssFiles.forEach((file) => {
        loadCss(`/themes/${_theme}/${file}`, `${file.replace('.css', '')}`)
    })
}

// 设置主题的函数
export function switchTheme(_theme: string) {
    if (config.themes && config.themes[_theme]) {
        theme.value = _theme
        loadTheme(theme.value, config.themes[_theme])
    } else {
        console.warn(`Theme ${_theme} not found in global config`)
    }
}

// useTheme hook
export const useTheme = function () {
    return { theme, themes, config, switchTheme }
}
