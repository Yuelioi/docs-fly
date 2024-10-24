<template>
  <div class="fixed bottom-0 flex items-center justify-between bg-base-200 w-full h-12 border">
    <div class="flex items-center justify-around flex-1">
      <div>
        <a href="#" class="flex flex-col items-center">
          <BIconHouseHeart class="text-icon-base"></BIconHouseHeart
          ><span class="icon-[lucide--home] size-6"></span
        ></a>
      </div>
      <div
        class=""
        @click.prevent.stop="
          router.push({
            name: 'books'
          })
        ">
        <span class="icon-[lucide--library-big] size-5"></span>
      </div>
      <div class="" @click="toggleDarkMode()">
        <span class="icon-[lucide--moon] size-5" v-if="theme == 'light'"></span>
        <span class="icon-[lucide--sun] size-5" v-else> </span>
      </div>
      <div class="">
        <span
          class="icon-[lucide--search] size-6"
          @click.prevent.stop="
            router.push({
              name: 'search'
            })
          "></span>
      </div>

      <div class="">
        <router-link :to="{ name: 'star' }" :key="'star'">
          <div
            class="flex flex-col items-center p-2 ml-2 rounded-lg outline-1 hover:outline fontsize">
            <span class="icon-[lucide--star] size-5"></span>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core'

const router = useRouter()

const theme = useLocalStorage('theme', 'light')
const _themes = ['light', 'dark']
const { themes, switchTheme } = useTheme(_themes)

function toggleDarkMode() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}
</script>
