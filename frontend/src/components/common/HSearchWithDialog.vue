<template>
    <transition name="fade">
        <div
            v-if="showSearchDialog"
            ref="searchDialogRef"
            class="search fixed w-screen h-full pt-8 z-50 flex flex-col align-center">
            <div
                @click.stop
                class="shadow-2xl dialog relative bg-theme-card dark:bg-dark-extra rounded-lg min-h-[16rem] max-h-[75%] w-[90%] left-[5%] md:w-[80%] md:left-[10%] lg:w-1/2 lg:left-1/4 max-h-1/2 z-50 top-18">
                <!-- 顶部工具 -->
                <div class="w-full flex-col h-16 z-50">
                    <div class="flex justify-around items-center border-b-2 mt-4 pb-4">
                        <!-- 书籍设置 -->
                        <div class="pl-4 md:pl-8 lg:pl-16 flex flex-col select-none">
                            <div
                                @mouseover="showSearchDropdown = true"
                                class="h-10 flex relative justify-center items-center text-center text-nowrap w-32 bg-transparent text-sm font-semibold">
                                <div class="pr-2">
                                    <BIconBook></BIconBook>
                                </div>
                                <span class="group" @mouseleave="showSearchDropdown = false">{{
                                    currentOption.name ? currentOption.name : '全站搜索'
                                }}</span>
                                <div
                                    v-show="currentOption.name.length > 0"
                                    class="pl-2 text-sm/[12px]"
                                    @click="currentOption.name = ''">
                                    <BIconX></BIconX>
                                </div>
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
                            <div class="relative flex ml-4 h-10 w-full">
                                <div
                                    class="absolute top-2/4 -mt-2 left-3 text-surface-400 dark:text-surface-600">
                                    <BIconSearch></BIconSearch>
                                </div>
                                <input
                                    v-model="search"
                                    placeholder="搜索..."
                                    @keydown.enter="handleSearch"
                                    class="pl-10 bg-transparent flex-1" />
                            </div>
                        </div>

                        <div class="toolbar ml-4 flex">
                            <div class="pr-2 lg:pr-4">
                                <div
                                    class="text-[1.25rem]"
                                    v-show="!pinSearchResult"
                                    @click="pinSearchResult = !pinSearchResult">
                                    <BIconLock></BIconLock>
                                </div>
                                <div
                                    class="text-[1.25rem]"
                                    v-show="pinSearchResult"
                                    @click="pinSearchResult = !pinSearchResult">
                                    <BIconUnlock></BIconUnlock>
                                </div>
                            </div>
                            <div class="pr-4">
                                <div class="text-[1.25rem]" @click="closeDialog(undefined)">
                                    <BIconX></BIconX>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- 结果展示 -->
                <div
                    class="content w-full flex pt-6 items-center justify-center min-h-[400px]"
                    style="height: calc(100% - 9rem)">
                    <transition name="result">
                        <div
                            v-if="searchResult.length"
                            class="h-full w-full overflow-scroll p-6 first:pt-2">
                            <div v-for="(data, index) in searchResult" :key="data.url">
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
                                                    class="description pt-6"
                                                    v-html="highLight(data.content)"></div>
                                            </div>
                                        </a>
                                    </div>
                                </div>
                            </div></div
                    ></transition>
                    <div v-if="searchResult.length == 0" class="">没有找到任何文章~</div>
                </div>
                <div v-if="searchResult.length > 0" class="text-center border-t-4 pt-4">
                    {{ '搜索耗时: ' + searchConsume }}
                </div>
            </div>
        </div></transition
    >
</template>

<script setup lang="ts">
// TODO 搜索结果分页

import { MetaData } from '@/models/base'
import { SearchData, Nav } from '@/models/home'
import { BIconBook, BIconLock, BIconSearch, BIconUnlock, BIconX } from 'bootstrap-icons-vue'

// 默认不pin
const pinSearchResult = ref(false)
const search = ref('')
const lastSearch = ref('')

const searchConsume = ref('')

const currentOption = ref<MetaData>(new MetaData())
const searchResult = ref<SearchData[]>([])

const searchDialogRef = ref(null)
const route = useRouter()

const showSearchDropdown = ref(false)
const showSearchDialog = defineModel('showSearchDialog', {
    type: Boolean,
    required: true
})

const navs = defineModel('navs', {
    type: Array as () => Nav[],
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
    showSearchDropdown.value = false
    currentOption.value = option
}

onMounted(async () => {
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
