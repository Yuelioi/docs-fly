<template>
    <div class="h-16 w-full">
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
                        v-for="(nav, index_nav) in filteredNavs"
                        :key="index_nav">
                        <span class="font-bold cursor-default">{{ nav.metadata.title }}</span>

                        <ul
                            class="absolute top-16 rounded-lg bg-theme-card scale-0 group-hover:scale-100 ease-in-out duration-300 origin-top-left z-50">
                            <router-link
                                class="px-6 py-3 flex items-center last:pb-4 first:hover:rounded-t-lg last:hover:rounded-b-lg hover:bg-theme-primary-hover w-full whitespace-nowrap"
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
                    <button @click="toggleDark()">
                        <div
                            class="outline-theme-primary outline-1 hover:outline ml-2 p-2 text-lg rounded-lg">
                            <BIconSun v-if="isDark"></BIconSun> <BIconMoon v-else></BIconMoon>
                        </div>
                    </button>

                    <button>
                        <router-link :to="{ name: 'star' }" :key="'star'">
                            <div
                                class="outline-theme-primary outline-1 hover:outline ml-2 fontsize p-2 text-lg rounded-lg">
                                <BIconStar></BIconStar>
                            </div>
                        </router-link>
                    </button>
                    <button @click="changeLocale">
                        <div
                            class="outline-theme-primary outline-1 hover:outline ml-2 fontsize p-2 text-lg rounded-lg">
                            <BIconTranslate></BIconTranslate>
                        </div>
                    </button>
                    <button v-if="!isAdmin" @click="showLoginWindow = true">
                        <div
                            class="outline-theme-primary outline-1 hover:outline ml-2 fontsize p-2 text-lg rounded-lg">
                            <BIconPerson></BIconPerson>
                        </div>
                    </button>
                    <button v-else @click="logout">
                        <div
                            class="outline-theme-primary outline-1 hover:outline ml-2 fontsize p-2 text-lg rounded-lg">
                            <BIconBoxArrowRight></BIconBoxArrowRight>
                        </div>
                    </button>

                    <div v-if="showLoginWindow">
                        <LoginWindow v-model:showLoginWindow="showLoginWindow"></LoginWindow>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <HSearchWithDialog v-model:showSearchDialog="showSearchDialog" v-model:navs="filteredNavs">
    </HSearchWithDialog>
</template>

<script setup lang="ts">
import { Nav } from '@/models/home'
import { MetaData } from '@/models/base'

const isDark = useDark()
const toggleDark = useToggle(isDark)

const route = useRoute()
const router = useRouter()

import { Message } from '@/plugins/message'
import {
    BIconBook,
    BIconBoxArrowRight,
    BIconMoon,
    BIconPerson,
    BIconStar,
    BIconSun,
    BIconTranslate
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
