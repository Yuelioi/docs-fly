<template>
    <div class="">
        <!-- 书籍章节大纲 Chapter -->
        <aside
            class="hidden mt-1 lg:block fixed lg:w-[15rem] xl:w-[18rem] z-20 top-[3.8125rem] border-r-2 right-auto pb-16 pl-8 h-full">
            <PostChapter v-model:chapters="chapters" />
        </aside>

        <!-- 文章 Document Content-->
        <article class="sm:w-full sm:px-[2rem] lg:px-[15rem] xl:px-[20rem] my-8">
            <PostContent v-model:postContent="postContent" v-model:postHtml="postHtml" />
        </article>

        <!-- 文章目录 Toc -->
        <div
            class="pl-6 mt-1 hidden lg:block fixed lg:w-[15rem] xl:w-[20rem] right-2 border-l-2 w-[14rem] top-[3.8125rem] h-full">
            <PostToc :toc="toc" />
        </div>
    </div>
</template>

<script setup lang="ts">
import type { ChapterInfo, Toc } from '@/models'
import type { RouteLocationNormalizedLoaded, RouteParams } from 'vue-router'

import { Message } from '@/plugins/message'

import { ref, onBeforeMount, watch, computed } from 'vue'
import { useRoute } from 'vue-router'

import { basicStore } from '@/stores'

import PostChapter from '@/components/PostChapter.vue'
import PostContent from '@/components/PostContent.vue'
import PostToc from '@/components/PostToc.vue'

import { getPostChapterData, addPostChapterData } from '@/database/index'
import { fetchPost, fetchChapter } from '@/handlers/index'
import { AddVisitorLog } from '@/handlers/index'

const basic = basicStore()

const locale = computed(() => basic.locale)

const postContent = ref('')
const postHtml = ref('')
const toc = ref<Toc[]>([])

const chapters = ref<ChapterInfo[]>([])

const route = useRoute()

/**
 * @param params
 * @param reload :是否更新章节信息
 */
async function refreshBookContent(params: RouteParams, reload: boolean = true) {
    const result: any = await getPostChapterData(params)

    // 已有数据存入数据库

    // 更新章节
    if (reload) {
        if (result) {
            chapters.value = result['data']
        } else {
            const [ok, data] = await fetchChapter(
                params['category'] as string,
                params['book'] as string,
                locale.value
            )

            if (ok) {
                chapters.value = data
                await addPostChapterData(params, JSON.parse(JSON.stringify(chapters.value)))
            } else {
                chapters.value = []
            }
        }
    }

    // 更新文章
    const [ok, data] = await fetchPost(
        params['category'] as string,
        params['book'] as string,
        params['locale'] as string,
        params['chapter'] as string,
        params['document'] as string
    )

    if (ok) {
        postContent.value = data['content_markdown']
        postHtml.value = data['content_html']
        toc.value = JSON.parse(data['toc'])
    } else {
        Message('获取文章失败', 'error')
        postContent.value = ''
        postHtml.value = ''
        toc.value = []
    }
}

watch(route, async (val: RouteLocationNormalizedLoaded, oldVal: RouteLocationNormalizedLoaded) => {
    let reload = true
    if (
        oldVal.params['category'] == val.params['category'] &&
        oldVal.params['book'] == val.params['book'] &&
        oldVal.params['locale'] == val.params['locale']
    ) {
        reload = false
    }

    await refreshBookContent(val.params, reload)
    await AddVisitorLog(val.params, val.fullPath)
})

onBeforeMount(async () => {
    await refreshBookContent(route.params)
    await AddVisitorLog(route.params, route.fullPath)
})
</script>

<style scoped></style>
