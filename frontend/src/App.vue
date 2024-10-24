<template>
  <main class="flex flex-col h-lvh w-full min-h-screen text-sm bg-base-200 justify-between">
    <!--  顶部导航栏  -->
    <div class="shadow-md sm:h-16 w-full sticky top-0 z-50">
      <TheHeader />
    </div>

    <!-- 主要内容 -->
    <div class="flex-1">
      <router-view></router-view>
    </div>

    <!--  页脚  -->
    <div class="pt-4 sticky">
      <!-- 版权申明 -->
      <TheFooter />
      <!-- 额外脚本 -->
      <FooterScript />
    </div>
    <!-- 手机端底部间距 -->
    <div class="hidden sm:block"></div>
  </main>

  <BackTop />
</template>
