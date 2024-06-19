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

    // 初始化nickname
    const nicknameLocal = localStorage.getItem('nickname')
    if (nicknameLocal) {
        nickname.value = nicknameLocal
    } else {
        nickname.value = ''
    }

    // 初始化数据库

    const appVersionLocal = localStorage.getItem('appVersion')
    if (appVersionLocal) {
        await fetchBasic(appVersion, '', getAppVersion)
        if (appVersionLocal != appVersion.value) {
            await refreshDB()
        }
    } else {
        await fetchBasic(appVersion, '', getAppVersion)
        await refreshDB()
    }
})
</script>
