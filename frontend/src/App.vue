<template>
    <div class="text-sm">
        <div class="flex flex-col h-lvh">
            <!--  顶部导航栏  -->
            <div class="border shadow-md">
                <TheHeader />
            </div>

            <div class="flex-grow">
                <router-view></router-view>
            </div>

            <!--  页脚  -->
            <div class="pt-4">
                <!-- 版权申明 -->
                <TheFooter />
                <!-- 额外脚本 -->
                <FooterScript />
            </div>
            <!-- 手机端底部间距 -->
            <div class="pb-12 md:pb-0"></div>
        </div>
    </div>

    <BackTop />
</template>
