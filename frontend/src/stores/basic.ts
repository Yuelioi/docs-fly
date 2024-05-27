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
        locale: {
            en: 'English',
            zh: '中文'
        },
        displayName: {
            en: 'Display Name',
            zh: '显示名称'
        },
        order: {
            en: 'Order',
            zh: '序号'
        },
        hidden: {
            en: 'Hidden',
            zh: '隐藏'
        }
    }

    /**
     * 翻译函数
     */
    const translate = function (source: string) {
        console.log(source)

        return i18n[source as keyof typeof i18n][locale.value as 'en' | 'zh']
    }

    return { isAdmin, token, locale, translate }
})
