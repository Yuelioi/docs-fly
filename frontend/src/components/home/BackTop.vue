<template>
    <transition appear>
        <button
            v-show="show"
            id="scrollToTop"
            class="fixed right-12 duration-200 transition-opacity ease-in-out bottom-12 w-8 h-8 bg-theme-backToTop-base hover:bg-theme-backToTop-hover p-3 rounded-full flex items-center justify-center"
            @click="scrollToTop()">
            <div class="text-[1.25rem]"><BIconCaretUp></BIconCaretUp></div></button
    ></transition>
</template>

<script lang="ts" setup>
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
