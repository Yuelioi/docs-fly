<template>
    <div id="nav" class="my-3 lg:text-sm lg:leading-6 h-full">
        <div class="absolute right-4 top-6 z-50" @click="handleCollapse">
            <i class="pi" :class="isNavCollapsed ? 'pi-align-right' : 'pi-align-justify'"></i>
        </div>

        <div
            ref="scrollContainerRef"
            class="relative h-full overflow-y-scroll"
            @scroll="handleScroll">
            <div
                class="list mb-3 absolute py-3 pb-12 pl-1 w-full text-ellipsis text-nowrap"
                ref="listRef"
                :style="{ top: listTop + 'px' }">
                <ul
                    v-for="(chapter, chapter_index) in filteredChapters"
                    :key="chapter.id"
                    class="scroll-item lg:mt-4 overflow-hidden"
                    :class="chapter.metadata.status ? '' : 'hidden'"
                    :data-index="chapter.id"
                    @click.prevent="chapter.collapsed = !chapter.collapsed">
                    <!-- 情况1. 没有章节 speedTree -->
                    <div v-if="chapter.documents.length == 0">
                        <router-link
                            :to="{
                                name: 'post',
                                params: {
                                    slug: getCat(chapter.metadata.filepath),
                                    document: getDocument(chapter.metadata.filepath)
                                }
                            }"
                            :data-index="chapter.id"
                            :key="chapter.metadata.filepath"
                            class="hover:border-slate-800 hover:pr-8 hover:bg-slate-300 dark:hover:border-slate-700 text-theme-text-base hover:rounded dark:hover:bg-slate-800">
                            <h5
                                class="select-none text-lg font-bold mb-4 lg:mb-3 text-theme-text-base">
                                {{ chapter.metadata.title }}
                            </h5></router-link
                        >
                    </div>

                    <!-- 情况2 有章节 -->

                    <div v-else class="">
                        <h5 class="select-none text-lg font-bold mb-4 lg:mb-3 text-theme-text-base">
                            {{ chapter.id + 1 + '. ' + chapter.metadata.title }}
                        </h5>

                        <Transition name="list">
                            <li v-show="!chapter.collapsed">
                                <div
                                    v-for="(section, section_id) in chapter.children"
                                    :key="section_id">
                                    <span> {{ section.metadata.title }}</span>
                                </div>

                                <div
                                    v-for="(document, document_index) in chapter.documents"
                                    :key="document.filepath"
                                    class="flex border-l border-l-slate-400">
                                    <router-link
                                        :to="{
                                            name: 'post',
                                            params: {
                                                slug: getCat($route.fullPath),
                                                document: getDocument($route.fullPath)
                                            }
                                        }"
                                        :key="chapter_index + document_index"
                                        class="pl-4 -ml-px hover:bg-theme-card hover:pl-3.5 hover:pr-4 hover:border-theme-primary hover:border-l-4"
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
            <div class="fill" :style="{ height: fillHeigh + 'px' }"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ChapterData, Chapter, MetaData } from '@/models'
import { ref, computed, watch, onMounted, nextTick } from 'vue'

import { addZero, getCat, getDocument } from '@/utils'
import { useRoute } from 'vue-router'

const chapters = defineModel('chapters', { type: Chapter, required: true })

const isNavCollapsed = ref(false)

const scrollContainerRef = ref<HTMLElement>()
const listRef = ref<HTMLElement>()

// 处理大数据
const listTop = ref(0)
const chapter_collapsed_height = ref(40) // 折叠后的章节高度(默认显示就是40)

const start = ref(0)
const end = ref(100)
const lastScrollTop = ref(0)

const virtual_limit_length = 50 // 多少个数据使用虚拟列表

const chaptersData = ref<ChapterData[]>([])

const containerHeigh = computed(() =>
    scrollContainerRef.value ? scrollContainerRef.value.clientHeight : 800
)

/**
 * 填充体高度
 * 填充体: 用来撑开整个元素, 撑开滚动条
 * */
const fillHeigh = computed(() => {
    if (chaptersData.value.length < virtual_limit_length) {
        return 0
    }
    return Math.max(
        chaptersData.value.length * chapter_collapsed_height.value,
        containerHeigh.value
    )
})

