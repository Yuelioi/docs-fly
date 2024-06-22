<template>
    <div class="flex">
        <div class="w-1/2 book-logo">
            <img :src="bookStatistic.bookCover" alt="" />
        </div>
        <div class="flex flex-col space-y-6 *:border-b-2 select-none">
            <div class="border-b-2 border-theme-base">
                <BIconBook class="inline-block"></BIconBook>
                <span class="pl-2">书籍名称: {{ bookStatistic.bookTitle }}</span>
            </div>
            <div>
                <BIconGraphUpArrow class="inline-block"></BIconGraphUpArrow
                ><span class="pl-2">阅读次数: {{ bookStatistic.readCount }}</span>
            </div>
            <div>
                <BIconJournal class="inline-block"></BIconJournal>
                <span class="pl-2">章节数量: {{ bookStatistic.chapterCount }}</span>
            </div>
            <div>
                <BIconFiletypeDoc class="inline-block"></BIconFiletypeDoc>
                <span class="pl-2">文章数量: {{ bookStatistic.documentCount }}</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { BIconBook, BIconFiletypeDoc, BIconGraphUpArrow, BIconJournal } from 'bootstrap-icons-vue'
import { BookStatistic } from '@/models/book'
const bookStatistic = ref<BookStatistic>(new BookStatistic())

const basic = basicStore()
const route = useRoute()
const locale = computed(() => basic.locale)

onMounted(async () => {
    const [ok2, data2] = await fetchStatisticBook(
        (route.params['bookPath'] as string[]).join('/'),
        locale.value
    )

    if (ok2) {
        const statisticData = data2['data']

        bookStatistic.value.bookCover = statisticData['book_cover']
        bookStatistic.value.bookTitle = statisticData['book_title']
        bookStatistic.value.readCount = statisticData['read_count']
        bookStatistic.value.chapterCount = statisticData['chapter_count']
        bookStatistic.value.documentCount = statisticData['document_count']
    } else {
        bookStatistic.value = new BookStatistic()
    }
})
</script>
