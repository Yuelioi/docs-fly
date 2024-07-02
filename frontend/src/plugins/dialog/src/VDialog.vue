<template>
    <Teleport to="body">
        <transition name="fade">
            <div
                v-if="show"
                class="fixed top-0 left-0 z-40 w-screen h-screen opacity-50 overlay bg-zinc-950"></div>
        </transition>
        <transition name="bounce">
            <div
                ref="dialogRef"
                v-if="show"
                class="fixed z-50 -translate-x-1/2 rounded-lg shadow-2xl left-1/2 right-1/2 top-32 bg-theme-card w-[calc(90vw)]">
                <div class="flex flex-col p-4">
                    <header class="flex items-center justify-between">
                        <slot name="header">{{ title }}</slot>

                        <BIconX class="m-2 ml-auto text-icon-md" @click="show = false"></BIconX>
                    </header>
                    <main>
                        <slot></slot>
                    </main>
                    <footer>
                        <slot name="footer"></slot>
                    </footer>
                </div>
            </div>
        </transition>
    </Teleport>
</template>

<script lang="ts" setup>
import { BIconX } from 'bootstrap-icons-vue'

const show = defineModel('show', { default: false, required: true })
const dialogRef = ref<HTMLElement | null>(null)

defineProps({
    title: {
        require: false,
        type: String
    }
})

const handleClickOutside = (e: MouseEvent) => {
    if (dialogRef.value && !dialogRef.value.contains(e.target as Node)) {
        show.value = false
    }
}

onMounted(() => {
    document.body.classList.add('dialog')
    document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
    document.body.classList.remove('dialog')
    document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}
</style>

<style scoped>
.bounce-enter-active {
    animation: bounce-in 0.5s;
}
.bounce-leave-active {
    animation: bounce-in 0.1s reverse;
}

@keyframes bounce-in {
    0% {
        transform: scale(0);
        transform: translateY(500px);
        transform: translateX(calc(-50%));
    }
    50% {
        transform: scale(1.1);
        transform: translateY(600px);
        transform: translateX(calc(-50%));
    }
    100% {
        transform: scale(1);
        transform: translateY(700px);
        transform: translateX(calc(-50%));
    }
}
</style>