/**
 * 计算后的数据
 */
const filteredChapters = computed(() => {
    if (chaptersData.value.length < virtual_limit_length) {
        return chaptersData.value
    } else {
        return chaptersData.value.slice(
            Math.max(start.value, 0),
            Math.min(end.value, chaptersData.value.length - 1)
        ) as ChapterData[]
    }
})

/**
 * 获取可视区第一个与最后一个元素的id
 */
function getVisibleFirstLastElementId(container: HTMLElement) {
    const ul = listRef.value
    let container_start = 0
    let container_end = 0
    if (ul) {
        const containerRect = container.getBoundingClientRect() // 获取父级元素的范围
        const elements = ul.querySelectorAll('ul') as NodeListOf<HTMLUListElement>
        let firstVisible: HTMLUListElement | null = null
        let lastVisible: HTMLUListElement | null = null

        for (let i = 0; i < elements.length - 1; i++) {
            const el = elements[i]
            const rect = el.getBoundingClientRect()
            if (rect.top >= containerRect.top && rect.bottom <= containerRect.bottom) {
                if (!firstVisible) {
                    firstVisible = el
                }

                lastVisible = el
            }
        }

        if (firstVisible && firstVisible.dataset.index) {
            container_start = parseInt(firstVisible.dataset.index) // @ts-ignore
        }

        if (lastVisible && lastVisible.dataset.index) {
            container_end = parseInt(lastVisible.dataset.index)
        }
    }
    return { container_start, container_end }
}

/**
 * 处理滚动条
 **/
function handleScroll() {
    // 小数据直接渲染 不处理
    if (chaptersData.value.length < virtual_limit_length) {
        return
    }

    const container = scrollContainerRef.value

    if (!container) {
        return
    }

    let { container_start, container_end } = getVisibleFirstLastElementId(container)
    const currentScrollTop = container.scrollTop

    // 向上滚动
    if (currentScrollTop - lastScrollTop.value < 0 && container_start - start.value < 5) {
        start.value = container_start - 85
        end.value = container_end + 15
        lastScrollTop.value = currentScrollTop
    }

    // 向下滚动
    if (currentScrollTop - lastScrollTop.value > 0 && container_end > end.value - 5) {
        start.value = container_start - 15
        end.value = container_end + 85
        lastScrollTop.value = currentScrollTop
    }
}

function handleCollapse() {
    isNavCollapsed.value = !isNavCollapsed.value
    filteredChapters.value.forEach((ele) => {
        ele.collapsed = isNavCollapsed.value
    })
}

watch(chapters, () => {
    init()
    nextTick(() => {
        updateDefaultChapterHeight()
    })
})

function calculateHeight(ele: Element) {
    const style = window.getComputedStyle(ele)
    var marginTop = parseInt(style.marginTop)
    var marginBottom = parseInt(style.marginBottom)
    var totalHeight = ele.getBoundingClientRect().height + marginTop + marginBottom
    return totalHeight
}

// 更新一下章节高度真实值
function updateDefaultChapterHeight() {
    for (let i = 0; i < filteredChapters.value.length; i++) {
        const chapter = filteredChapters.value[i]
        if (chapter.ref) {
            if (chapter.collapsed) {
                chapter_collapsed_height.value = calculateHeight(chapter.ref)
                return
            }
        }
    }
}

function init() {
    const cacheChaptersData: ChapterData[] = []

    chapters.value.documents.forEach((doc: MetaData, index: number) => {
        const chapterData = {
            metadata: doc,
            documents: [],
            children: [],
            ref: document.createElement('div'),
            collapsed: chapters.value.children.length > virtual_limit_length,
            id: index
        } as ChapterData

        cacheChaptersData.push(chapterData)
    })

    chapters.value.children.forEach((chapter: Chapter, index: number) => {
        const chapterData = {
            ...chapter,
            collapsed: chapters.value.children.length > virtual_limit_length,
            id: index
        } as ChapterData

        cacheChaptersData.push(chapterData)
    })

    chaptersData.value = cacheChaptersData
}

onMounted(async () => {
    const route = useRoute()
    console.log(route.fullPath)

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
