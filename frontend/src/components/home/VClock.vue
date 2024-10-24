<template>
  <div class="mx-auto h-36">
    <div class="container flex h-full items-center justify-center inset-4 animated-clock">
      <img
        class="h-full"
        v-for="(data, index) in clockData"
        :key="index"
        :src="imageCache[data]"
        alt="" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { getImages } from '@/utils/divergence'

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

const imageCache = ref<{ [key: string]: string }>({})

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
  if (Object.keys(imageCache.value).length === 0) {
    imageCache.value = getImages()
  }
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
