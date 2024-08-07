<template>
    <div class="container flex items-center h-16">
        <div class="flex items-center justify-around w-full sm:justify-between">
            <!-- # region logo区域 -->
            <div class="flex items-center">
                <a rel="home" href="#">
                    <img
                        class="h-8"
                        itemprop="logo"
                        src="https://cdn.yuelili.com/docs/web/assets/ydocs-256.png" />
                </a>
            </div>
            <!-- # endregion -->
            <!-- # region PC左侧菜单  -->
            <div v-if="width > 576" class="flex flex-row">
                <div
                    class="relative flex items-center h-16 px-2 rounded-lg group"
                    v-for="(nav, index_nav) in filteredNavs"
                    :key="index_nav">
                    <span class="text-sm font-bold cursor-default">{{ nav.metadata.title }}</span>

                    <ul
                        class="absolute top-[4.25rem] z-50 duration-300 ease-in-out origin-top-left scale-0 rounded-b-lg group-hover:scale-100">
                        <router-link
                            class="flex items-center w-full px-3 py-2 last:pb-4 last:hover:rounded-b-lg hover: whitespace-nowrap"
                            v-for="(child, index_item) in sortMeta(nav.children)"
                            :key="index_item"
                            :to="{
                                name: 'book',
                                params: {
                                    bookPath: child.url.split('/')
                                }
                            }"
                            ><div class="text-[1rem]"><BIconBook></BIconBook></div>

                            <span class="pl-2">{{ child.title }}</span></router-link
                        >
                    </ul>
                </div>
            </div>
            <!-- # endregion  -->

            <!-- # region 手机端右侧菜单 -->
            <div v-if="width < 576" class="flex ml-auto group z-50">
                <BIconList class="text-icon-base"></BIconList>
                <div class="absolute top-[4.25rem] right-2 w-24 group-hover:scale-100 scale-0">
                    <div
                        class="flex flex-col h-8 items-center m-1 last:rounded-b px-4 py-1 hover: group/sub duration-300 ease-in-out origin-top-left"
                        v-for="(nav, index_nav) in filteredNavs"
                        :key="index_nav">
                        <div class="">
                            <span class="text-sm cursor-default">{{ nav.metadata.title }}</span>
                        </div>

                        <ul
                            class="group-hover/sub:scale-100 scale-0 absolute right-[100%] top-0 rounded-b-lg">
                            <router-link
                                class="flex items-center m-1 w-32 h-8 px-2 hover:"
                                v-for="(child, index_item) in sortMeta(nav.children)"
                                :key="index_item"
                                :to="{
                                    name: 'book',
                                    params: {
                                        bookPath: child.url.split('/')
                                    }
                                }"
                                ><div class="text-[1rem]"><BIconBook></BIconBook></div>

                                <span class="pl-2 truncate">{{ child.title }}</span></router-link
                            >
                        </ul>
                    </div>
                </div>
            </div>
            <!-- # endregion-->

            <!-- # region PC端右侧菜单 -->
            <div class="items-center justify-center relative hidden sm:flex">
                <div class="relative flex items-center justify-center h-16 select-none">
                    <div class="absolute -mt-2 top-2/4 left-3">
                        <BIconSearch></BIconSearch>
                    </div>
                    <span
                        @click.prevent.stop="showSearchDialog = true"
                        class="flex items-center h-10 pl-10 pr-4 text-sm bg-transparent border-2 rounded-full hover:">
                        搜索文档...
                    </span>
                </div>
                <button @click="toggleColorScheme()">
                    <div class="p-2 ml-2 rounded-lg outline-1 hover:outline">
                        <BIconSun v-if="isDark"></BIconSun>
                        <BIconMoonStars v-else></BIconMoonStars>
                    </div>
                </button>
                <div class="flex items-center justify-center">
                    <div
                        class="p-2 ml-2 text-lg group/theme rounded-lg outline-1 hover:outline fontsize">
                        <BIconPalette></BIconPalette>
                        <div
                            class="group-hover/theme:scale-100 scale-0 absolute top-full z-40 rounded-b-lg transition-transform duration-500">
                            <select
                                v-model="theme"
                                class="select select-sm text-base-content text-sm data-choose-theme w-full max-w-xs">
                                <option :value="_theme" v-for="_theme in themes" :key="_theme">
                                    {{ _theme }}
                                </option>
                            </select>
                        </div>
                    </div>

                    <button>
                        <router-link :to="{ name: 'star' }" :key="'star'">
                            <div class="p-2 ml-2 rounded-lg outline-1 hover:outline fontsize">
                                <BIconStar></BIconStar>
                            </div>
                        </router-link>
                    </button>
                    <button @click="changeLocale">
                        <div class="p-2 ml-2 text-lg rounded-lg outline-1 hover:outline fontsize">
                            <BIconTranslate></BIconTranslate>
                        </div>
                    </button>
                    <button v-if="!isAdmin" @click.prevent.stop="showLoginWindow = true">
                        <div class="p-2 ml-2 text-lg rounded-lg outline-1 hover:outline fontsize">
                            <BIconPerson></BIconPerson>
                        </div>
                    </button>
                    <button v-else @click="logout">
                        <div class="p-2 ml-2 text-lg rounded-lg outline-1 hover:outline fontsize">
                            <BIconBoxArrowRight></BIconBoxArrowRight>
                        </div>
                    </button>

                    <VDialog v-model:show="showLoginWindow">
                        <LoginWindow v-model:show="showLoginWindow"></LoginWindow>
                    </VDialog>
                </div>
            </div>
            <!-- # endregion-->
        </div>
    </div>

    <!-- # region 手机端底部导航 -->
    <div class="fixed bottom-0 flex items-center justify-between w-full h-12 mt-12 md:hidden">
        <div class="flex items-center justify-around flex-1">
            <div>
                <a href="#" class="flex flex-col items-center">
                    <BIconHouseHeart class="text-icon-base"></BIconHouseHeart
                    ><span class="text-xs">首页</span></a
                >
            </div>
            <div>
                <div class="flex flex-col items-center">
                    <BIconSearch></BIconSearch>
                    <span
                        @click.prevent.stop="showSearchDialog = true"
                        class="flex items-center text-sm bg-transparent">
                        搜索
                    </span>
                </div>
            </div>
        </div>
        <div class="center">中间</div>
        <div class="flex items-center justify-around flex-1">
            <button>
                <router-link :to="{ name: 'star' }" :key="'star'">
                    <div
                        class="flex flex-col items-center p-2 ml-2 rounded-lg outline-1 hover:outline fontsize">
                        <BIconStar class="text-icon-base"></BIconStar>
                        <span class="text-xs">收藏</span>
                    </div>
                </router-link>
            </button>
            <button v-if="!isAdmin" @click.prevent.stop="showLoginWindow = true">
                <div
                    class="flex flex-col items-center p-2 ml-2 rounded-lg outline-1 hover:outline fontsize">
                    <BIconPerson class="text-icon-base"></BIconPerson>
                    <span class="text-xs">登录</span>
                </div>
            </button>
            <button v-else @click="logout">
                <div class="p-2 ml-2 rounded-lg outline-1 hover:outline fontsize">
                    <BIconBoxArrowRight class="text-icon-base"></BIconBoxArrowRight>
                    <span class="text-xs">注销</span>
                </div>
            </button>
        </div>
    </div>
    <!-- # endregion -->

    <HSearchWithDialog v-model:showSearchDialog="showSearchDialog" v-model:navs="filteredNavs">
    </HSearchWithDialog>
