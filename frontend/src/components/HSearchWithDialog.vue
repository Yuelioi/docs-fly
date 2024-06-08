<template>
    <transition name="fade">
        <div
            v-if="showSearchDialog"
            ref="searchDialogRef"
            class="search fixed w-screen h-full pt-8 z-50 flex flex-col align-center">
            <div
                @click.stop
                class="dialog relative bg-theme-card dark:bg-dark-extra rounded-lg min-h-[16rem] max-h-[75%] w-[90%] left-[5%] md:w-[80%] md:left-[10%] lg:w-1/2 lg:left-1/4 max-h-1/2 z-50 top-18">
                <!-- 顶部工具 -->
                <div class="w-full flex-col h-16 z-50">
                    <div class="flex justify-around border-b-2 mt-4 pb-4">
                        <!-- 书籍设置 -->
                        <div class="pl-4 md:pl-8 lg:pl-16 flex flex-col select-none">
                            <div
                                @mouseover="showSearchDropdown = true"
                                class="h-10 flex relative justify-center items-center text-center text-nowrap w-32 bg-transparent text-sm font-semibold">
                                <i class="pi pi-book pr-2"></i>
                                <span class="group" @mouseleave="showSearchDropdown = false">{{
                                    currentOption.name ? currentOption.name : '全站搜索'
                                }}</span>
                                <i
                                    v-show="currentOption.name.length > 0"
                                    class="pi pi-times pl-2 text-sm/[12px]"
                                    @click="currentOption.name = ''"></i>
                                <div
                                    v-if="showSearchDropdown"
                                    class="absolute z-[100] top-[3rem] rounded h-32">
                                    <ul
                                        @mouseleave="showSearchDropdown = false"
                                        class="bg-theme-base dark:bg-dark-light overflow-y-scroll mt-1 max-h-56 w-full rounded-md py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                        <li
                                            v-for="(option, index) in options"
                                            class="px-6 py-3 last:pb-4 hover:bg-theme-primary-hover hover:rounded-lg w-full whitespace-normal"
                                            :key="index"
                                            @click="select(option)">
                                            <span>{{ option.title }}</span>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                        <!-- 搜索框 -->
                        <div class="pl-8 flex-1 flex items-center">
                            <div class="relative flex ml-4 h-10">
                                <i
                                    class="pi pi-search absolute top-2/4 -mt-2 left-3 text-surface-400 dark:text-surface-600" />
                                <input
                                    v-model="search"
                                    placeholder="搜索..."
                                    @keydown.enter="handleSearch"
                                    class="pl-10 bg-transparent" />
                            </div>
                        </div>

                        <div class="toolbar absolute top-6 right-4 flex">
                            <div class="pr-2 lg:pr-4">
                                <i
                                    class="pi pi-lock"
                                    v-show="!pinSearchResult"
                                    @click="pinSearchResult = !pinSearchResult">
                                </i>
                                <i
                                    class="pi pi-lock-open"
                                    v-show="pinSearchResult"
                                    @click="pinSearchResult = !pinSearchResult">
                                </i>
                            </div>
                            <div class="pr-2">
                                <i class="pi pi-times" @click="closeDialog(undefined)"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- 结果展示 -->
                <div
                    class="content w-full flex pt-6 items-center justify-center"
                    style="height: calc(100% - 9rem)">
                    <transition name="result">
                        <div
                            v-if="searchResult.result.length"
                            class="h-full w-full overflow-scroll p-6 first:pt-2">
                            <div
                                v-for="(data, index) in searchResult.result"
                                :key="data.document_name">
                                <div class="p-2">
                                    <div
                                        class="hover:bg-theme-primary-hover relative border-b rounded-lg hover:rounded-lg p-4">
                                        <a
                                            class="relative max-w-[1/2]"
                                            v-bind:href="conventLink(data)"
                                            @click.prevent="jumpToDocument(conventLink(data))">
                                            <div class="w-[90%]">
                                                <div class="text-lg">
                                                    <span class="font-bold"
                                                        >{{ index + 1 + '.' }}
                                                        {{ data.document_title }}</span
                                                    >
                                                    <div class="absolute right-4 top-0">
                                                        <i class="pi pi-book pr-2"></i>
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
                                                    class="description pt-6"
                                                    v-html="highLight(data.content)"></div>
                                            </div>
                                        </a>
                                    </div>
                                </div>
                            </div></div
                    ></transition>
                    <div v-if="searchResult.result.length == 0" class="">没有找到任何文章~</div>
                </div>
                <div v-if="searchResult.result.length > 0" class="text-center">
                    {{ '搜索耗时: ' + searchResult.search_time }}
                </div>
            </div>
        </div></transition
    >
