<template>
    <div class="justify-center items-center relative flex h-16 select-none">
        <div class="absolute top-2/4 -mt-2 left-3 text-surface-400 dark:text-surface-600">
            <BIconSearch></BIconSearch>
        </div>
        <span
            @click="showSearchDialog = true"
            class="pl-10 pr-4 h-10 bg-transparent text-sm rounded-full border-2 border-theme-base hover:border-theme-primary flex items-center">
            搜索文档...
        </span>
    </div>
</template>

<script setup lang="ts">
import { BIconSearch } from 'bootstrap-icons-vue'

const showSearchDialog = defineModel('showSearchDialog')
</script>
<style scoped></style>
