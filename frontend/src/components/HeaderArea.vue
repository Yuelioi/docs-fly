<template>
    <div class="h-16 w-full basis-10 shadow-md">
        <div
            class="flex flex-row h-16 items-center justify-around sm:mx-[2rem] md:mx-[3rem] lg:mx-[5rem] xl:mx-[7.5rem] 2xl:mx-[10rem]">
            <div class="left flex cat-menu">
                <div class="justify-start h-16 hidden lg:block">
                    <a rel="home" href="#"
                        ><img
                            class="h-full"
                            itemprop="logo"
                            src="https://cdn.yuelili.com/web/assets/logo.webp"
                    /></a>
                </div>
                <!-- 分类 Start -->
                <div class="flex flex-row">
                    <div
                        class="relative group items-center flex rounded-lg h-full p-1 pl-4 pr-4"
                        v-for="(nav, index_nav) in navs"
                        :key="index_nav">
                        <span class="font-bold cursor-default">{{ nav.display_name }}</span>

                        <ul
                            class="absolute top-16 rounded-lg bg-theme-card scale-0 group-hover:scale-100 ease-in-out duration-300 origin-top-left z-50">
                            <router-link
                                class="px-6 py-3 flex items-center last:pb-4 hover:bg-theme-primary-hover w-full whitespace-nowrap"
                                v-for="(child, index_item) in nav.children"
                                :key="index_item"
                                :to="{
                                    name: 'book',
                                    params: {
                                        category: nav.identity,
                                        book: child.identity,
                                        locale: locale
                                    }
                                }"
                                ><i class="pi pi-book"></i
                                ><span class="pl-2">{{ child.display_name }}</span></router-link
                            >
                        </ul>
                    </div>
                </div>
                <!-- 分类 End-->
            </div>
            <div
                class="header-banner-item header-banner-right items-center justify-center flex h-full">
                <!-- 搜索 Start -->
                <div class="search">
                    <HSearch v-model:showSearchDialog="showSearchDialog" />
                </div>
                <!-- 搜索 End -->

                <!-- 右侧工具 -->
                <div class="flex items-center">
                    <!-- pi-sun -->
                    <button @click="toggleDark()">
                        <i
                            class="pi pi-moon ml-2 p-2 text-lg rounded-lg"
                            :class="isDark ? 'i-sun' : 'pi-moon'"></i>
                    </button>

                    <button>
                        <router-link :to="{ name: 'star' }" :key="'star'">
                            <i class="pi pi-star ml-2 fontsize p-2 text-lg rounded-lg"></i>
                        </router-link>
                    </button>
                    <button @click="changeLocale">
                        <i class="pi pi-language ml-2 fontsize p-2 text-lg rounded-lg"></i>
                    </button>
                    <button v-if="!isAdmin" @click="showLoginWindow = true">
                        <i class="pi pi-user ml-2 fontsize p-2 text-lg rounded-lg"></i>
                    </button>
                    <button v-else @click="logout">
                        <i class="pi pi-sign-in ml-2 fontsize p-2 text-lg rounded-lg"></i>
                    </button>

                    <div v-if="showLoginWindow">
                        <LoginWindow v-model:showLoginWindow="showLoginWindow"></LoginWindow>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <HSearchWithDialog v-model:showSearchDialog="showSearchDialog"> </HSearchWithDialog>
</template>

<script setup lang="ts">
import { NavData } from '@/models'
import { ref, onMounted } from 'vue'
import { fetchNav } from '@/handlers/index'
import { getNav, addNav } from '@/database'

import LoginWindow from '@/components/LoginWindow.vue'
import HSearch from '@/components/HSearch.vue'
import HSearchWithDialog from '@/components/HSearchWithDialog.vue'

import { storeToRefs } from 'pinia'

import { useRoute, useRouter } from 'vue-router'

import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)

const route = useRoute()
const router = useRouter()

import { basicStore } from '@/stores/index'
import { Message } from '@/plugins/message'
const basic = basicStore()
const { locale, token, isAdmin } = storeToRefs(basic)
const { translate } = basic

const navs = ref<NavData[]>([])

const showLoginWindow = ref(false)
const showSearchDialog = ref(false)

// 更改语言设置
function changeLocale() {
    if (locale.value == 'en') {
        locale.value = 'zh'
    } else {
        locale.value = 'en'
    }

    Message(`已切换为${translate('locale')}`)
    localStorage.setItem('locale', locale.value)

    const newParams = { ...route.params, locale: locale.value } // 替换 'newLocaleValue' 为你想要设置的新的参数值
    const newRoute = { ...route, params: newParams }

    router.replace(newRoute)
}

// 登出
function logout() {
    token.value = ''
    isAdmin.value = false
    localStorage.removeItem('token')
    Message('已成功登出', 'success')
}

onMounted(async () => {
    const nav_data = await getNav()

    if (nav_data) {
        navs.value = nav_data.data
    } else {
        const [ok, data] = await fetchNav()
        if (ok) {
            navs.value = data.sort((pre: NavData, next: NavData) => pre.order - next.order)
            await addNav(navs.value)
        } else {
            navs.value = [new NavData()]
        }
    }
})
</script>

<style scoped lang="css">
.cat-menu a.router-link-active.router-link-exact-active {
    border-left: #0088ff solid 6px;
}
</style>
