<script setup lang="ts">
import HeaderArea from '@/components/HeaderArea.vue'
import BackTop from '@/components/BackTop.vue'
import Footer from '@/components/FooterArea.vue'

import { themeState } from './utils/themeManager'
</script>

<template>
    <div :class="themeState.currentTheme" class="bg-theme-base text-theme-base border-theme-base">
        <div class="flex flex-col h-lvh">
            <!--  顶部导航栏  -->
            <div class="header-area shadow-md"><HeaderArea /></div>
            <!--  顶部导航栏结束  -->

            <!--  主要内容  -->
            <div class="content flex-grow bg-theme-base">
                <router-view></router-view>
            </div>
            <!--  主要内容结束 -->

            <!--  页脚  -->
            <div class="py-8 footer bg-theme-footer">
                <Footer />
            </div>
            <!--  页脚结束  -->
        </div>
    </div>
    <BackTop />
</template>
