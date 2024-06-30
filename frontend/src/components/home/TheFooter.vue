<template>
    <div class=""><div class="py-2 text-center">Copyright © 2024 月离万事屋</div></div>
</template>

<script setup lang="ts">
const basic = basicStore()
let { isAdmin, locale, nickname } = storeToRefs(basic)

const appVersion = ref('')

async function refreshDB() {
    try {
        localStorage.setItem('appVersion', appVersion.value)
        await dbManager.clearDatabase()
        await Message({ message: 'Database refresh successfully' })
        console.log('Database refresh successfully')
    } catch (error) {
        Message({ message: 'Failed to clear database', type: 'warn' })
        console.error('Failed to clear database:', error)
    }
}

onMounted(async () => {
    // 初始化token
    const localToken = localStorage.getItem('token')
    if (localToken) {
        await fetchHandleBasicCallback(
            isAdmin,
            false,
            fetchCheckToken,
            localToken,
            'data',
            async () => {
                isAdmin.value = true
            },
            async () => {
                isAdmin.value = false
                localStorage.removeItem('token')
            }
        )
    } else {
        localStorage.removeItem('token')
    }

    // 读取本地语言/昵称
    // initLocalStorageParam(locale, 'zh', 'locale')
    // initLocalStorageParam(nickname, 'zh', '看书大王')

    // 读取本地数据库
    const appVersionLocal = localStorage.getItem('appVersion')
    if (appVersionLocal) {
        await fetchHandleBasic(appVersion, '', getAppVersion)
        if (appVersionLocal != appVersion.value) {
            await refreshDB()
        }
    } else {
        await fetchHandleBasic(appVersion, '', getAppVersion)
        await refreshDB()
    }
})
</script>
