<template>
    <div id="nav" class="h-full my-3">
        <div class="flex w-full toolbar">
            <div v-if="chaptersData.length > virtual_limit_length">
                <nav
                    class="inline-flex -space-x-px rounded-md shadow-sm isolate"
                    aria-label="Pagination">
                    <a
                        @click.prevent="(currentPage -= 1), (isNavCollapsed = true)"
                        class="relative inline-flex items-center px-2 py-2 text-gray-400 rounded-l-md ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">
                        <span class="sr-only">Previous</span>
                        <div class="text-icon-md"><BIconCaretLeft></BIconCaretLeft></div>
                    </a>
                    <a
                        @click="(currentPage = 1), (isNavCollapsed = true)"
                        :class="currentPage == 1 ? ' ' : ''"
                        class="relative inline-flex items-center px-4 py-2 text-sm font-semibold select-none ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
                        >1</a
                    >
                    <a
                        :class="
                            currentPage != 1 &&
                            currentPage != Math.ceil(chaptersData.length / virtual_limit_length)
                                ? ' '
                                : ''
                        "
                        class="relative inline-flex items-center px-4 py-2 text-sm font-semibold focus:z-20 ring-1 ring-inset ring-gray-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                        >{{
                            Math.min(
                                Math.max(currentPage, 2),
                                Math.ceil(chaptersData.length / virtual_limit_length) - 1
                            )
                        }}</a
                    >
                    <a
                        :class="
                            currentPage == Math.ceil(chaptersData.length / virtual_limit_length)
                                ? ' '
                                : ''
                        "
                        @click="
                            (currentPage = Math.ceil(chaptersData.length / virtual_limit_length)),
                                (isNavCollapsed = true)
                        "
                        class="relative inline-flex items-center px-4 py-2 text-sm font-semibold select-none ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
                        >{{ Math.ceil(chaptersData.length / virtual_limit_length) }}</a
                    >
                    <a
                        @click.prevent="(currentPage += 1), (isNavCollapsed = true)"
                        class="relative inline-flex items-center px-2 py-2 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">
                        <span class="sr-only">Next</span>
                        <div class="text-icon-md"><BIconCaretRight></BIconCaretRight></div>
                    </a>
                </nav>
            </div>

            <div
                @click="handleCollapse"
                class="inline-flex ml-auto text-icon-md items-center justify-center ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">
                <div v-if="isNavCollapsed" class="px-2 py-2 rounded-r-md">
                    <BIconFilterLeft></BIconFilterLeft>
                </div>
                <div v-else class="px-2 py-2 rounded-r-md">
                    <BIconJustify></BIconJustify>
                </div>
            </div>
        </div>

        <div class="relative h-full overflow-y-scroll">
            <div class="absolute w-full py-3 pb-12 pl-1 mb-3 list text-ellipsis text-nowrap">
                <ul
                    v-for="(chapter, chapter_index) in filteredChapters"
                    :key="chapter.id"
                    class="mt-2 overflow-hidden scroll-item"
                    :class="chapter.metadata.status ? '' : 'hidden'"
                    :data-index="chapter.id">
                    <!-- 1. 没有章节 speedTree -->
                    <div v-if="chapter.documents.length == 0 && chapters.children.length == 0">
                        <router-link
                            :to="{
                                name: 'post',
                                params: {
                                    postPath: chapter.metadata.url.split('/')
                                }
                            }"
                            :data-index="chapter.id"
                            :key="chapter.metadata.url"
                            class="hover:border-slate-800 hover:pr-8 hover:bg-slate-300 dark:hover:border-slate-700 hover:rounded dark:hover:bg-slate-800">
                            <h5 class="mb-2 text-lg select-none">
                                {{ chapter.metadata.title }}
                            </h5></router-link
                        >
                    </div>

                    <!-- 2. 有章节 -->

                    <div v-else class="">
                        <!-- 2.1章节目录 -->
                        <router-link
                            :to="{
                                name: 'post',
                                params: {
                                    postPath: chapter.metadata.url.split('/')
                                }
                            }"
                            :key="chapter.metadata.url">
                            <h5 class="mb-4 text-lg font-bold">
                                {{ chapter_index + 1 + '. ' + chapter.metadata.title }}
                            </h5>
                        </router-link>

                        <Transition name="list">
                            <li v-if="!chapter.collapsed">
                                <div
                                    v-for="(section, section_id) in chapter.children"
                                    :key="section_id">
                                    <router-link
                                        :to="{
                                            name: 'post',
                                            params: {
                                                postPath: section.metadata.url.split('/')
                                            }
                                        }"
                                        :key="chapter.metadata.url"
                                        ><span> {{ section.metadata.title }}</span></router-link
                                    >

                                    <div
                                        v-for="(document, document_index) in section.documents"
                                        :key="document.url"
                                        class="flex border-l border-l-slate-400 last:pb-4">
                                        <router-link
                                            :to="{
                                                name: 'post',
                                                params: {
                                                    postPath: document.url.split('/')
                                                }
                                            }"
                                            :key="chapter_index + document_index"
                                            class="pl-6 -ml-px hover: hover:pl-7 hover:pr-4 hover: hover:border-l-4"
                                            ><span>
                                                {{
                                                    addZero(document_index + 1, 2) +
                                                    '. ' +
                                                    document.title
                                                }}</span
                                            ></router-link
                                        >
                                    </div>
                                </div>

                                <div
                                    v-for="(document, document_index) in chapter.documents"
                                    :key="document.url"
                                    class="flex border-l border-l-slate-400">
                                    <router-link
                                        :to="{
                                            name: 'post',
                                            params: {
                                                postPath: document.url.split('/')
                                            }
                                        }"
                                        :key="chapter_index + document_index"
                                        class="pl-4 -ml-px hover: hover:pl-3.5 hover:pr-4 hover: hover:border-l-4"
                                        @click.stop
                                        ><span>{{
                                            addZero(document_index + 1, 2) + '. ' + document.title
                                        }}</span></router-link
                                    >
                                </div>
                            </li></Transition
                        >
                    </div>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ChapterData, Chapter } from '@/models/post'
