<template>
    <div id="toc">
        <div
            class="toc-item text-sm border-l-2 hover:bg-theme-card border-theme-base hover:border-theme-primary hover:px-6 p-1 pl-4 select-none cursor-pointer"
            v-for="(item, i) in toc"
            :key="i">
            <div v-if="item.depth > 1" @click="jump">{{ item.title }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import type { Toc } from '@/models'

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
