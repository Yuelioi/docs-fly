<template>
    <transition name="bounce">
        <div
            v-show="show"
            ref="dialog"
            :key="Date.now().toString()"
            :class="data[severity as keyof typeof data].main"
            class="relative font-bold border text-wrap flex min-w-60 max-w-[24rem] rounded-lg items-center">
            <div class="flex items-center w-full">
                <i class="ml-4 pi" :class="data[severity as keyof typeof data].icon"></i>
                <span class="pl-4 w-4/5 py-2 break-words">{{ message }}</span>
                <i class="pi pi-times absolute right-3" @click="close"></i>
            </div>
        </div>
    </transition>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref } from 'vue'
import type { MessageType } from './'
defineProps<{
    severity: MessageType
    message: string
}>()

const show = ref(false)
const dialog = ref<HTMLElement | null>(null)
const data = {
    success: {
        main: 'bg-green-50 border-green-300 text-green-600',
        icon: 'pi-check-circle'
    },
    info: {
        main: 'bg-blue-50 border-blue--300 text-blue-600',
        icon: 'pi-info-circle'
    },
    warn: {
        main: 'bg-yellow-50 border-yellow-300 text-yellow-600',
        icon: 'pi-exclamation-triangle'
    },
    error: {
        main: 'bg-red-50 border-red-300 text-red-600',
        icon: 'pi-times-circle'
    },
    secondary: {
        main: 'bg-slate-50 border-slate-300 text-slate-600',
        icon: ''
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
