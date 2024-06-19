<template>
    <div class="relative flex items-center justify-center h-16 select-none">
        <div class="absolute -mt-2 top-2/4 left-3 text-surface-400 dark:text-surface-600">
            <BIconSearch></BIconSearch>
        </div>
        <span
            @click="showSearchDialog = true"
            class="flex items-center h-10 pl-10 pr-4 text-sm bg-transparent border-2 rounded-full border-theme-base hover:border-theme-primary">
            搜索文档...
        </span>
    </div>
</template>

<script setup lang="ts">
import { BIconSearch } from 'bootstrap-icons-vue'

const showSearchDialog = defineModel('showSearchDialog')
</script>
<style scoped></style>
