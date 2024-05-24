<template>
    <div class="w-full">
        <div class="flex flex-col">
            <!-- 主页手风琴 -->
            <div class="w-full h-[50vh] relative">
                <div class="flex justify-center items-center h-full select-none">
                    <span
                        class="animated-text absolute z-20 top-1/3 text-red-50 md:text-2xl text-xl lg:text-3xl"
                        >{{ yiyan }}</span
                    >
                </div>

                <img
                    class="absolute top-0 w-full h-full object-cover"
                    src="https://cdn.yuelili.com/docs/web/assert/banner-anime-girl-roller.jpg"
                    alt="" />
            </div>
            <!-- 主页手风琴结束 -->

            <!--  -->
            <div class="md:mt-[-14rem]">
                <div class="mx-auto max-w-7xl py-12 sm:px-8 sm:py-24 lg:px-8">
                    <div
                        class="relative isolate overflow-hidden bg-slate-100 dark:bg-dark-base px-8 pt-16 shadow-2xl sm:rounded-3xl sm:px-16 md:pt-24 lg:flex lg:gap-x-20 lg:px-24 lg:pt-0">
                        <svg
                            viewBox="0 0 1024 1024"
                            class="absolute left-1/2 top-1/2 -z-10 h-[64rem] w-[64rem] -translate-y-1/2 [mask-image:radial-gradient(closest-side,white,transparent)] sm:left-full sm:-ml-80 lg:left-1/2 lg:ml-0 lg:-translate-x-1/2 lg:translate-y-0"
                            aria-hidden="true">
                            <circle
                                cx="512"
                                cy="512"
                                r="512"
                                fill="url(#759c1415-0410-454c-8f7c-9a820de03641)"
                                fill-opacity="1" />
                            <defs>
                                <radialGradient id="759c1415-0410-454c-8f7c-9a820de03641">
                                    <stop stop-color="#7775D6" />
                                    <stop offset="1" stop-color="#E935C1" />
                                </radialGradient>
                            </defs>
                        </svg>
                        <div
                            class="mx-auto text-center lg:mx-0 lg:flex-auto py-12 lg:py-24 lg:text-left">
                            <h2 class="text-3xl font-bold tracking-tight sm:text-4xl">
                                欢迎来到月离文档站.<br />Welcome to Yueli Docs
                            </h2>
                            <p class="mt-6 text-lg leading-8 text-black-300">
                                国内最大CG文档网站(不是)<br />
                                Largest CG Documentation Site(X).
                            </p>
                            <div
                                class="mt-10 flex items-center justify-center gap-x-6 lg:justify-start">
                                <a href="#" class="btn-primary px-3.5 py-2.5">开始阅读</a>
                                <a href="#" class="text-sm font-semibold leading-6"
                                    >Start Reading <span aria-hidden="true">→</span></a
                                >
                            </div>
                        </div>
                        <div class="mt-8 relative lg:w-1/2 h-80 lg:-mt-8">
                            <div class="h-80 w-full lg:-mt-8 absolute">
                                <img
                                    class="absolute left-0 top-8 w-full lg:w-[60rem] max-w-none rounded-md bg-white-base/5 ring-1 ring-white/10"
                                    src="https://cdn.yuelili.com/docs/web/assert/anime-girl.jpg"
                                    alt="App screenshot"
                                    width="1824"
                                    height="1080" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <!--  -->

            <!-- 主页信息统计区域 -->
            <div class="bg-white-base dark:bg-dark-base">
                <div class="mx-auto max-w-7xl py-18 sm:px-6 sm:py-32 lg:px-8">
                    <div
                        class="max-h-96 relative rounded-b-none rounded-3xl border-radius isolate overflow-hidden bg-gray-900 py-24 sm:py-32">
                        <img
                            src="https://cdn.yuelili.com/docs/web/assert/anime-girl-dream.jpg"
                            alt=""
                            class="relative -translate-y-1/4 inset-0 -z-10 h-full w-full object-cover object-right md:object-center" />

                        <div class="mx-auto max-w-7xl px-6 lg:px-8">
                            <div class="mx-auto max-w-2xl lg:mx-0">
                                <h2
                                    class="text-3xl font-bold tracking-tight text-white sm:text-4xl">
                                    站点统计
                                </h2>
                            </div>
                        </div>
                    </div>
                    <div class="flex text-center shadow-2xl sm:rounded-3xl">
                        <div class="rounded-es-lg border-2 border-r-0 basis-1/4 py-6">
                            <div class="text-sm">书籍数量</div>
                            <div class="text-lg font-bold">{{ statistic?.book_count }}</div>
                        </div>
                        <div class="border-2 border-r-0 basis-1/4 p-6">
                            <div class="text-sm">文章数量</div>
                            <div class="text-lg font-bold">{{ statistic?.document_count }}</div>
                        </div>
                        <div class="border-2 border-r-0 basis-1/4 p-6">
                            <div class="text-sm">历史访问人数</div>
                            <div class="text-lg font-bold">
                                {{ statistic?.historical_visitor_count }}
                            </div>
                        </div>
                        <div class="rounded-lg border-2 rounded-tr-none basis-1/4 p-6">
                            <div class="text-sm">今日访问人数</div>
                            <div class="text-lg font-bold">
                                {{ statistic?.today_visitor_count }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <!-- 主页信息统计区域结束 -->
            <!-- 时间 -->
            <Clock />

            <!-- 时间结束 -->
        </div>
    </div>
</template>

<script setup lang="ts">
import type { HomeStatistic } from '@/models'
import { onMounted, ref } from 'vue'
import Clock from '@/components/VClock.vue'

import { fetchStatisticHome, fetchYiYan } from '@/handlers/index'

const yiyan = ref('')

const statistic = ref<HomeStatistic>()

onMounted(async () => {
    const [ok, data] = await fetchStatisticHome()
    if (ok) {
        statistic.value = data
    }

    const [ok2, data2] = await fetchYiYan()

    if (ok2) {
        yiyan.value = data2['hitokoto']
    }
})
</script>

<style scoped lang="css">
:deep(.el-statistic__head),
:deep(.el-statistic__content) {
    text-align: center;
}
.animated-text {
    animation: fadeInOut 3s ease-in-out;
}

@keyframes fadeInOut {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
}
</style>
