<template>
    <transition name="bounce">
        <div
            v-show="visible"
            ref="messageRef"
            :class="data[props.type as keyof typeof data].main"
            :key="Date.now().toString()"
            class="mt-4 relative font-bold border text-wrap flex min-w-60 max-w-[24rem] rounded-lg items-center">
            <div class="flex items-center w-full">
                <span class="w-4/5 py-2 pl-4 break-words">{{ message }}</span>
                <component class="pl-4" :is="data[props.type as messageType].icon"></component>
                <BIconX @click="close" v-if="props.showClose"></BIconX>
            </div>
        </div>
    </transition>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref } from 'vue'
import type { messageType } from './model'

import {
    BIconCheck2Circle,
    BIconInfoCircle,
    BIconExclamationCircle,
    BIconXCircle,
    BIconX
} from 'bootstrap-icons-vue'

const props = defineProps<{
    type: messageType
    message: string
    duration: number
    showClose: boolean
}>()

const visible = ref(false)
const messageRef = ref<HTMLElement | null>(null)
const data: Record<messageType, { main: string; icon: any }> = {
    success: {
        main: 'bg-green-50 border-green-300 text-green-600',
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
    secondary: {
        main: 'bg-black border-slate-300 text-slate-200',
        icon: BIconCheck2Circle
    },
    contrast: {
        main: 'bg-black border-slate-300 text-slate-200',
        icon: ''
    }
}

onBeforeUnmount(() => {})
onMounted(() => {
    visible.value = true
    // 定时删除子元素
    // setTimeout(() => {}, 3000)
})

function close() {
    visible.value = false
    if (messageRef.value) {
        const messageDiv = messageRef.value.parentNode as HTMLElement
        messageDiv.parentNode?.removeChild(messageDiv)
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
