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
import { Chapter, type Toc } from '@/models/post'
import type { RouteLocationNormalizedLoaded, RouteParams } from 'vue-router'

import { Message } from '@/plugins/message'

const postContent = ref('')
const postHtml = ref('')
const toc = ref<Toc[]>([])

const chapters = ref<Chapter>(new Chapter())

const route = useRoute()

/**
 * @param params
 * @param reload :是否更新章节信息
 */
async function refreshBookContent(params: RouteParams, reload: boolean = true) {
    // 已有数据存入数据库

    // 更新章节
    if (reload) {
        const path = (params['postPath'] as string[]).slice(0, 3).join('/')
        const data: any = await getPostChapterData(path)
        if (data) {
            chapters.value = data['data']
        } else {
            const [ok, data] = await getChapter((params['postPath'] as string[]).join('/'))

            if (ok) {
                chapters.value = data['data']
                await addPostChapterData(
                    chapters.value.metadata.url,
                    JSON.parse(JSON.stringify(chapters.value))
                )
            } else {
                chapters.value = new Chapter()
            }
        }
    }

    // 更新文章
    const [ok, data] = await getPost(
        (params['postPath'] as string[]).join('/'),
        params['document'] as string
    )

    if (ok) {
        postContent.value = data['data']['content_markdown']
        postHtml.value = data['data']['content_html']
        const tocData = JSON.parse(data['data']['toc'])
        toc.value = tocData
    } else {
        await Message({ message: '获取文章失败', type: 'error' })
        postContent.value = ''
        postHtml.value = ''
        toc.value = []
    }
}

watch(route, async (val: RouteLocationNormalizedLoaded, oldVal: RouteLocationNormalizedLoaded) => {
    let reload = true
    if (oldVal.params['postPath'] == val.params['postPath']) {
        reload = false
    }

    await refreshBookContent(val.params, reload)
    await AddVisitorLog((val.params['postPath'] as string[]).join('/'))
})

onBeforeMount(async () => {
    await refreshBookContent(route.params)
    await AddVisitorLog((route.params['postPath'] as string[]).join('/'))
})
</script>

<style scoped></style>
