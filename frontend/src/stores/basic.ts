import { defineStore } from 'pinia'
import { ref } from 'vue'

/**
 * 基础存储
 */
export const basicStore = defineStore('basic', () => {
    /**
     * 管理员
     */
    const isAdmin = ref(false)

    /**
     * 验证token
     */
    const token = ref('')

    /**
     * 语言
     */
    const locale = ref('')

    /**
     * 语言字典,后续可以单独保存配置文件
     */
    const i18n = {
        en: 'English',
        zh: '中文'
    }

    /**
     * 翻译函数
     */
    const t = function (source: string) {
        return i18n[source as keyof typeof i18n]
    }

    return { isAdmin, token, locale, t }
})
