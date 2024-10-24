<template>
  <div class="flex items-center justify-around w-full sm:justify-between px-6">
    <!-- logo -->
    <div class="flex items-center">
      <a rel="home" href="#">
        <img
          class="h-8"
          itemprop="logo"
          src="https://cdn.yuelili.com/docs/web/assets/ydocs-256.png" />
      </a>
    </div>

    <!-- 功能 -->
    <div class="items-center justify-center flex relative space-x-4">
      <div class="relative flex items-center justify-center h-16 cursor-pointer">
        <span class="icon-[lucide--search] size-5 absolute left-3"></span>
        <span
          @click.prevent.stop="
            router.push({
              name: 'search'
            })
          "
          class="flex items-center h-10 pl-10 pr-4 text-sm select-none border-2 rounded-full hover:">
          搜索文档...
        </span>
      </div>
      <div
        class="btn btn-sm btn-outline btn-circle"
        @click.prevent.stop="
          router.push({
            name: 'books'
          })
        ">
        <span class="icon-[lucide--library-big] size-5"></span>
      </div>

      <label class="swap swap-rotate btn btn-sm btn-outline btn-circle">
        <input type="checkbox" @click="toggleDarkMode()" />
        <span class="icon-[lucide--moon] size-5 swap-on"></span>
        <span class="icon-[lucide--sun] size-5 swap-off"></span>
      </label>

      <div class="btn btn-sm btn-outline btn-circle">
        <router-link :to="{ name: 'star' }" :key="'star'">
          <span class="icon-[lucide--star] size-5"></span>
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
  console.log(theme.value)
  switchTheme(theme.value)
}
</script>
