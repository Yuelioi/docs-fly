<script setup lang="ts">
import HeaderArea from '@/components/HeaderArea.vue'
import BackTop from '@/components/BackTop.vue'
import Footer from '@/components/FooterArea.vue'
</script>

<template>
    <div>
        <div class="flex flex-col h-lvh dark:bg-dark-base dark:text-slate-200">
            <!--  顶部导航栏  -->
            <div class="header-area shadow-md border-slate-200"><HeaderArea /></div>
            <!--  顶部导航栏结束  -->

            <!--  主要内容  -->
            <div class="content flex-grow dark:bg-dark-base">
                <router-view></router-view>
            </div>
            <!--  主要内容结束 -->

            <!--  页脚  -->
            <div class="py-8 footer dark:bg-dark-base">
                <Footer />
            </div>
            <!--  页脚结束  -->
        </div>
    </div>
    <BackTop />
</template>
