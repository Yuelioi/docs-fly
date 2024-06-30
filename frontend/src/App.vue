<script setup lang="ts">
const { theme } = useTheme()
</script>

<template>
    <div :class="theme" class="text-sm bg-theme-base text-theme-base border-theme-base">
        <div class="flex flex-col h-lvh">
            <!--  顶部导航栏  -->
            <div class="border shadow-md border-b-theme-base">
                <TheHeader />
            </div>

            <div class="flex-grow border-2 bg-theme-base border-t-theme-muted">
                <router-view></router-view>
            </div>

            <!--  页脚  -->
            <div class="py-4 bg-theme-footer">
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
