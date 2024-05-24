<template>
    <div class="container mx-auto px-16 py-16 md:max-w-screen-md">
        <div class="flex">
            <div class="book-logo md:basis-1/3 lg:basis-1/4">
                <img src="https://docs.yuelili.com/uploads/202305/1761bef95f7235da.jpg" alt="" />
            </div>
            <div class="flex flex-col px-6">
                <div class="py-3 border-b">
                    <BIconBook class="inline-block"></BIconBook>
                    <span class="pl-2">书籍名称: {{ bookDatas.book.display_name }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconGraphUpArrow class="inline-block"></BIconGraphUpArrow
                    ><span class="pl-2">阅读次数: {{ bookReadCount }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconJournal class="inline-block"></BIconJournal>
                    <span class="pl-2">章节数量: {{ bookChapterCount }}</span>
                </div>
                <div class="py-3 border-b">
                    <BIconFiletypeDoc class="inline-block"></BIconFiletypeDoc>
                    <span class="pl-2">文章数量: {{ bookDocumentCount }}</span>
                </div>
            </div>
        </div>

        <div class="tab pt-16">
            <div class="border-b">
                <div
                    class="text-sm font-medium text-center text-gray-500 border-b dark:text-gray-400 dark:border-gray-700">
                    <ul class="flex flex-wrap -mb-px">
                        <li :class="['me-2', 'group', { active: tabId === 1 }]" @click="tabId = 1">
                            <span
                                class="inline-block p-4 group-[.active]:text-blue-600 group-[.active]:border-blue-600 group-[.active]:border-b-2 rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
                                >章节</span
                            >
                        </li>
                        <!-- TODO 评论区 -->
                        <li :class="['me-2', 'group', { active: tabId === 2 }]" @click="tabId = 2">
                            <span
                                class="inline-block p-4 group-[.active]:text-blue-600 group-[.active]:border-blue-600 group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >评论</span
                            >
                        </li>
                        <!-- TODO 编辑元数据 -->
                        <li
                            v-if="isAdmin"
                            :class="['me-2', 'group', { active: tabId === 3 }]"
                            @click="tabId = 3">
                            <span
                                class="inline-block p-4 group-[.active]:text-blue-600 group-[.active]:border-blue-600 group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >编辑元数据</span
                            >
                        </li>
                    </ul>
                </div>
            </div>
            <div class="mt-4">
                <div class="tab-item" v-if="tabId == 1">
                    <div class="flex flex-col">
                        <div class="" v-if="bookDatas.children.length == 0">
                            本书尚未有{{ translate(locale) }}版本
                        </div>
                        <router-link
                            v-else
                            v-for="(chapter, index) in bookDatas.children"
                            :key="index"
                            class="py-2 px-4 border-b hover:bg-slate-100 border-dashed dark:hover:bg-slate-700 rounded-md flex items-center"
                            :to="{
                                name: 'post',
                                params: {
                                    category: bookDatas.category.identity,
                                    book: bookDatas.book.identity,
                                    locale: locale,
                                    chapter: chapter.chapter,
                                    section: chapter.section,
                                    document: chapter.document
                                }
                            }">
                            <span class="pl-2">{{ chapter.display_name }}</span></router-link
                        >
                    </div>
                </div>

                <div class="tab-item" v-if="tabId == 2">评论</div>
                <div class="tab-item" v-if="tabId == 3">
                    <div class="toolbar flex flex-row-reverse">
                        <button type="button" class="btn-primary px-3 py-1" @click="saveMeta">
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
                                    <th class="border px-4 py-2">{{ translate('hidden') }}</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="meta in metas" :key="meta.identity">
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class=""
                                            disabled
                                            v-model="meta.identity" />
                                    </td>
                                    <td class="px-4 py-2">
                                        <input
                                            type="text"
                                            class="p-2 border rounded-sm w-full"
                                            v-model="meta.display_name" />
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
                                            v-model="meta.hidden"
                                            :true-value="false"
                                            :false-value="true" />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { RouteParams, RouteLocationNormalizedLoaded } from 'vue-router'

import { BookData, MetaData } from '@/models'
import { Message } from '@/plugins/message'
import { fetchBook, fetchBookMeta, fetchStatisticBook, saveBookMeta } from '@/handlers/index'

import { getBookData, addBookData } from '@/database'

import { ref, onMounted, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import { basicStore } from '@/stores/index'

const basic = basicStore()
const locale = computed(() => basic.locale)
const isAdmin = computed(() => basic.isAdmin)

const metas = ref<MetaData[]>([])

const translate = basic.translate

const tabId = ref(3)

const route = useRoute()
const bookReadCount = ref(0)
const bookChapterCount = ref(0)
const bookDocumentCount = ref(0)

const bookDatas = ref<BookData>(new BookData())

async function saveMeta() {
    await saveBookMeta(
        route.params['category'] as string,
        route.params['book'] as string,
        locale.value,
        metas.value
    )
}

async function refreshBook(params: RouteParams) {
    // /book/Ae/basic

    const db_data = await getBookData(params)
    if (db_data) {
        bookDatas.value = db_data.data
    } else {
        const [ok, data] = await fetchBook(
            params['category'] as string,
            params['book'] as string,
            locale.value
        )

        if (ok) {
            bookDatas.value = data
            await addBookData(params, data)
        } else {
            Message('未找到书籍数据', 'warn')
        }
    }

    const [ok2, statisticData] = await fetchStatisticBook(
        route.params['category'] as string,
        route.params['book'] as string
    )

    if (ok2) {
        bookReadCount.value = statisticData['read_count']
        bookChapterCount.value = statisticData['chapter_count']
        bookDocumentCount.value = statisticData['document_count']
    }

    let ok3
    ;[ok3, metas.value] = await fetchBookMeta(
        params['category'] as string,
        params['book'] as string,
        locale.value
    )

    if (ok3) {
        console.log(metas.value)
    }
}

watch(route, async (val: RouteLocationNormalizedLoaded) => {
    refreshBook(val.params)
})

onMounted(async () => {
    refreshBook(route.params)
})
</script>
