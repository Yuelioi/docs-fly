<script setup lang="ts">
import { themeState } from './utils/themeManager'
</script>

<template>
    <div :class="themeState.currentTheme" class="bg-theme-base text-theme-base border-theme-base">
        <div class="flex flex-col h-lvh">
            <!--  顶部导航栏  -->
            <div class="border shadow-md border-b-theme-base"><TheHeader /></div>

            <div class="flex-grow border-2 bg-theme-base border-t-theme-muted">
                <router-view></router-view>
            </div>

            <!--  页脚  -->
            <div class="py-8 bg-theme-footer">
                <!-- 版权申明 -->
                <TheFooter />
                <!-- 额外脚本 -->
                <FooterScript />
            </div>
        </div>
    </div>

    <BackTop />
</template>

<style scoped></style>
