<template>
  <footer class="footer bg-base-100 text-neutral-content py-8 hidden sm:flex">
    <div class="ml-auto"></div>
    <nav class="basis-6/12 md:basis-4/12 lg:basis-3/12">
      <h6 class="footer-title">其他网站</h6>
      <div class="flex flex-col space-y-3">
        <a class="link link-hover" href="https://www.yuelili.com/" target="_blank">主站</a>
        <a class="link link-hover" href="https://www.yuelili.com/" target="_blank">导航站</a>
      </div>
    </nav>

    <nav class="basis-6/12 md:basis-4/12 lg:basis-3/12">
      <h6 class="footer-title">联系</h6>
      <div class="flex flex-col space-y-3">
        <div class="flex space-x-2">
          <span class="icon-[fa6-brands--github] size-5"></span>
          <a class="link link-hover ml-1" href="https://github.com/Yuelioi" target="_blank">
            Github
          </a>
        </div>

        <div class="flex space-x-2">
          <span class="icon-[fa6-brands--bilibili] size-5"></span>
          <a class="link link-hover ml-1" href="https://space.bilibili.com/4279370" target="_blank">
            Bilibili</a
          >
        </div>
        <div class="flex space-x-2">
          <span class="icon-[mdi--qqchat] size-5"></span>
          <a class="link link-hover ml-1" href="https://qm.qq.com/q/SslCMd4XO6" target="_blank">
            QQ</a
          >
        </div>
      </div>
    </nav>
  </footer>

  <div class="bg-base-200">
    <div class="py-2 text-center">Copyright © 2024 月离万事屋 ・ 苏ICP备2024118526号-1</div>
  </div>
</template>

<script setup lang="ts">
const basic = basicStore()
let { isAdmin, locale, nickname } = storeToRefs(basic)

const appVersion = ref('')

async function refreshDB() {
  try {
    localStorage.setItem('appVersion', appVersion.value)
    await dbManager.clearDatabase()
    Message({ message: 'Database refresh successfully' })
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
