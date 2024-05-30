<script setup lang="ts">
import HeaderArea from '@/components/HeaderArea.vue'
import BackTop from '@/components/BackTop.vue'
import Footer from '@/components/FooterArea.vue'

import { themeState } from './utils/themeManager'
</script>

<template>
    <div :class="themeState.currentTheme">
        <div class="flex flex-col h-lvh bg-theme-base text-theme-base">
            <!--  顶部导航栏  -->
            <div class="header-area shadow-md border-theme-base"><HeaderArea /></div>
            <!--  顶部导航栏结束  -->

            <!--  主要内容  -->
            <div class="content flex-grow bg-theme-base">
                <router-view></router-view>
            </div>
            <!--  主要内容结束 -->

            <!--  页脚  -->
            <div class="py-8 footer bg-theme-base">
                <Footer />
            </div>
            <!--  页脚结束  -->
        </div>
    </div>
    <BackTop />
</template>
