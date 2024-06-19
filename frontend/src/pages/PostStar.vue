<template>
    <div class="w-1/2 ml-[25%]">
        <div class="mx-2 my-8">
            <div class="px-4 py-4 rounded-md bottom-2 bg-theme-warn-base">
                注意, 此收藏夹存于本地浏览器, 请勿删除浏览器数据
            </div>
        </div>

        <ul role="list" class="mt-8 divide-y">
            <router-link
                class="flex justify-between px-4 py-4 my-4 rounded gap-x-6 hover:bg-theme-primary-hover"
                v-for="(star, index) in stars"
                :key="star.createdTime.toString()"
                :to="{
                    name: 'post',
                    params: {
                        postPath: star.postPath
                    }
                }">
                <div class="flex min-w-0 gap-x-8">
                    <div class="flex-auto min-w-0">
                        <div class="flex items-center">
                            <p
                                class="text-sm font-semibold leading-6 text-gray-900 dark:text-slate-200">
                                {{ index + 1 + '. ' + star.params }}
                            </p>
                            <BIconBook class="pl-4 pr-2"></BIconBook>
                            <span class="text-sm leading-6">
                                {{ star.params + ' / ' + star.params }}
                            </span>
                        </div>

                        <p
                            class="mt-1 text-xs leading-5 text-gray-500 truncate dark:text-slate-400">
                            这里放文章概述{{ star.params }}
                        </p>
                    </div>
                </div>
                <div class="shrink-0 sm:flex sm:flex-col sm:items-end">
                    <i class="pi pi-times text-sm/[18px]" @click.prevent="deleteData(star.key)"></i>
                    <p class="mt-1 text-xs leading-5 text-gray-500">
                        收藏日期
                        <time datetime="2023-01-23T13:23Z">{{ formatDate(star.createdTime) }}</time>
                    </p>
                </div>
            </router-link>
        </ul>
    </div>
</template>

<script setup lang="ts">
// TODO 星星结果分页

import { getPostStarsData, deletePostStarData } from '@/database/star'
import { PostStar } from '@/models/star'

import { formatDate } from '@/utils'

const stars = ref<PostStar[]>([])

async function deleteData(key: string) {
    await deletePostStarData(key)
    await Message({ message: '删除成功' })
    refresh()
}
async function refresh() {
    const data = await getPostStarsData()
    if (data) {
        stars.value = data
    }
}

onMounted(async () => {
    refresh()
})
</script>

<style scoped></style>