</template>

<script setup lang="ts">
import { Nav } from '@/models/home'
import { MetaData } from '@/models/base'
import { useLocalStorage } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)

function toggleColorScheme() {
    const element = document.querySelector('html')
    element?.classList.toggle('dark')
    toggleDark()
}

import { useTheme } from '@/hooks/useTheme'

const theme = useLocalStorage('theme', '')
const _themes = [
    'light',
    'corporate',
    'lofi',
    'wireframe',
    'nord',
    'dark',
    'dracula',
    'dim',
    'sunset'
]
const { themes, switchTheme } = useTheme(_themes)

const route = useRoute()
const router = useRouter()

import { useWindowSize } from '@vueuse/core'

const { width } = useWindowSize()

import {
    BIconBook,
    BIconBoxArrowRight,
    BIconPerson,
    BIconStar,
    BIconTranslate,
    BIconList,
    BIconSearch,
    BIconHouseHeart,
    BIconPalette,
    BIconSun,
    BIconMoonStars
} from 'bootstrap-icons-vue'
const basic = basicStore()
const { locale, isAdmin } = storeToRefs(basic)
const { translate } = basic

const navs = ref<Nav[]>([])

const filteredNavs = computed(() => {
    return navs.value.slice().sort((pre, next) => {
        return (pre.metadata.order = next.metadata.order)
    })
})

const showLoginWindow = ref(false)
const showSearchDialog = ref(false)

// 更改语言设置
async function changeLocale() {
    const lastLocale = locale.value

    if (locale.value == 'en') {
        locale.value = 'zh'
    } else {
        locale.value = 'en'
    }

    Message({ message: `已切换为${translate('locale')}` })
    localStorage.setItem('locale', locale.value)

    const routeParams = route.params
    const updatedParams: RouteParams = {}

    Object.keys(routeParams).forEach((key) => {
        if (Array.isArray(routeParams[key])) {
            const paramArray = routeParams[key] as string[]

            const newParamArray = paramArray.map((ele: string) =>
                ele === lastLocale ? locale.value : ele
            )
            updatedParams[key] = newParamArray
        } else {
            updatedParams[key] = routeParams[key] === lastLocale ? locale.value : routeParams[key]
        }
    })

    const newParams = { ...updatedParams } // 替换 'newLocaleValue' 为你想要设置的新的参数值
    const newRoute = { ...route, params: newParams }

    router.replace(newRoute)
}

// 登出
async function logout() {
    isAdmin.value = false
    localStorage.removeItem('token')
    Message({ message: '已成功登出', type: 'success' })
}

function sortMeta(data: MetaData[]) {
    return data.sort((pre: MetaData, next: MetaData) => {
        return pre.order - next.order
    })
}

watch(theme, () => {
    switchTheme(theme.value)
})

onMounted(async () => {
    const nav_data = await getDBNav()
    // 验证Nav数据库信息并排序
    if (nav_data) {
        navs.value = nav_data.data
    } else {
        const [ok, data] = await getNav()
        if (ok) {
            navs.value = data['data'].sort(
                (pre: Nav, next: Nav) => pre.metadata.order - next.metadata.order
            )
            await addDBNav(navs.value)
        } else {
            navs.value = [new Nav()]
        }
    }

    switchTheme(theme.value)
})
</script>

<style scoped lang="css">
.cat-menu a.router-link-active.router-link-exact-active {
    border-left: #0088ff solid 6px;
}
</style>
