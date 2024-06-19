<template>
    <div class="w-full">
        <button @click="addNotify">添加事件</button>
        <div class="flex flex-col -mt-1">
            <!-- 主页手风琴 -->
            <div class="w-full h-[50vh] relative">
                <div class="flex items-center justify-center h-full select-none">
                    <span
                        class="absolute z-20 text-xl animated-text text-slate-200 top-1/3 md:text-2xl lg:text-3xl"
                        >{{ yiyan }}</span
                    >
                </div>

                <img
                    class="absolute top-0 object-cover w-full h-full"
                    src="https://cdn.yuelili.com/docs/web/assert/banner-anime-girl-roller.jpg"
                    alt="" />
            </div>
            <!-- 主页手风琴结束 -->

            <!--  -->
            <div class="md:mt-[-14rem]">
                <div class="py-12 mx-auto max-w-7xl sm:px-8 sm:py-24 lg:px-8">
                    <div
                        class="relative px-8 pt-16 overflow-hidden shadow-2xl isolate bg-theme-base sm:rounded md:pt-24 lg:flex lg:gap-x-20 lg:px-24 sm:px-0 lg:pt-0">
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
                            class="py-12 mx-auto text-center lg:mx-0 lg:flex-auto lg:py-24 lg:text-left">
                            <h2 class="text-3xl font-bold tracking-tight sm:text-4xl">
                                欢迎来到月离文档站.<br />Welcome to Yueli Docs
                            </h2>
                            <p class="mt-6 text-lg leading-8 text-black-300">
                                国内最大CG文档网站(不是)<br />
                                Largest CG Documentation Site(X).
                            </p>
                            <div
                                class="flex items-center justify-center mt-10 gap-x-6 lg:justify-start">
                                <router-link
                                    class="btn bg-theme-primary-base hover:bg-theme-primary-hover px-3.5 py-2.5"
                                    :to="{
                                        name: 'post',
                                        params: {
                                            postPath: rndPostUrl
                                        }
                                    }"
                                    :key="rndPostUrl.join('/')">
                                    <span>开始阅读</span>
                                </router-link>
                                <a href="#" class="text-sm font-semibold leading-6"
                                    >Start Reading <span aria-hidden="true">→</span></a
                                >
                            </div>
                        </div>
                        <div class="relative mt-8 lg:w-1/2 h-80 lg:-mt-8">
                            <div class="absolute w-full h-80 lg:-mt-8">
                                <img
                                    class="absolute left-0 top-8 w-full lg:w-[60rem] max-w-none rounded-md ring-1 ring-white/10"
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
            <div class="">
                <div class="mx-auto max-w-7xl py-18 sm:px-6 sm:py-32 lg:px-8">
                    <div
                        class="relative py-24 overflow-hidden rounded-b-none max-h-96 rounded-3xl border-radius isolate sm:py-32">
                        <img
                            src="https://cdn.yuelili.com/docs/web/assert/anime-girl-dream.jpg"
                            alt=""
                            class="relative inset-0 object-cover object-right w-full h-full -translate-y-1/4 -z-10 md:object-center" />
                    </div>
                    <div class="flex text-center shadow-2xl sm:rounded-3xl">
                        <div class="py-6 border-2 border-r-0 rounded-es-lg basis-1/4">
                            <div class="text-sm">书籍数量</div>
                            <div class="text-lg font-bold">{{ statistic?.book_count }}</div>
                        </div>
                        <div class="p-6 border-2 border-r-0 basis-1/4">
                            <div class="text-sm">文章数量</div>
                            <div class="text-lg font-bold">{{ statistic?.document_count }}</div>
                        </div>
                        <div class="p-6 border-2 border-r-0 basis-1/4">
                            <div class="text-sm">历史访问人数</div>
                            <div class="text-lg font-bold">
                                {{ statistic?.historical_visitor_count }}
                            </div>
                        </div>
                        <div class="p-6 border-2 rounded-lg rounded-tr-none basis-1/4">
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
import { MetaData } from '@/models/base'
import { HomeStatistic } from '@/models/home'
import { onMounted, ref } from 'vue'
import Clock from '@/components/home/VClock.vue'

const yiyan = ref('')

const statistic = ref<HomeStatistic>()

const rndPostUrl = ref<string[]>(['intro'])

function addNotify() {
    Message({ message: '111', type: 'success', duration: 9999999 })
    Message({ message: '112', type: 'secondary', duration: 9999999 })
    Message({ message: '113', type: 'info', duration: 9999999 })
    Message({ message: '114', type: 'warn', duration: 9999999 })
    Message({ message: '115', type: 'error', duration: 9999999 })
    Message({ message: '116', type: 'contract', duration: 9999999 })
}

onMounted(async () => {
    await fetchBasic(statistic, new HomeStatistic(), fetchStatisticHome)
    await fetchBasic(yiyan, '最短的捷径就是绕远路。', fetchYiYan, '', 'hitokoto')

    const [ok, data] = await getRandPost()
    if (ok) {
        rndPostUrl.value = (data['data'] as MetaData).url.split('/')
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
