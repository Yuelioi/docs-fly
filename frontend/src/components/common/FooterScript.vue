<template>
    <div></div>
</template>

<script setup lang="ts">
onMounted(async () => {
    const modules = import.meta.glob('@/assets/js/*.js')
    for (const path in modules) {
        try {
            await modules[path]()
            console.log(`${path} loaded`)
        } catch (error) {
            console.error(`Failed to load ${path}`, error)
        }
    }
})
</script>
