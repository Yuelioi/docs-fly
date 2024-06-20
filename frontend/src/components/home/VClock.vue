<template>
    <div
        class="container flex items-center justify-center w-screen pb-8 mx-auto mt-6 mb-2 h-36 animated-clock meters">
        <img
            class="h-full"
            v-for="(data, index) in clockData"
            :key="index"
            :src="getImageUrl(data)"
            alt="" />
    </div>
</template>

<script setup lang="ts">
import { getImageUrl } from '@/utils/divergence'

class Clock {
    pre: string = '1.'
    month: string = '00'
    day: string = '00'
    hour: string = '00'
    minute: string = '00'
    second: string = '00'

    getProp(propName: keyof Clock): string {
        return this[propName] as string
    }
}

const clock = ref<Clock>(new Clock())
const clockData = ref<string[]>([])

function refreshClock() {
    const date = new Date()

    clock.value.month = ('0' + (date.getMonth() + 1)).slice(-2)
    clock.value.day = ('0' + date.getDate()).slice(-2)
    clock.value.hour = ('0' + date.getHours()).slice(-2)
    clock.value.minute = ('0' + date.getMinutes()).slice(-2)
    clock.value.second = ('0' + date.getSeconds()).slice(-2)

    let res = []

    for (const key of Object.keys(clock.value)) {
        const value = clock.value[key as keyof Clock] as string
        for (const y of value) {
            res.push(y)
        }
    }

    clockData.value = res
}

setInterval(() => {
    refreshClock()
}, 999)
</script>
<style scoped>
img {
    width: calc(100% / 12);
}
.animated-clock {
    animation: fadeInOut 3s ease-in-out;
}

@keyframes fadeInOut {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
}
</style>
