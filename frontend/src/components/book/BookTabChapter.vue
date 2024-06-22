<template>
    <div class="flex flex-col">
        <div v-if="bookDatas.length == 0">本书尚未有{{ translate('locale') }}版本</div>
        <router-link
            v-else
            v-for="(chapter, index) in sortedBookDatas"
            :key="index"
            class="flex items-center px-4 py-2 border-b border-dashed rounded-md hover:bg-theme-card"
            :to="{
                name: 'post',
                params: {
                    postPath: chapter.url.split('/')
                }
            }">
            <span class="pl-2">{{
                addZero(chapter.metadata.order, 3) + '. ' + chapter.metadata.title
            }}</span></router-link
        >
    </div>
</template>

<script setup lang="ts">
import type { RouteParams, RouteLocationNormalizedLoaded } from 'vue-router'

import { addZero } from '@/utils'
import { BookData } from '@/models/book'
const bookDatas = ref<BookData[]>([])

const basic = basicStore()
const route = useRoute()
const locale = computed(() => basic.locale)
const translate = basic.translate

const sortedBookDatas = computed(() => {
    return bookDatas.value.slice().sort((pre, next) => pre.metadata.order - next.metadata.order)
})

async function refreshBookChapter(params: RouteParams) {
    const db_data = await getDBBookData(params['bookPath'] as string[], locale.value)

    if (db_data) {
        bookDatas.value = db_data.data
    } else {
        const [ok, data] = await getBookData(
            (params['bookPath'] as string[]).join('/'),
            locale.value
        )

        if (ok) {
            bookDatas.value = data['data']
            await addDBBookData(params['bookPath'] as string[], locale.value, data['data'])
        } else {
            bookDatas.value = []
            await Message({ message: '未找到书籍数据', type: 'warn' })
        }
    }
}

watch(locale, async () => {
    await refreshBookChapter(route.params)
})

watch(route, async (val: RouteLocationNormalizedLoaded) => {
    refreshBookChapter(val.params)
})

onMounted(async () => {
    refreshBookChapter(route.params)
})
</script>
