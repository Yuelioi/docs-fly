<template>
    <!--container divide space -->
    <div class="container w-20 space-y-0 divide-y-8 divide-yellow-400">
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
    </div>

    <!-- 文字隔断  line-clamp限制行数 truncate限制一行-->
    <div class="w-48 truncate line-clamp-2">
        Lorem ipsum dolor sit amet consectetur, adipisicing elit. Iure debitis illum soluta
        quibusdam quam officia dolores in, suscipit, sapiente deserunt provident pariatur ipsum quis
        nobis numquam praesentium. Amet, velit minus?
    </div>
    <!-- 过渡渐变 -->
    <div
        class="w-full h-48 bg-gradient-to-r from-20% from-orange-500 to-slate-400 via-sky-400"></div>

    <!-- ring 使用box-shadow的伪边框 -->
    <button class="w-24 h-12 bg-sky-600 ring-4 ring-red-600">Tex</button>

    <AsyncComp></AsyncComp>
</template>

<script setup lang="ts">
const AsyncComp = defineAsyncComponent({
    // the loader function
    loader: () => new Promise((res, rej) => h('div', ['<span>Loading</span>'])),

    // A component to use while the async component is loading
    loadingComponent: h('div', ['<span>Loading</span>']),
    // Delay before showing the loading component. Default: 200ms.
    delay: 200,

    // A component to use if the load fails
    errorComponent: h('div', ['<span>Loading</span>']),
    // The error component will be displayed if a timeout is
    // provided and exceeded. Default: Infinity.
    timeout: 3000
})
</script>
