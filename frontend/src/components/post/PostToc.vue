<template>
    <div id="toc" class="overflow-scroll h-[calc(100%-4rem)]">
        <div
            class="p-1 pl-4 text-sm border-l-2 cursor-pointerselect-none toc-item hover: hover: hover: hover:px-6"
            v-for="(item, i) in toc"
            :key="i">
            <div v-if="item.depth > 1" @click="jump">{{ item.title }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import type { Toc } from '@/models/post'

const props = defineProps({
    toc: {
        type: Array as () => Toc[],
        required: true
    }
})

const { toc } = toRefs(props)

function jump(e: MouseEvent) {
    const header = (e.target as HTMLElement).innerText
    document.querySelectorAll(`h2`).forEach((element: HTMLElement) => {
        if (element.innerText == header) {
            element.scrollIntoView({ behavior: 'smooth' })
        }
    })
}
</script>
