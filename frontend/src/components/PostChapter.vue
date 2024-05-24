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
                    :ref="(el) => setChapterRef(el, chapter)"
                    @click.prevent="chapter.collapsed = !chapter.collapsed">
                    <!-- 情况1. 没有章节 speedTree -->
                    <div v-if="chapter.document.order">
                        <router-link
                            :to="{
                                name: 'post',
                                params: {
                                    category: chapter.category.identity,
                                    chapter: chapter.chapter.identity,
                                    locale: locale,
                                    document: chapter.document.identity
                                }
                            }"
                            :key="chapter.document.identity"
                            class="hover:border-slate-800 hover:pr-8 hover:bg-slate-300 dark:hover:border-slate-700 text-slate-700 dark:text-slate-400 hover:rounded dark:hover:bg-slate-800">
                            {{ chapter.document.display_name }}</router-link
                        >
                    </div>

                    <!-- 情况2 有章节 -->

                    <h5
                        class="select-none text-lg font-bold mb-4 lg:mb-3 text-slate-900 dark:text-slate-200">
                        {{ chapter.id + 1 + '. ' + chapter.chapter.display_name }}
                    </h5>
                    <Transition name="list">
                        <li v-show="!chapter.collapsed">
                            <div
                                v-for="(section, section_id) in chapter.sections"
                                :key="section_id">
                                {{ '小节' + section.display_name }}
                            </div>

                            <div
                                v-for="(document, document_index) in chapter.documents"
                                :key="chapter_index + document_index"
                                class="flex">
                                <router-link
                                    :to="{
                                        name: 'post',
                                        params: {
                                            category: category,
                                            book: book,
                                            locale: locale,
                                            chapter: chapter.chapter.identity,
                                            document: document.identity
                                        }
                                    }"
                                    :key="chapter_index + document_index"
                                    class="block border-l pl-4 -ml-px border-slate-300 hover:pl-3.5 hover:border-slate-800 hover:pr-8 hover:bg-slate-300 dark:hover:border-slate-700 text-slate-700 dark:text-slate-400 hover:rounded dark:hover:bg-slate-800"
                                    @click.stop
                                    ><span>{{
                                        addZero(document_index + 1, 2) +
                                        '. ' +
                                        document.display_name
                                    }}</span></router-link
                                >
                            </div>
                        </li></Transition
                    >
                </ul>
            </div>
            <div class="fill" :style="{ height: fillHeigh + 'px' }"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ChapterData, ChapterInfo } from '@/models'
import { ref, computed, watch, onMounted, nextTick } from 'vue'

import { useRoute } from 'vue-router'
const route = useRoute()

const category = route.params['category']
const book = route.params['book']

import { addZero } from '@/utils'
import { basicStore } from '@/stores'
const basic = basicStore()

const chapters = defineModel('chapters', { type: Array as () => ChapterInfo[], required: true })

const isNavCollapsed = ref(false)
const locale = computed(() => basic.locale)

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
) // 容器高度

function setChapterRef(el: any, chapter: ChapterData) {
    chapter.ref = el
}

// 填充总高度
const fillHeigh = computed(() => {
    if (chaptersData.value.length < virtual_limit_length) {
        return 0
    }
    return Math.max(
        chaptersData.value.length * chapter_collapsed_height.value,
        containerHeigh.value
    )
})

const filteredChapters = computed(() => {
    if (chaptersData.value.length < virtual_limit_length) {
        return chaptersData.value
    } else {
        return chaptersData.value.slice(start.value, end.value) as ChapterData[]
    }
})

function handleScroll(event: UIEvent) {
    console.log(event)

    event.preventDefault()
    // 小数据直接渲染 不处理
    if (chaptersData.value.length < virtual_limit_length) {
        return
    }

    const container = scrollContainerRef.value

    if (!container) {
        return
    }

    // 当前滚动高度
    const currentScrollTop = container.scrollTop

    if (Math.abs(currentScrollTop - lastScrollTop.value) > chapter_collapsed_height.value * 50) {
        const newStart = Math.floor(currentScrollTop / chapter_collapsed_height.value)
        listTop.value = currentScrollTop - chapter_collapsed_height.value * 50
        start.value = newStart - 50
        end.value = newStart + 100
        lastScrollTop.value = currentScrollTop
        console.log(currentScrollTop, listTop.value, newStart, start.value, end.value)
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
        update()
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
function update() {
    for (let i = 0; i < filteredChapters.value.length; i++) {
        const chapter = filteredChapters.value[i]
        if (chapter.ref) {
            // 大数据只计算折叠高度
            if (chapter.collapsed) {
                chapter_collapsed_height.value = calculateHeight(chapter.ref)
                console.log('已更新')
                return
            }
        }
    }
}

function init() {
    const cacheChaptersData: ChapterData[] = []

    chapters.value.forEach((chapter: ChapterInfo, index: number) => {
        const chapterData = {
            ...chapter,
            collapsed: true,
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
