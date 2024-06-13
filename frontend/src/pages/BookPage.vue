<template>
    <div class="container mx-auto px-16 py-16 md:max-w-screen-md">
        <div class="flex">
            <div class="book-logo md:basis-1/3 lg:basis-1/4">
                <img src="https://docs.yuelili.com/uploads/202305/1761bef95f7235da.jpg" alt="" />
            </div>
            <div class="flex flex-col px-6 select-none">
                <div class="py-3 border-b">
                    <BIconBook class="inline-block"></BIconBook>
                    <span class="pl-2">书籍名称: {{ route.params['category'] }}</span>
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
                            <!-- 本书尚未有{{ translate('locale') }}版本 -->
                        </div>
                        <router-link
                            v-else
                            v-for="(chapter, index) in bookDatas"
                            :key="index"
                            class="py-2 px-4 border-b hover:bg-theme-card border-dashed rounded-md flex items-center"
                            :to="{
                                name: 'post',
                                params: {
                                    book: getCat(chapter.url),
                                    document: getDocument(chapter.url)
                                }
                            }">
                            <span class="pl-2">{{
                                addZero(chapter.metadata.order, 3) + '. ' + chapter.metadata.title
                            }}</span></router-link
                        >
                    </div>
                </div>

                <div class="tab-item" v-if="tabId == 2">
                    <div class="w-full">
                        <textarea
                            name=""
                            id=""
                            cols="30"
                            rows="3"
                            class="w-full py-3 px-4 rounded-md min-h-12"
                            :placeholder="poem"></textarea>
                    </div>
                    <div class="mt-2 flex">
                        <button
                            class="btn bg-theme-primary-base hover:bg-theme-primary-hover px-3 py-1 ml-auto"
                            @click="postNewComment">
                            发布
                        </button>
                    </div>
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
                                <tr v-for="meta in metas.categorys" :key="meta.url_path">
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
                                <tr v-for="meta in metas.documents" :key="meta.url_path">
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
                        <button class="btn primary px-3 py-1 ml-auto" @click="postNewComment">
                            更新
                        </button>
                        <button class="btn warn px-3 py-1 ml-3" @click="postNewComment">
                            禁用
                        </button>
                        <button class="btn danger px-3 py-1 ml-3" @click="postNewComment">
                            删除
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { RouteParams, RouteLocationNormalizedLoaded } from 'vue-router'

import { addZero } from '@/utils'

import { BookData, LocalMetaDatas } from '@/models'
import { Message } from '@/plugins/message'
import {
    getBookData,
    getBookMeta,
    fetchStatisticBook,
    saveBookMeta,
    updateBookMeta,
    getRandNickname,
    postBookComment,
    getRandPoem
} from '@/handlers/index'

import { getDBBookData, addDBBookData } from '@/database'

import { getCat, getDocument } from '@/utils'

import { ref, onMounted, watch, computed } from 'vue'
import { useRoute } from 'vue-router'
import { basicStore } from '@/stores/index'

const basic = basicStore()
const locale = computed(() => basic.locale)
const isAdmin = computed(() => basic.isAdmin)
const nickname = computed(() => basic.nickname)

const metas = ref<LocalMetaDatas>(new LocalMetaDatas())

const translate = basic.translate

const tabId = ref(1)

const route = useRoute()
const bookReadCount = ref(0)
const bookChapterCount = ref(0)
const bookDocumentCount = ref(0)
const poem = ref('')

const bookDatas = ref<BookData[]>([])

watch(tabId, async (newVal: number, old: number) => {
    if (newVal == 3) {
        let [ok, data] = await getBookMeta(
            (route.params['slug'] as string[]).join('/'),
            locale.value
        )

        if (ok) {
            metas.value = data
        }
    }
})

async function saveMeta() {
    await saveBookMeta((route.params['slug'] as string[]).join('/'), locale.value, metas.value)
}

async function postNewComment() {}

async function updateMeta() {
    await updateBookMeta()
}

async function refreshBook(params: RouteParams) {
    // /book/Ae/basic

    const db_data = await getDBBookData(params['slug'] as string[], locale.value)

    if (db_data) {
        bookDatas.value = db_data.data
        console.log(bookDatas.value)
    } else {
        const [ok, data] = await getBookData((params['slug'] as string[]).join('/'), locale.value)

        if (ok) {
            bookDatas.value = data
            await addDBBookData(params['slug'] as string[], locale.value, data)
        } else {
            Message('未找到书籍数据', 'warn')
        }
    }

    const [ok2, statisticData] = await fetchStatisticBook(
        (route.params['slug'] as string[]).join('/'),
        locale.value
    )

    if (ok2) {
        bookReadCount.value = statisticData['read_count']
        bookChapterCount.value = statisticData['chapter_count']
        bookDocumentCount.value = statisticData['document_count']
    }
}

watch(route, async (val: RouteLocationNormalizedLoaded) => {
    refreshBook(val.params)
})

onMounted(async () => {
    refreshBook(route.params)

    poem.value = '山重水复疑无路，柳暗花明又一村。'

    const [ok, data] = await getRandPoem()

    if (ok) {
        poem.value = data['content']
    }
})
</script>
