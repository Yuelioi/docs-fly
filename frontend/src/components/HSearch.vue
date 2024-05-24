<template>
    <div class="justify-center items-center relative flex h-16 select-none">
        <i
            class="pi pi-search absolute top-2/4 -mt-2 left-3 text-surface-400 dark:text-surface-600" />
        <span
            @click="showSearchDialog = true"
            class="pl-10 pr-4 h-10 bg-transparent text-sm rounded-full border-2 border-neutral-300 hover:border-sky-400 flex items-center">
            搜索文档...
        </span>
    </div>
</template>

<script setup lang="ts">
const showSearchDialog = defineModel('showSearchDialog')
</script>
<style scoped></style>
