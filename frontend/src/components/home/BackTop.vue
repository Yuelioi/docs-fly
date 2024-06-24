<template>
    <transition appear>
        <button
            v-show="show"
            id="scrollToTop"
            class="fixed flex items-center justify-center transition-opacity duration-200 ease-in-out rounded-full size-6 right-4 bottom-16 bg-theme-backToTop-base hover:bg-theme-backToTop-hover"
            @click="scrollToTop()">
            <div class="text-icon-md"><BIconCaretUp></BIconCaretUp></div></button
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
            document.body.classList.add('scrolling')
            show.value = true
        } else {
            show.value = false
            document.body.classList.remove('scrolling')
        }
    })
})
</script>
