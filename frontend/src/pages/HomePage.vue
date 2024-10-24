<template>
  <div class="flex flex-col -mt-1 h-full bg-base-200 overflow-hidden">
    <!-- 主页手风琴 -->
    <div class="w-full h-[20vh] lg:h-[50vh] relative">
      <div class="flex items-center justify-center h-full select-none">
        <span
          class="absolute z-20 text-sm sm:text-lg font-bold md:text-2xl lg:text-3xl animated-text text-slate-200 top-1/3"
          >{{ yiyan }}</span
        >
      </div>
      <img
        class="absolute top-0 object-cover w-full h-full"
        src="https://cdn.yuelili.com/docs/web/assets/banner-anime-girl-roller.jpg"
        alt="" />
    </div>
    <!-- 主页手风琴结束 -->

    <!--  -->
    <div class="mx-auto my-8 w-full px-8 sm:w-10/12 md:w-9/12 lg:w-8/12 xl:w-6/12">
      <div
        class="relative mt-4 flex flex-col h-full sm:flex-row justify-stretch overflow-hidden rounded-lg shadow-2xl isolate">
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
        <div class="sm:relative h-64 sm:flex-1 w-full sm:block top-0">
          <img
            class="object-cover w-full h-full ring-1 ring-white/10"
            src="https://cdn.yuelili.com/docs/web/assets/anime-girl.jpg"
            alt="App screenshot"
            width="1824"
            height="1080" />
        </div>
        <div class="mx-auto sm:flex-1">
          <div class="h-24 sm:h-full flex items-center justify-center">
            <router-link
              class="btn btn-primary btn-outline text-base-200 px-3.5 py-2.5"
              :to="{
                name: 'books'
              }">
              <span class="">开始阅读 →</span>
            </router-link>
          </div>
        </div>
      </div>
    </div>
    <!--  -->

    <!-- 主页信息统计区域 -->
    <div class="mx-auto my-8 w-full px-8 sm:w-10/12 md:w-9/12 lg:w-8/12 xl:w-6/12">
      <div class="shadow-2xl">
        <div class="rounded-b-none rounded-3xl border-radius isolate">
          <img
            src="https://cdn.yuelili.com/docs/web/assets/anime-girl-dream.jpg"
            alt=""
            class="relative inset-0 object-cover object-top w-full h-64 -translate-y-0 sm:h-64 -z-10" />
        </div>
        <div class="flex text-sm text-center *:pt-2">
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
    </div>

    <!-- 主页信息统计区域结束 -->
    <!-- 时间 -->
    <keep-alive>
      <VClock class="max-w-[48rem] my-8" />
    </keep-alive>
    <!-- 时间结束 -->
  </div>
</template>

<script setup lang="ts">
import { MetaData } from '@/models/base'
import { HomeStatistic } from '@/models/home'

import { ref, onMounted } from 'vue'

const yiyan = ref('')
const statistic = ref<HomeStatistic>()

const rndPostUrl = ref<string[]>(['intro'])

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
