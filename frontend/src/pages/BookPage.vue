<template>
    <div class="container mx-auto px-16 py-16 md:max-w-screen-md">
        <div class="flex">
            <div class="book-logo md:basis-1/3 lg:basis-1/4">
                <img src="https://docs.yuelili.com/uploads/202305/1761bef95f7235da.jpg" alt="" />
            </div>
            <div class="flex flex-col px-6 select-none">
                <div class="py-3 border-b">
                    <BIconBook class="inline-block"></BIconBook>
                    <span class="pl-2">书籍名称: {{ bookStatistic.bookTitle }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconGraphUpArrow class="inline-block"></BIconGraphUpArrow
                    ><span class="pl-2">阅读次数: {{ bookStatistic.readCount }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconJournal class="inline-block"></BIconJournal>
                    <span class="pl-2">章节数量: {{ bookStatistic.chapterCount }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconFiletypeDoc class="inline-block"></BIconFiletypeDoc>
                    <span class="pl-2">文章数量: {{ bookStatistic.documentCount }}</span>
                </div>
            </div>
        </div>

        <div class="tab pt-16">
            <div class="border-b-2 border-theme-muted">
                <div class="text-sm font-medium text-center">
                    <ul class="flex flex-wrap -mb-px select-none">
                        <li :class="['me-2', 'group', { active: tabId === 1 }]" @click="tabId = 1">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                >章节</span
                            >
                        </li>
                        <li :class="['me-2', 'group', { active: tabId === 2 }]" @click="tabId = 2">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >评论</span
                            >
                        </li>
                        <li
                            v-if="isAdmin"
                            :class="['me-2', 'group', { active: tabId === 3 }]"
                            @click="tabId = 3">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >编辑元数据</span
                            >
                        </li>
                        <li
                            v-if="isAdmin"
                            @click="tabId = 4"
                            :class="['me-2', 'group', { active: tabId === 4 }]">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >书籍设置</span
                            >
                        </li>
                    </ul>
                </div>
            </div>

            <div class="mt-4">
                <div class="tab-item" v-if="tabId == 1">
                    <div class="flex flex-col">
                        <div class="" v-if="bookDatas.length == 0">
                            本书尚未有{{ translate('locale') }}版本
                        </div>
                        <router-link
                            v-else
                            v-for="(chapter, index) in sortedBookDatas"
                            :key="index"
                            class="py-2 px-4 border-b hover:bg-theme-card border-dashed rounded-md flex items-center"
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
                </div>

                <div class="tab-item" v-if="tabId == 2">
                    <VComment></VComment>
                </div>
                <div class="tab-item" v-if="tabId == 3">
                    <div class="toolbar flex pb-4">
                        <button
                            type="button"
                            class="btn bg-theme-primary-base hover:bg-theme-primary-hover ml-auto px-3 py-1"
                            @click="updateMeta">
                            更新
                        </button>
                        <button
                            type="button"
                            class="btn bg-theme-primary-base hover:bg-theme-primary-hover ml-3 px-3 py-1"
                            @click="saveMeta">
                            保存
                        </button>
                    </div>
                    <div>
                        <table class="table-auto w-full border-collapse">
                            <thead>
                                <tr>
                                    <th class="border px-4 py-2">ID</th>
                                    <th class="border px-4 py-2">{{ translate('displayName') }}</th>
                                    <th class="border px-4 py-2">{{ translate('order') }}</th>
                                    <th class="border px-4 py-2">{{ translate('status') }}</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="meta in metas.categorys" :key="meta.url">
                                    <td class="px-4 py-2">
                                        <input type="text" class="" disabled v-model="meta.name" />
                                    </td>
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class="p-2 border rounded-sm w-full"
                                            v-model="meta.title" />
                                    </td>
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class="p-2 border rounded-sm w-full"
                                            v-model.number="meta.order" />
                                    </td>
                                    <td class="px-4 py-2 text-center">
                                        <input
                                            type="checkbox"
                                            id="checkbox"
                                            v-model="meta.status"
                                            :true-value="false"
                                            :false-value="true" />
                                    </td>
                                </tr>
                            </tbody>

                            <tbody>
                                <tr v-for="meta in metas.documents" :key="meta.url">
                                    <td class="px-4 py-2">
                                        <input type="text" class="" disabled v-model="meta.name" />
                                    </td>
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class="p-2 border rounded-sm w-full"
                                            v-model="meta.title" />
                                    </td>
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class="p-2 border rounded-sm w-full"
                                            v-model.number="meta.order" />
                                    </td>
                                    <td class="px-4 py-2 text-center">
                                        <input
                                            type="checkbox"
                                            id="checkbox"
                                            v-model="meta.status"
                                            :true-value="false"
                                            :false-value="true" />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>

                <div class="tab-item" v-if="tabId == 4">
                    <div class="toolbar flex pb-4">
                        <button class="btn primary px-3 py-1 ml-auto">更新</button>
                        <button class="btn warn px-3 py-1 ml-3">禁用</button>
                        <button class="btn danger px-3 py-1 ml-3">删除</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { RouteParams, RouteLocationNormalizedLoaded } from 'vue-router'

import { addZero } from '@/utils'

import { BookData, BookStatistic, LocalMetaDatas } from '@/models'
import { Message } from '@/plugins/message'
import {
    getBookData,
    getBookMeta,
    fetchStatisticBook,
    saveBookMeta,
    updateBookMeta
} from '@/services/index'

import VComment from '@/components/common/VComment.vue'

import { getDBBookData, addDBBookData } from '@/database'

import { ref, onMounted, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import { basicStore } from '@/stores/index'
import { BIconBook, BIconFiletypeDoc, BIconGraphUpArrow, BIconJournal } from 'bootstrap-icons-vue'

const basic = basicStore()
const locale = computed(() => basic.locale)
const isAdmin = computed(() => basic.isAdmin)

const metas = ref<LocalMetaDatas>(new LocalMetaDatas())

const translate = basic.translate

const tabId = ref(1)

const route = useRoute()
const bookStatistic = ref<BookStatistic>(new BookStatistic())

const bookDatas = ref<BookData[]>([])
const sortedBookDatas = computed(() => {
    return bookDatas.value.slice().sort((pre, next) => pre.metadata.order - next.metadata.order)
})

watch(tabId, async (newVal: number) => {
    if (newVal == 3) {
        let [ok, data] = await getBookMeta(
            (route.params['bookPath'] as string[]).join('/'),
            locale.value
        )

        if (ok) {
            metas.value = data['data']
        }
    }
})

watch(locale, async () => {
    await refreshBook(route.params)
})

async function saveMeta() {
    const [ok, data] = await saveBookMeta(
        (route.params['bookPath'] as string[]).join('/'),
        locale.value,
        metas.value
    )
    if (ok) {
        Message({ message: '保存成功' })
    } else {
        Message({ message: '保存失败', type: 'warn' })
    }
}

async function updateMeta() {
    await updateBookMeta()
}

async function refreshBook(params: RouteParams) {
    // /book/Ae/basic

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

    const [ok2, data2] = await fetchStatisticBook(
        (route.params['bookPath'] as string[]).join('/'),
        locale.value
    )

    if (ok2) {
        const statisticData = data2['data']
        bookStatistic.value.bookTitle = statisticData['book_title']
        bookStatistic.value.readCount = statisticData['read_count']
        bookStatistic.value.chapterCount = statisticData['chapter_count']
        bookStatistic.value.documentCount = statisticData['document_count']
    } else {
        bookStatistic.value = new BookStatistic()
    }
}

watch(route, async (val: RouteLocationNormalizedLoaded) => {
    refreshBook(val.params)
})

onMounted(async () => {
    refreshBook(route.params)
})
</script>
