<template>
    <div>测试</div>
    {{ chapter }}
</template>

<script setup lang="ts">
const chapter = ref()
onMounted(async () => {
    chapter.value = await getChapter('Ue/blueprint/en/AI/AIMoveTo.md')
})
</script>
