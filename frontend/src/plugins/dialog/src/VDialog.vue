<template>
    <Teleport to="body">
        <transition name="fade">
            <div
                v-if="show"
                class="fixed top-0 left-0 z-40 w-screen h-screen opacity-50 overlay"></div>
        </transition>
        <transition name="bounce">
            <div v-if="show" class="fixed inset-0 flex items-center justify-center z-50">
                <div ref="dialogRef" class="bg-neutral-content w-[75%]">
                    <div class="flex flex-col p-4">
                        <header class="flex items-center justify-between">
                            <slot name="header">{{ title }}</slot>

                            <i class="m-2 ml-auto text-icon-md" @click="closeDialog"
                                ><svg
                                    t="1720599521905"
                                    class="icon"
                                    viewBox="0 0 1024 1024"
                                    version="1.1"
                                    xmlns="http://www.w3.org/2000/svg"
                                    p-id="7759"
                                    width="20"
                                    height="20">
                                    <path
                                        d="M301.226667 210.773333a64 64 0 1 0-90.453334 90.453334L421.461333 512l-210.773333 210.773333a64 64 0 0 0 90.538667 90.453334L512 602.538667l210.773333 210.773333a64 64 0 0 0 90.453334-90.538667L602.538667 512l210.773333-210.773333a64 64 0 0 0-90.538667-90.453334L512 421.461333l-210.773333-210.773333z"
                                        fill="#515151"
                                        p-id="7760"></path></svg
                            ></i>
                        </header>
                        <div class="divider"></div>
                        <main
                            class="text-base-content min-w-[375px] min-h-[500px] max-h-[65vh] overflow-y-scroll">
                            <slot></slot>
                        </main>
                        <footer>
                            <slot name="footer"></slot>
                        </footer>
                    </div>
                </div>
            </div>
        </transition>
    </Teleport>
</template>

<script lang="ts" setup>
import { BIconX } from 'bootstrap-icons-vue'

defineProps({
    title: {
        type: String,
        required: false
    }
})

const show = defineModel('show', { default: false, required: true })
const dialogRef = ref<HTMLElement | null>(null)

const closeDialog = () => {
    show.value = false
}

const handleClickOutside = (e: MouseEvent) => {
    if (dialogRef.value && !dialogRef.value.contains(e.target as Node)) {
        closeDialog()
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

.bounce-enter-active {
    animation: bounce-in 0.5s;
}

.bounce-leave-active {
    animation: bounce-in 0.1s reverse;
}

@keyframes bounce-in {
    0% {
        transform: scale(0);
    }
    50% {
        transform: scale(1.1);
    }
    100% {
        transform: scale(1);
    }
}
</style>
