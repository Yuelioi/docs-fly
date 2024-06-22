<template>
    <div class="container px-16 py-16 mx-auto">
        <BookHeader></BookHeader>

        <div class="pt-16 tab">
            <div class="border-b-2 border-theme-muted">
                <div class="text-sm font-medium text-center">
                    <ul class="flex flex-wrap -mb-px select-none">
                        <li
                            :class="['me-2', 'group', { active: currentTab == BookTabChapter }]"
                            @click="currentTab = BookTabChapter">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                >章节</span
                            >
                        </li>
                        <li
                            :class="['me-2', 'group', { active: currentTab == VComment }]"
                            @click="currentTab = VComment">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >评论</span
                            >
                        </li>
                        <li
                            v-if="isAdmin"
                            :class="['me-2', 'group', { active: currentTab == BookTabMeta }]"
                            @click="currentTab = BookTabMeta">
                            <span
                                class="inline-block p-4 group-[.active]:text-theme-primary group-[.active]:border-theme-primary group-[.active]:border-b-2 rounded-t-lg"
                                aria-current="page"
                                >编辑元数据</span
                            >
                        </li>
                    </ul>
                </div>
            </div>

            <div class="mt-4">
                <keep-alive> <div class="tab-item" is="vue:currentTab"></div> </keep-alive>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const basic = basicStore()
const isAdmin = computed(() => basic.isAdmin)

import VComment from '@/components/common/VComment.vue'
import BookTabChapter from '@/components/book/BookTabChapter.vue'
import BookTabMeta from '@/components/book/BookTabMeta.vue'

const currentTab = shallowRef<Component>(BookTabChapter)
</script>