</template>

<script setup lang="ts">
// TODO 搜索结果分页
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

import { useRouter } from 'vue-router'

import { fetchKeyword, fetchSearchOptions } from '@/handlers'
import { MetaData, SearchData, SearchOption, SearchResult } from '@/models'

import { getSearchOption, addSearchOption } from '@/database'

// 默认不pin
const pinSearchResult = ref(false)
const search = ref('')
const lastSearch = ref('')

const currentOption = ref<MetaData>(new MetaData())
const searchResult = ref<SearchResult>(new SearchResult())

const searchDialogRef = ref(null)
const route = useRouter()

const searchOptions = ref<SearchOption[]>([])

const showSearchDropdown = ref(false)
const showSearchDialog = defineModel('showSearchDialog', {
    type: Boolean,
    required: true
})

function closeDialog(event: KeyboardEvent | any = undefined) {
    // 如果是单击事件 并且作用在dialog上 就关闭对话框
    if (event && event.type == 'click' && searchDialogRef.value == event.target) {
        showSearchDialog.value = false
    }
    // 如果是esc关闭对话框
    if (!event || event.key === 'Escape') {
        showSearchDialog.value = false
    }
}

function highLight(content: string) {
    return content.replace(lastSearch.value, `<span class="highlight">${lastSearch.value}</span>`)
}

function jumpToDocument(link: string) {
    if (!pinSearchResult.value) {
        showSearchDialog.value = false
    }
    route.push(link)
}

const conventLink = function (data: SearchData) {
    const linkList = [
        data.category_name,
        data.book_name,
        data.locale,
        data.chapter_name,
        data.section_name,
        data.document_name
    ]
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
    for (let cat of searchOptions.value) {
        for (const book of cat.children) {
            const data = new MetaData()
            data.hidden = book.hidden
            data.icon = book.icon
            data.name = cat.name + '/' + book.name
            data.title = cat.title + '/' + book.title
            res.push(data)
        }
    }
    return res
})

async function handleSearch() {
    if (search.value == '') {
        searchResult.value = new SearchResult()
        return
    }

    let option_list
    if (currentOption.value.name == '') {
        option_list = ['', '']
    } else {
        option_list = currentOption.value.name.split('/')
    }

    const [ok, data] = await fetchKeyword(option_list[0], option_list[1], search.value)

    if (ok) {
        searchResult.value = data
        lastSearch.value = search.value
    } else {
        searchResult.value = new SearchResult()
    }
}

function select(option: MetaData) {
    showSearchDropdown.value = false
    currentOption.value = option
}

onMounted(async () => {
    const search_option = await getSearchOption()
    if (search_option) {
        searchOptions.value = search_option.data
    } else {
        const [ok, data] = await fetchSearchOptions()
        if (ok) {
            searchOptions.value = data
            await addSearchOption(searchOptions.value)
        } else {
            searchOptions.value = []
        }
    }

    window.addEventListener('keydown', closeDialog)
    document.addEventListener('click', closeDialog)
})

onBeforeUnmount(() => {
    window.removeEventListener('keydown', closeDialog)
    document.removeEventListener('click', closeDialog)
})
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
