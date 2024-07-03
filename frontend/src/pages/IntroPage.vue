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
    <input type="text" v-model="num" />

    <div class="w-24 h-24" :style="getStyle()"></div>

    <canvas ref="container" width="1000" height="1000"></canvas>

    <canvas ref="container2" width="1000" height="1000" style="width: 100%; height: auto"></canvas>
</template>

<script setup lang="ts">
const show = ref(false)
import { Rive } from '@rive-app/canvas'

const container = ref()
const container2 = ref()

function getStyle() {
    return { background: 'red' }
}

const num = ref(0)
const AsyncComp = defineAsyncComponent({
    loader: () => import('@/components/VTest.vue'),
    loadingComponent: h('div', ['<span>Loading</span>']),
    delay: 2000,
    timeout: 3000
})

onMounted(() => {
    new Rive({
        src: '/rive/water_bubble.riv',
        canvas: container.value,
        stateMachines: 'State Machine 1',
        autoplay: true
    })
    new Rive({
        src: 'https://cdn.motiondesign.school/uploads/2024/03/Belly.riv',
        canvas: container2.value,
        stateMachines: 'State Machine 1',
        autoplay: true
    })
})
</script>
