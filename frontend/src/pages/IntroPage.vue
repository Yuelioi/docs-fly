<template>
    <!--container divide space -->
    <div class="container w-20 space-y-0 divide-y-8 divide-yellow-400">
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
    </div>
    <button @click="show = true">111</button>
    <Suspense><AsyncComp v-if="show"></AsyncComp></Suspense>

    <!-- <VTest v-if="show"></VTest> -->
</template>

<script setup lang="ts">
const show = ref(false)
const AsyncComp = defineAsyncComponent({
    loader: () => import('@/components/VTest.vue'),
    loadingComponent: h('div', ['<span>Loading</span>']),
    delay: 2000,
    timeout: 3000
})
</script>
