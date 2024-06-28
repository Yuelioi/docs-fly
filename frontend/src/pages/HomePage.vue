<template>
    <div class="">
        <button @click="addNotify">添加事件</button>
        <div class="flex flex-col -mt-1">
            <!-- 主页手风琴 -->
            <div class="w-full h-[33vh] relative">
                <div class="flex items-center justify-center h-full select-none">
                    <span
                        class="absolute z-20 text-sm sm:text-base animated-text text-slate-200 top-1/3"
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
            <div class="container -mt-[15vh]">
                <div class="w-full">
                    <div
                        class="relative mt-4 overflow-hidden rounded-lg shadow-2xl sm:flex isolate bg-theme-base">
                        <svg
                            viewBox="0 0 1024 1024"
                            class="absolute left-1/2 top-1/2 sm:left-1/4 sm:-translate-y-4 -z-10 size-[64rem] sm:size-[32rem] -translate-y-1/2 [mask-image:radial-gradient(closest-side,white,transparent)]"
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
                        <div class="py-6 mx-auto text-center sm:w-1/2">
                            <h2 class="font-bold tracking-tight sm:text-lg sm:py-4">
                                欢迎来到月离文档站.<br />Welcome to Yueli Docs
                            </h2>
                            <p class="mt-6 leading-1 text-black-300">
                                国内最大CG文档网站<br />
                                Largest CG Documentation Site.
                            </p>
                            <div class="flex items-center justify-center mt-10 gap-x-6">
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
                        <div class="relative hidden w-1/2 sm:block">
                            <img
                                class="object-cover w-full h-full ring-1 ring-white/10"
                                src="https://cdn.yuelili.com/docs/web/assert/anime-girl.jpg"
                                alt="App screenshot"
                                width="1824"
                                height="1080" />
                        </div>
                    </div>
                </div>
            </div>
            <!--  -->

            <!-- 主页信息统计区域 -->
            <div class="container my-8">
                <div class="rounded-b-none rounded-3xl border-radius isolate">
                    <img
                        src="https://cdn.yuelili.com/docs/web/assert/anime-girl-dream.jpg"
                        alt=""
                        class="relative inset-0 object-cover object-top w-full h-full -translate-y-0 sm:h-64 sm:mt-16 -z-10" />
                </div>
                <div class="flex text-sm text-center shadow-2xl *:pt-2">
                    <div class="border-2 border-r-0 rounded-es-lg basis-1/4">
                        <div class="">书籍数量</div>
                        <div class="font-bold">{{ statistic?.book_count }}</div>
                    </div>
                    <div class="border-2 border-r-0 basis-1/4">
                        <div class="">文章数量</div>
                        <div class="font-bold">{{ statistic?.document_count }}</div>
                    </div>
                    <div class="border-2 border-r-0 basis-1/4">
                        <div class="text-sm">历史访问</div>
                        <div class="font-bold">
                            {{ statistic?.historical_visitor_count }}
                        </div>
                    </div>
                    <div class="border-2 basis-1/4">
                        <div class="text-sm">今日访问</div>
                        <div class="text-lg font-bold">
                            {{ statistic?.today_visitor_count }}
                        </div>
                    </div>
                </div>
            </div>
            <!-- 主页信息统计区域结束 -->
            <!-- 时间 -->
            <keep-alive>
                <VClock class="max-w-[48rem]" />
            </keep-alive>
            <!-- 时间结束 -->
        </div>
    </div>
</template>

<script setup lang="ts">
import { MetaData } from '@/models/base'
import { HomeStatistic } from '@/models/home'

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
    await fetchHandleBasic(statistic, new HomeStatistic(), fetchStatisticHome)
    await fetchHandleBasic(yiyan, '最短的捷径就是绕远路。', fetchYiYan, '', 'hitokoto')

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