import { MetaData } from '@/models/base'

import { addZero } from '@/utils'
import { BIconCaretLeft, BIconCaretRight, BIconFilterLeft, BIconJustify } from 'bootstrap-icons-vue'

const chapters = defineModel('chapters', { type: Object as () => Chapter, required: true })

const isNavCollapsed = ref(false)

const virtual_limit_length = 50 // 多少个数据使用分页

const chaptersData = ref<ChapterData[]>([])
const currentPage = ref(1)

// TODO 如果滑动过多会丢失

/**
 * 计算后的数据
 */
const filteredChapters = computed(() => {
    if (chaptersData.value.length < virtual_limit_length) {
        return chaptersData.value
    } else {
        return chaptersData.value.slice(
            Math.max((currentPage.value - 1) * 50, 0),
            Math.min(currentPage.value * 50, chaptersData.value.length - 1)
        ) as ChapterData[]
    }
})

function handleCollapse() {
    isNavCollapsed.value = !isNavCollapsed.value
    filteredChapters.value.forEach((ele) => {
        ele.collapsed = isNavCollapsed.value
    })
}

watch(chapters, () => {
    init()
})

function init() {
    const cacheChaptersData: ChapterData[] = []

    chapters.value.documents.forEach((doc: MetaData, index: number) => {
        const chapterData = {
            metadata: doc,
            documents: [],
            children: [],
            filepath: '',
            collapsed: chapters.value.children.length > virtual_limit_length,
            id: index
        } as ChapterData

        cacheChaptersData.push(chapterData)
    })

    chapters.value.children.forEach((chapter: Chapter, index: number) => {
        const chapterData = {
            ...chapter,
            collapsed: false,
            id: index
        } as ChapterData

        cacheChaptersData.push(chapterData)
    })

    chaptersData.value = cacheChaptersData
}

onMounted(async () => {
    init()
})
</script>
<style scoped>
.hidden {
    display: none;
}

ul:hover::-webkit-scrollbar {
    background-color: #5858584d;
}
#nav ul:hover::-webkit-scrollbar-thumb {
    background-color: #27272791 !important;
}

#nav ul::-webkit-scrollbar-thumb {
    background: transparent;
}

.list-enter-active,
.list-leave-active {
    transition: all 0.25s ease;
}

.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(-40px);
}
</style>
