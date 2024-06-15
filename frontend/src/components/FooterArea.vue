<template>
    <div class=""><div class="text-center py-2">Copyright © 2024 月离万事屋</div></div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { fetchCheckToken } from '@/handlers/index'

import { basicStore } from '@/stores/index'
const basic = basicStore()
let { isAdmin, locale } = storeToRefs(basic)

onMounted(async () => {
    // 初始化token
    const localToken = localStorage.getItem('token')
    if (localToken) {
        const [ok, result] = await fetchCheckToken(localToken)
        if (ok) {
            isAdmin.value = true
        } else {
            isAdmin.value = false
            localStorage.removeItem('token')
        }
    } else {
        localStorage.removeItem('token')
    }

    // 初始化语言
    const localeLocal = localStorage.getItem('locale')
    if (localeLocal) {
        locale.value = localeLocal
    } else {
        locale.value = 'zh'
    }

    // 初始化数据库
})
</script>
