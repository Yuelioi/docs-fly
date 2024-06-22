<template>
    <transition appear>
        <button
            v-show="show"
            id="scrollToTop"
            class="fixed flex items-center justify-center w-8 h-8 p-3 transition-opacity duration-200 ease-in-out rounded-full right-12 bottom-12 bg-theme-backToTop-base hover:bg-theme-backToTop-hover"
            @click="scrollToTop()">
            <div class="text-[1.25rem]"><BIconCaretUp></BIconCaretUp></div></button
    ></transition>
</template>

<script lang="ts" setup>
import { BIconCaretUp } from 'bootstrap-icons-vue'

function scrollToTop() {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    })
}

const show = ref(false)

onMounted(() => {
    window.addEventListener('scroll', function () {
        if (window.scrollY > 100) {
            show.value = true
        } else {
            show.value = false
        }
    })
})
</script>
