<template>
    <transition name="bounce">
        <div
            v-show="show"
            ref="dialog"
            :key="Date.now().toString()"
            :class="data[type as keyof typeof data].main"
            class="relative font-bold border text-wrap flex min-w-60 max-w-[24rem] rounded-lg items-center">
            <div class="flex items-center w-full">
                <component class="pl-4" :is="data[type as messageType].icon"></component>
                <span class="w-4/5 py-2 pl-4 break-words">{{ message }}</span>
                <BIconX @click="close"></BIconX>
            </div>
        </div>
    </transition>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref } from 'vue'
import type { messageType } from './message'

import {
    BIconCheck2Circle,
    BIconInfoCircle,
    BIconExclamationCircle,
    BIconXCircle,
    BIconX
} from 'bootstrap-icons-vue'

defineProps<{
    type: messageType
    message: string
}>()

const show = ref(false)
const dialog = ref<HTMLElement | null>(null)
const data: Record<messageType, { main: string; icon: any }> = {
    success: {
        main: 'bg-theme-success-base border-green-300 text-green-600',
        icon: BIconCheck2Circle
    },
    secondary: {
        main: 'bg-black border-slate-300 text-slate-200',
        icon: BIconCheck2Circle
    },
    info: {
        main: 'bg-blue-50 border-blue--300 text-blue-600',
        icon: BIconInfoCircle
    },
    warn: {
        main: 'bg-yellow-50 border-yellow-300 text-yellow-600',
        icon: BIconExclamationCircle
    },
    error: {
        main: 'bg-red-50 border-red-300 text-red-600',
        icon: BIconXCircle
    },
    contrast: {
        main: 'bg-black border-slate-300 text-slate-200',
        icon: ''
    }
}

onBeforeUnmount(() => {
    show.value = false
})
onMounted(() => {
    show.value = true
})

function close() {
    if (dialog.value) {
        ;(dialog.value.parentNode as HTMLElement).removeChild(dialog.value)
    }
}
</script>

<style scoped>
.bounce-enter-active {
    animation: bounce-in 0.5s;
}
.bounce-leave-active {
    animation: bounce-in 0.5s reverse;
}

@keyframes bounce-in {
    0% {
        transform: scale(0);
        transform: translateX(20px);
    }
    50% {
        transform: scale(1.1);
    }
    100% {
        transform: scale(1);
    }
}
</style>
