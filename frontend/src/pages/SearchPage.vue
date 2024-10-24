<template>
  <div class="py-8 h-full">
    <div class="flex h-full flex-col w-10/12 mx-auto items-center border-2 rounded-lg bg-base-100">
      <!-- Toolbar -->
      <div class="flex h-16 py-4 bg-base-200 w-full">
        <!-- 书籍设置 -->
        <div
          class="flex ml-auto pl-8 relative items-center justify-center text-sm join font-semibold text-center text-nowrap">
          <button class="btn btn-sm join-item">筛选:</button>

          <div class="dropdown m-1">
            <div tabindex="0" role="button" class="btn flex justify-center flex-nowrap">
              <span>选择书籍</span><span class="icon-[lucide--crosshair]"></span>
            </div>

            <ul
              tabindex="0"
              class="dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow">
              <li
                v-for="(option, index) in options"
                :key="index"
                @click="select(option)"
                class="flex items-center px-4 py-2 cursor-pointer">
                <span class="icon-[lucide--book-open] size-5"></span>
                <span class="w-24 text-sm truncate">{{ option.title }}</span>
              </li>
            </ul>
          </div>
        </div>

        <!-- 搜索框 -->
        <div class="flex items-center pl-8 w-full pr-16">
          <div class="relative flex items-center w-full">
            <span class="icon-[lucide--search] size-5 absolute left-3"></span>
            <span
              class="flex items-center h-10 pl-10 pr-4 w-full text-sm bg-transparent border-2 rounded-full hover:">
              <input
                v-model="search"
                placeholder="搜索..."
                @keydown.enter="handleSearch"
                class="flex-1 bg-transparent" />
            </span>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div class="search-content flex-1 flex w-full items-center justify-center">
        <transition name="result">
          <div v-if="searchResult.length" class="first:pt-2">
            <div v-for="(data, index) in searchResult" :key="data.url">
              <div class="p-2">
                <div class="relative p-4 border-b rounded-lg hover: hover:rounded-lg">
                  <a class="relative max-w-[1/2]" v-bind:href="conventLink(data)">
                    <div class="">
                      <div class="">
                        <span class="font-bold"
                          >{{ index + 1 + '.' }} {{ data.document_title }}</span
                        >
                        <div class="absolute top-0 right-4">
                          <span class="icon-[lucide--book-open] size-5"></span>
                          <span class=""> {{ data.category_title + '/' + data.book_title }}</span>
                        </div>
                      </div>

                      <div class="pt-6 description" v-html="highLight(data.content)"></div>
                    </div>
                  </a>
                </div>
              </div>
            </div></div
        ></transition>
        <div v-if="searchResult.length == 0" class="">没有找到任何文章~</div>
      </div>

      <!-- Footer -->
      <div class="bg-base-200 w-full h-8">
        <div v-if="searchResult.length > 0" class="pt-4 text-center border-t-4">
          {{ '搜索耗时: ' + searchConsume }}
        </div>
        <div v-else><div class="">Footer</div></div>
      </div>
    </div>
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

const navs = ref<Nav[]>([])

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
