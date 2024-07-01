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
            <!-- # endregion logo区域 -->
            <!-- # region PC左侧菜单  -->
            <div v-if="width > 576" class="flex flex-row">
                <div
                    class="relative flex items-center h-16 px-2 rounded-lg group"
                    v-for="(nav, index_nav) in filteredNavs"
                    :key="index_nav">
                    <span class="text-sm font-bold cursor-default">{{ nav.metadata.title }}</span>

                    <ul
                        class="absolute top-[calc(100%-1px)] z-50 duration-300 ease-in-out origin-top-left scale-0 rounded-b-lg bg-theme-card group-hover:scale-100">
                        <router-link
                            class="flex items-center w-full px-3 py-2 last:pb-4 last:hover:rounded-b-lg hover:bg-theme-primary-hover whitespace-nowrap"
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
            <!-- # endregion PC左侧菜单  -->

            <!-- # region 手机端右侧菜单 -->
            <div v-if="width < 576" class="flex ml-auto">
                <BIconList class="text-icon-base"></BIconList>
            </div>
            <!-- # endregion 手机端右侧菜单 手机端-->

            <!-- # region PC端右侧菜单 -->
            <div class="items-center justify-center hidden sm:flex">
                <!-- 搜索 Start -->

                <div class="relative flex items-center justify-center h-16 select-none">
                    <div
                        class="absolute -mt-2 top-2/4 left-3 text-surface-400 dark:text-surface-600">
                        <BIconSearch></BIconSearch>
                    </div>
                    <span
                        @click.prevent.stop="showSearchDialog = true"
                        class="flex items-center h-10 pl-10 pr-4 text-sm bg-transparent border-2 rounded-full border-theme-base hover:border-theme-primary">
                        搜索文档...
                    </span>
                </div>

                <!-- 搜索 End -->

                <div class="flex items-center">
                    <button @click="toggleDark()">
                        <div
                            class="p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline">
                            <BIconSun v-if="isDark"></BIconSun>
                            <BIconMoonStars v-else></BIconMoonStars>
                        </div>
                    </button>
                    <button class="relative h-16 group">
                        <div
                            class="p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline">
                            <BIconPalette class=""></BIconPalette>
                            <div
                                class="absolute top-[calc(100%+1px)] -right-1/2 z-50 duration-300 ease-in-out origin-top-left scale-0 rounded-b-lg bg-theme-card group-hover:scale-100">
                                <div
                                    class="w-full px-4 py-2 last:pb-4 last:hover:rounded-b-lg hover:bg-theme-primary-hover whitespace-nowrap"
                                    v-for="theme in themes"
                                    @click="switchTheme(theme)"
                                    :key="theme">
                                    <span class="">{{ theme }}</span>
                                </div>
                            </div>
                        </div>
                    </button>

                    <button>
                        <router-link :to="{ name: 'star' }" :key="'star'">
                            <div
                                class="p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                                <BIconStar></BIconStar>
                            </div>
                        </router-link>
                    </button>
                    <button @click="changeLocale">
                        <div
                            class="p-2 ml-2 text-lg rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                            <BIconTranslate></BIconTranslate>
                        </div>
                    </button>
                    <button v-if="!isAdmin" @click.prevent.stop="showLoginWindow = true">
                        <div
                            class="p-2 ml-2 text-lg rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                            <BIconPerson></BIconPerson>
                        </div>
                    </button>
                    <button v-else @click="logout">
                        <div
                            class="p-2 ml-2 text-lg rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                            <BIconBoxArrowRight></BIconBoxArrowRight>
                        </div>
                    </button>

                    <VDialog v-model:show="showLoginWindow">
                        <LoginWindow></LoginWindow>
                    </VDialog>
                </div>
            </div>
            <!-- # endregion PC端右侧菜单 PC端-->
        </div>
    </div>

    <!-- # region 手机端底部导航 -->
    <div
        class="fixed bottom-0 flex items-center justify-between w-full h-12 mt-12 md:hidden bg-theme-card">
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
                        class="flex flex-col items-center p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                        <BIconStar class="text-icon-base"></BIconStar>
                        <span class="text-xs">收藏</span>
                    </div>
                </router-link>
            </button>
            <button v-if="!isAdmin" @click.prevent.stop="showLoginWindow = true">
                <div
                    class="flex flex-col items-center p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                    <BIconPerson class="text-icon-base"></BIconPerson>
                    <span class="text-xs">登录</span>
                </div>
            </button>
            <button v-else @click="logout">
                <div
                    class="p-2 ml-2 rounded-lg outline-theme-primary outline-1 hover:outline fontsize">
                    <BIconBoxArrowRight class="text-icon-base"></BIconBoxArrowRight>
                    <span class="text-xs">注销</span>
                </div>
            </button>
        </div>
    </div>
    <!-- # endregion 手机端底部导航 -->

    <HSearchWithDialog v-model:showSearchDialog="showSearchDialog" v-model:navs="filteredNavs">
    </HSearchWithDialog>
</template>

<script setup lang="ts">
import { Nav } from '@/models/home'
import { MetaData } from '@/models/base'

const isDark = useDark()
const toggleDark = useToggle(isDark)

import { themes, switchTheme } from '@/hooks/useTheme'

const route = useRoute()
const router = useRouter()

import { useWindowSize } from '@vueuse/core'

const { width } = useWindowSize()

import {
    BIconBook,
    BIconBoxArrowRight,
    BIconMoonStars,
    BIconPerson,
    BIconStar,
    BIconSun,
    BIconTranslate,
    BIconList,
    BIconSearch,
    BIconHouseHeart,
    BIconPalette
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

    await Message({ message: `已切换为${translate('locale')}` })
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
    await Message({ message: '已成功登出', type: 'success' })
}

function sortMeta(data: MetaData[]) {
    return data.sort((pre: MetaData, next: MetaData) => {
        return pre.order - next.order
    })
}

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
})
</script>

<style scoped lang="css">
.cat-menu a.router-link-active.router-link-exact-active {
    border-left: #0088ff solid 6px;
}
</style>
