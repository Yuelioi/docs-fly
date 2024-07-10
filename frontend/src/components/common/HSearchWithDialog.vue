<template>
    <div class="">
        <!-- 顶部工具 -->

        <VDialog v-model:show="showSearchDialog">
            <template #header>
                <div class="flex items-center justify-between w-full">
                    <!-- 书籍设置 -->
                    <div
                        class="flex relative items-center justify-center text-sm join font-semibold text-center text-nowrap">
                        <button class="btn btn-sm join-item">分类:</button>
                        <div
                            class="join-item rounded-md h-8 w-24 flex items-center justify-center group">
                            <span class="text-sm truncate bg-neutral-content px-2">{{
                                currentOption.name ? currentOption.name : '选择分类'
                            }}</span>
                            <BIconCaretDown class="ml-2"></BIconCaretDown>

                            <div
                                class="absolute top-full left-0 mt-1 scale-0 group-hover:scale-100 transition-transform duration-300 shadow-lg bg-white dark:bg-dark-light z-50">
                                <ul
                                    class="overflow-y-scroll max-h-60 ring-1 ring-neutral ring-opacity-5 focus:outline-none">
                                    <li
                                        v-for="(option, index) in options"
                                        :key="index"
                                        @click="select(option)"
                                        class="flex items-center px-4 py-2 hover:bg-gray-200 cursor-pointer">
                                        <BIconBook class="mr-2 text-sm"></BIconBook>
                                        <span class="w-24 text-sm truncate">{{
                                            option.title
                                        }}</span>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>

                    <!-- 搜索框 -->
                    <div class="flex items-center">
                        <div class="relative">
                            <input
                                v-model="search"
                                placeholder="搜索..."
                                @keydown.enter="handleSearch"
                                class="flex-1 bg-transparent" />
                        </div>
                    </div>

                    <div class="flex ml-4 toolbar">
                        <div class="text-icon-md">
                            <BIconSearch></BIconSearch>
                        </div>
                        <div
                            class="text-icon-md"
                            v-if="!pinSearchResult"
                            @click="pinSearchResult = !pinSearchResult">
                            <BIconLock></BIconLock>
                        </div>
                        <div
                            class="text-icon-md"
                            v-else
                            @click="pinSearchResult = !pinSearchResult">
                            <BIconUnlock></BIconUnlock>
                        </div>
                    </div></div
            ></template>

            <div class="flex items-center justify-center">
                <transition name="result">
                    <div v-if="searchResult.length" class="first:pt-2">
                        <div v-for="(data, index) in searchResult" :key="data.url">
                            <div class="p-2">
                                <div
                                    class="relative p-4 border-b rounded-lg hover: hover:rounded-lg">
                                    <a class="relative max-w-[1/2]" v-bind:href="conventLink(data)">
                                        <div class="">
                                            <div class="">
                                                <span class="font-bold"
                                                    >{{ index + 1 + '.' }}
                                                    {{ data.document_title }}</span
                                                >
                                                <div class="absolute top-0 right-4">
                                                    <BIconBook class="pr-2"></BIconBook>

                                                    <span class="">
                                                        {{
                                                            data.category_title +
                                                            '/' +
                                                            data.book_title
                                                        }}</span
                                                    >
                                                </div>
                                            </div>

                                            <div
                                                class="pt-6 description"
                                                v-html="highLight(data.content)"></div>
                                        </div>
                                    </a>
                                </div>
                            </div>
                        </div></div
                ></transition>
                <div v-if="searchResult.length == 0" class="">没有找到任何文章~</div>
            </div>
            <template #footer>
                <div v-if="searchResult.length > 0" class="pt-4 text-center border-t-4">
                    {{ '搜索耗时: ' + searchConsume }}
                </div></template
            >
        </VDialog>
    </div>
</template>

<script setup lang="ts">
// TODO 搜索结果分页

import { MetaData } from '@/models/base'
import { SearchData, Nav } from '@/models/home'
import {
    BIconBook,
    BIconLock,
    BIconSearch,
    BIconUnlock,
    BIconX,
    BIconCaretDown
} from 'bootstrap-icons-vue'

// 默认不pin
const pinSearchResult = ref(false)
const search = ref('')
const lastSearch = ref('')

const searchConsume = ref('')

const currentOption = ref<MetaData>(new MetaData())
const searchResult = ref<SearchData[]>([])

const showSearchDialog = defineModel('showSearchDialog')

const navs = defineModel('navs', {
    type: Array as () => Nav[],
    required: true
})

function highLight(content: string) {
    return content.replace(lastSearch.value, `<span class="highlight">${lastSearch.value}</span>`)
}

const conventLink = function (data: SearchData) {
    const linkList = [data.url]
    const filteredLink = linkList.filter(function (item: string) {
        return item != ''
    })

    return '/post/' + filteredLink.join('/')
}

const options = computed(() => {
    const res: MetaData[] = []

    const all = new MetaData()
    all.name = ''
    all.title = '全部'

    res.push(all)
    for (let nav of navs.value) {
        for (const book of nav.children) {
            const data = new MetaData()
            data.status = book.status
            data.icon = book.icon
            data.name = nav.metadata.name + '/' + book.name
            data.title = nav.metadata.title + '/' + book.title
            res.push(data)
        }
    }
    return res
})

async function handleSearch() {
    if (search.value == '') {
        searchResult.value = []
        return
    }

    const [ok, data] = await fetchKeyword(currentOption.value.name, search.value, 1, 20)

    if (ok) {
        const msTotal =
            new Date(data['server_time'] as string).getTime() -
            new Date(data['client_time'] as string).getTime()

        const seconds = Math.floor(msTotal / 1000)
        const ms = msTotal % 1000

        if (seconds > 1) {
            searchConsume.value = `${seconds}秒${ms}毫秒`
        } else {
            searchConsume.value = `${ms}毫秒`
        }

        searchResult.value = data['data']
        lastSearch.value = search.value
    } else {
        searchResult.value = []
        searchConsume.value = ''
    }
}

function select(option: MetaData) {
    currentOption.value = option
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
    transition: all 0.375s;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
    transform: translateY(-300px);
}

.result-enter-active {
    transition: all 0.375s;
}

.result-enter-from,
.result-leave-to {
    opacity: 0;
    transform: translateX(-30px);
}
</style>
