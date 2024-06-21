<script setup lang="ts">
import HeaderArea from '@/components/home/HeaderArea.vue'
import BackTop from '@/components/home/BackTop.vue'
import Footer from '@/components/home/FooterArea.vue'

import FooterScript from '@/components/common/FooterScript.vue'

import { themeState } from './utils/themeManager'
</script>

<template>
    <div :class="themeState.currentTheme" class="bg-theme-base text-theme-base border-theme-base">
        <div class="flex flex-col h-lvh">
            <!--  顶部导航栏  -->
            <div class="border shadow-md border-b-theme-base"><HeaderArea /></div>
            <!--  顶部导航栏结束  -->

            <!--  主要内容  -->
            <div class="flex-grow border-2 bg-theme-base border-t-theme-muted">
                <router-view></router-view>
            </div>
            <!--  主要内容结束 -->

            <!--  页脚  -->
            <div class="py-8 bg-theme-footer">
                <Footer />
            </div>
            <!--  页脚结束  -->
        </div>
    </div>
    <FooterScript></FooterScript>
    <BackTop />
</template>

<style scoped></style>
